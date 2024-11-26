package mqtt

import (
	pahoMqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/yayaleslie/golib/uuid"
	"go.uber.org/zap"
)

type Client struct {
	c       *Config
	client  *pahoMqtt.Client
	logger  *zap.SugaredLogger
	handler Handler
}

// NewClient 初始化
func NewClient(c *Config, handler Handler) (*Client, error) {
	addr := c.GetAddr()
	opts := pahoMqtt.NewClientOptions()
	opts.AddBroker("tcp://" + addr)
	opts.SetUsername(c.User)
	opts.SetPassword(c.Password)
	//opts.SetKeepAlive(120 * time.Second)
	//opts.SetPingTimeout(60 * time.Second)

	mqttUuid := uuid.GenerateUuid(addr, 1)
	opts.SetProtocolVersion(c.Protocol)
	opts.SetAutoReconnect(true)
	opts.SetClientID(opts.Username + "_" + mqttUuid)
	opts.SetCleanSession(true)

	if handler != nil {
		if handler.DefaultPublish != nil {
			opts.SetDefaultPublishHandler(handler.DefaultPublish)
		}
		if handler.OnConnect != nil {
			opts.SetOnConnectHandler(handler.OnConnect)
		}
		if handler.ConnectionLost != nil {
			opts.SetConnectionLostHandler(handler.ConnectionLost)
		}
		if handler.Reconnecting != nil {
			opts.SetReconnectingHandler(handler.Reconnecting)
		}
	}

	newClient := pahoMqtt.NewClient(opts)
	if token := newClient.Connect(); token.Wait() && token.Error() != nil {
		c.Logger.Errorf("[Mqtt] mqtt connect to %s(%s:%s) failed | client_id: %s | err: %s", addr, opts.Username, opts.Password, opts.ClientID, token.Error().Error())
		return nil, token.Error()
	}

	c.Logger.Infof("[Mqtt] mqtt connect to %s(%s:%s) success | client_id: %s", addr, opts.Username, opts.Password, opts.ClientID)
	return &Client{c: c, client: &newClient, logger: c.Logger}, nil
}

// Publish 发布
func (c *Client) Publish(topic string, qos byte, data interface{}) error {
	if topic == "" {
		return ErrNoTopic
	}

	if *c.client == nil {
		return ErrNoClient
	}

	token := (*c.client).Publish(topic, qos, false, data)
	token.Wait()
	if err := token.Error(); err != nil {
		return err
	}

	return nil
}

// Subscribe 订阅
func (c *Client) Subscribe(topic string, qos byte, callback HandlerFunc) error {
	if topic == "" {
		return ErrNoTopic
	}

	if *c.client == nil {
		return ErrNoClient
	}
	_ = c.UnsubscribeTopic(topic)

	token := (*c.client).Subscribe(topic, qos, func(client pahoMqtt.Client, message pahoMqtt.Message) {
		if callback != nil {
			callback(message.Topic(), message.Payload())
		}
	})

	token.Wait()
	if err := token.Error(); err != nil {
		return err
	}

	return nil
}

// SubscribeTopics 批量订阅
func (c *Client) SubscribeTopics(topics map[string]byte, callback HandlerFunc) error {
	if len(topics) == 0 {
		return ErrNoTopic
	}

	if *c.client == nil {
		return ErrNoClient
	}

	var topicSlice []string
	for topic := range topics {
		topicSlice = append(topicSlice, topic)
	}
	_ = c.UnsubscribeTopic(topicSlice...)

	token := (*c.client).SubscribeMultiple(topics, func(client pahoMqtt.Client, message pahoMqtt.Message) {
		if callback != nil {
			callback(message.Topic(), message.Payload())
		}
	})

	token.Wait()
	if err := token.Error(); err != nil {
		return err
	}
	return nil
}

// UnsubscribeTopic 取消订阅
func (c *Client) UnsubscribeTopic(topics ...string) error {
	if len(topics) == 0 {
		return ErrNoTopic
	}

	if *c.client == nil {
		return ErrNoClient
	}

	token := (*c.client).Unsubscribe(topics...)
	token.Wait()
	if err := token.Error(); err != nil {
		c.logger.Errorf("[SubscribeTopics] subscribeErr | topics: %+v | err: %s", topics, err)
		return err
	}

	return nil
}

// DisConnect 断开连接
func (c *Client) DisConnect(quiesce uint) error {
	if *c.client == nil {
		return ErrNoClient
	}

	(*c.client).Disconnect(quiesce)
	return nil
}

// Qos 获取默认qos
func (c *Client) Qos() byte {
	if c.c == nil {
		return 0
	}
	return c.c.Qos
}
