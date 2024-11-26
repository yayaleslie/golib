package kafka

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/IBM/sarama"
)

const (
	// assignor
	AssignorSticky     = "sticky"
	AssignorRoundRobin = "roundrobin"
	AssignorRange      = "range"
)

type ConsumerConfig struct {
	Debug    bool              `toml:"debug" yaml:"debug" json:"debug"`
	Brokers  []string          `toml:"brokers" yaml:"brokers" json:"brokers"`
	Topics   map[string]string `toml:"topics" yaml:"topics" json:"topics"`
	Group    string            `toml:"groupid" yaml:"groupid" json:"groupid"`
	Version  string            `toml:"version" yaml:"version" json:"version"`
	Assignor string            `toml:"assignor" yaml:"assignor" json:"assignor"`
	Newest   bool              `toml:"newest" yaml:"newest" json:"newest"`
	Verbose  bool              `toml:"verbose" yaml:"verbose" json:"verbose"`
	Worker   int               `toml:"worker" yaml:"worker" json:"worker"`
	Logger   logger            `json:"-"`
	TopicArr []string          `json:"-"`
}

func (cc *ConsumerConfig) SetTopics() {
	if cc == nil {
		return
	}

	cc.TopicArr = make([]string, 0, len(cc.Topics))
	for _, v := range cc.Topics {
		cc.TopicArr = append(cc.TopicArr, v)
	}
}

func (cc *ConsumerConfig) GetTopics() []string {
	if cc == nil {
		return nil
	}
	return cc.TopicArr
}

func (cc *ConsumerConfig) GetTopic(k string) string {
	if cc == nil {
		return ""
	}
	return cc.Topics[k]
}

type Consumer struct {
	c       *ConsumerConfig
	client  sarama.ConsumerGroup
	handler func(*sarama.ConsumerMessage) error
	logger  logger
	ctx     context.Context
	cancel  context.CancelFunc
	wg      *sync.WaitGroup
	signals chan os.Signal

	Ready chan bool
}

func (c *Consumer) run() {
	c.wg.Add(2)

	go c.receive()
	go c.logErr()
}

// Stop 停止 kafka
func (c *Consumer) Stop() {
	c.cancel()

	if err := c.client.Close(); err != nil {
		c.logger.Errorf("kafka consumer close failed | brokers: %v | groupid: %s | err: %v", c.c.Brokers, c.c.Group, err)
	}

	c.wg.Wait()
	c.logger.Warn("Consumer is stopped")
}

// 记录error
func (c *Consumer) logErr() {
	defer c.wg.Done()

	for {
		select {
		case <-c.ctx.Done():
			return
		case err := <-c.client.Errors():
			c.logger.Errorf("kafka consumer receive failed | brokers: %v | groupid: %s | err: %v", c.c.Brokers, c.c.Group, err)
		}
	}
}

func (c *Consumer) SignalsNotify() {
	select {
	case <-c.ctx.Done():
		return
	case sign := <-c.signals:
		c.logger.Debugf("kafka consumer receive signals notification | brokers: %v | groupid: %s | signal: %s", c.c.Brokers, c.c.Group, sign)

	}
}

func (c *Consumer) receive() {
	defer c.wg.Done()

	for {
		if err := c.client.Consume(c.ctx, c.c.GetTopics(), c); err != nil {
			if err.Error() == "context canceled" {
				c.logger.Warnf("[kafka][receive] Context canceled: | brokers: %v | groupid: %s", c.c.Brokers, c.c.Group)
			} else {
				c.logger.Errorf("[kafka][receive] Error from consumer: | brokers: %v | groupid: %s | err: %v", c.c.Brokers, c.c.Group, err)
			}
		}

		if e := c.ctx.Err(); e != nil {
			if e.Error() == "context canceled" {
				c.logger.Warnf("[kafka][receive] Context canceled: | brokers: %v | groupid: %s", c.c.Brokers, c.c.Group)
			} else {
				c.logger.Errorf("[kafka][receive] Context error: | brokers: %v | groupid: %s | err: %v", c.c.Brokers, c.c.Group, e)
			}
			return
		}

		c.Ready = make(chan bool)
	}
}

func (c *Consumer) Handler(handler func(message *sarama.ConsumerMessage) error) {
	c.handler = handler
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.Ready)
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		if c.c.Debug {
			c.logger.Infof("kafka consumer message | brokers: %v | groupid: %s | message: %v", c.c.Brokers, c.c.Group, toString(msg))
		}

		if c.handler != nil {
			if err := c.handler(msg); err != nil {
				c.logger.Errorf("kafka consumer handler failed | brokers: %v | groupid: %s | err: %v", c.c.Brokers, c.c.Group, err)
				continue
			}
		}

		session.MarkMessage(msg, "")
	}

	return nil
}

func NewConsumer(c *ConsumerConfig, handler ...func(message *sarama.ConsumerMessage) error) (consumer *Consumer, err error) {
	if c.Verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	version, err := sarama.ParseKafkaVersion(c.Version)
	if err != nil {
		c.Logger.Errorf("parsing kafka version failed: %v", err)
		return
	}

	conf := sarama.NewConfig()

	conf.Version = version
	conf.Consumer.Return.Errors = true
	conf.Consumer.Offsets.AutoCommit.Enable = true
	conf.Consumer.Offsets.AutoCommit.Interval = time.Second
	conf.ClientID = os.Getenv("HOSTNAME")

	if c.Newest {
		conf.Consumer.Offsets.Initial = sarama.OffsetNewest
	} else {
		conf.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	switch c.Assignor {
	case AssignorSticky:
		conf.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
	case AssignorRoundRobin:
		conf.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	case AssignorRange:
		conf.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	default:
		c.Logger.Errorf("unrecognized consumer group partition assignor: %s", c.Assignor)
		return
	}

	var hand func(message *sarama.ConsumerMessage) error
	if len(handler) > 0 {
		hand = handler[0]
	}

	consumer = &Consumer{
		c:       c,
		handler: hand,
		logger:  c.Logger,
		wg:      &sync.WaitGroup{},
		signals: make(chan os.Signal),
		Ready:   make(chan bool),
	}

	client, err := sarama.NewConsumerGroup(c.Brokers, c.Group, conf)

	if err != nil {
		return
	}

	consumer.client = client

	signal.Notify(consumer.signals, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	consumer.ctx, consumer.cancel = context.WithCancel(context.Background())
	consumer.run()
	return
}
