package kafka

import (
	"context"
	"errors"
	"os"
	"sync"

	"github.com/IBM/sarama"
)

type ProducerConfig struct {
	Brokers []string          `toml:"brokers" yaml:"brokers" json:"brokers"`
	Topics  map[string]string `toml:"topics" yaml:"topics" json:"topics"`
	Logger  logger            `json:"-"`
}

func (p *ProducerConfig) Topic(topic string) string {
	if v, ok := p.Topics[topic]; ok {
		return v
	}

	return topic
}

type Producer struct {
	c      *ProducerConfig
	client sarama.AsyncProducer
	logger logger
	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup
}

func (p *Producer) Stop() {
	p.cancel()

	if err := p.client.Close(); err != nil {
		p.logger.Errorf("kafka producer close failed | brokers: %v | err: %v", p.c.Brokers, err)
	}

	p.wg.Wait()
}

func (p *Producer) Send(topic string, payload []byte, key ...[]byte) error {
	select {
	case <-p.ctx.Done():
		return errors.New("producer is stopped")
	default:
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(payload),
		}

		if len(key) > 0 {
			msg.Key = sarama.ByteEncoder(key[0])
		}

		//p.logger.Infof("send kafka msg success | topic: %s", topic)
		p.client.Input() <- msg
		return nil
	}
}

func (p *Producer) logErr() {
	defer p.wg.Done()

	for {
		select {
		case <-p.ctx.Done():
			p.logger.Warn("Producer is stopped")
			return
		case err := <-p.client.Errors():
			p.logger.Errorf("kafka producer send failed | brokers: %v | err: %v", p.c.Brokers, err)
		case <-p.client.Successes():

		}
	}
}

func NewProducer(c *ProducerConfig) (producer *Producer, err error) {
	conf := sarama.NewConfig()
	//conf.Net.KeepAlive = 60 * time.Second
	//conf.Producer.Return.Successes = false
	//conf.Producer.Flush.Frequency = time.Second
	//conf.Producer.Flush.MaxMessages = 10

	hostname := os.Getenv("HOSTNAME")
	conf.ClientID = hostname
	conf.Producer.RequiredAcks = sarama.WaitForLocal
	conf.Producer.Retry.Max = 3
	conf.Producer.Return.Successes = true
	conf.Producer.Return.Errors = true

	client, err := sarama.NewAsyncProducer(c.Brokers, conf)

	if err != nil {
		return
	}

	producer = &Producer{
		c:      c,
		client: client,
		logger: c.Logger,
		wg:     &sync.WaitGroup{},
	}

	producer.ctx, producer.cancel = context.WithCancel(context.Background())

	producer.wg.Add(1)
	go producer.logErr()

	return
}
