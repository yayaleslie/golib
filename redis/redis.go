package redis

import (
	"context"
	"errors"
	"strconv"
	"sync"

	redis "github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string `toml:"host" yaml:"host" json:"host"`
	Port     int    `toml:"port" yaml:"port" json:"port"`
	Password string `toml:"password" yaml:"password" json:"password"`
	Database int    `toml:"database" yaml:"database" json:"database"`
}

func (c *Config) GetAddr() string {
	return c.Host + ":" + strconv.Itoa(c.Port)
}

type Pool struct {
	locker  sync.RWMutex
	clients map[string]*redis.Client
}

func (p *Pool) Add(name string, c *Config) {
	p.locker.Lock()
	defer p.locker.Unlock()

	p.clients[name] = redis.NewClient(&redis.Options{
		Addr:     c.GetAddr(),
		Password: c.Password,
		DB:       c.Database,
	})
}

func (p *Pool) Get(name string) (*redis.Client, error) {
	p.locker.RLock()
	defer p.locker.RUnlock()

	client, ok := p.clients[name]

	if ok {
		return client, nil
	}

	return nil, errors.New("no redis client")
}

func (p *Pool) Ping(ctx context.Context, name string) bool {
	client, ok := p.clients[name]

	if !ok {
		return false
	}

	if client != nil && client.Get(ctx, "PING").String() == "PONG" {
		return true
	}

	return false
}

func NewPool() *Pool {
	return &Pool{clients: make(map[string]*redis.Client, 64)}
}
