package mqtt

import (
	"errors"
	"fmt"

	pahoMqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
)

var (
	ErrNoTopic  = errors.New("no topic")
	ErrNoClient = errors.New("mqtt client no initialize")
)

type HandlerFunc func(topic string, payload []byte)

type Handler interface {
	DefaultPublish(pahoMqtt.Client, pahoMqtt.Message)
	OnConnect(pahoMqtt.Client)
	ConnectionLost(pahoMqtt.Client, error)
	Reconnecting(pahoMqtt.Client, *pahoMqtt.ClientOptions)
}

type Config struct {
	Addr     string             `toml:"addr" yaml:"addr" json:"addr"`
	Host     string             `toml:"host" yaml:"host" json:"host"`
	Port     int                `toml:"port" yaml:"port" json:"port"`
	User     string             `toml:"user" yaml:"user" json:"user"`
	Password string             `toml:"password" yaml:"password" json:"password"`
	ClientID string             `toml:"client_id" yaml:"client_id" json:"client_id"`
	Qos      byte               `toml:"qos" yaml:"qos" json:"qos"`
	Protocol uint               `toml:"protocol" yaml:"protocol" json:"protocol"`
	Logger   *zap.SugaredLogger `json:"-"`
}

func (c *Config) GetAddr() string {
	if c.Addr == "" {
		c.Addr = fmt.Sprintf("%s:%d", c.Host, c.Port)
	}

	return c.Addr
}
