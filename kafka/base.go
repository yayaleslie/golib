package kafka

import (
	"encoding/json"
	"log"
)

type logger interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Panic(...interface{})
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Panicf(string, ...interface{})
}

type Config struct {
	Producer *ProducerConfig `toml:"producer" yaml:"producer" json:"producer"`
	Consumer *ConsumerConfig `toml:"consumer" yaml:"consumer" json:"consumer"`
}

func (c *Config) P() *ProducerConfig {
	if c == nil {
		return nil
	}

	return c.Producer
}

func (c *Config) C() *ConsumerConfig {
	if c == nil {
		return nil
	}
	c.Consumer.SetTopics()
	return c.Consumer
}

func toString(v interface{}) string {
	b, e := json.Marshal(v)
	if e != nil {
		log.Println(e)
	}
	return string(b)
}
