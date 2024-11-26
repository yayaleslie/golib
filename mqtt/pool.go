package mqtt

import (
	"sync"

	"go.uber.org/zap"
)

type Pool struct {
	locker  sync.RWMutex
	clients map[string]*Client
	logger  *zap.SugaredLogger
}

func NewPool(logger *zap.SugaredLogger) *Pool {
	return &Pool{clients: make(map[string]*Client, 64), logger: logger}
}

func (p *Pool) Add(name string, c *Config, handler Handler) error {
	p.locker.Lock()
	defer p.locker.Unlock()

	client, err := NewClient(c, handler)
	if err != nil {
		return err
	}
	p.clients[name] = client

	return nil
}

func (p *Pool) Get(name string) (*Client, error) {
	p.locker.RLock()
	defer p.locker.RUnlock()

	client, ok := p.clients[name]
	if ok {
		return client, nil
	}

	return nil, ErrNoClient
}
