package mongo

// 新版mongo驱动，旧版mongo不支持，旧版用 lib/mgo

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

const (
	TimeOutDefault = 10 // second
)

type Config struct {
	User          string        `yaml:"user" json:"user"`
	Password      string        `yaml:"password" json:"password"`
	Hosts         []string      `yaml:"hosts" json:"hosts"`
	Database      string        `yaml:"database" json:"database"`
	ReplicaSet    string        `yaml:"replica_set" json:"replica_set"`
	DirectConnect bool          `yaml:"direct_connect" json:"direct_connect"`
	PoolMax       uint64        `yaml:"pool_max" json:"pool_max"`
	PoolMin       uint64        `yaml:"pool_min" json:"pool_min"`
	Timeout       time.Duration `yaml:"timeout" json:"timeout"`
	Url           string        `yaml:"url" json:"url"`
}

type Logger struct {
	logger *zap.SugaredLogger
}

func (l *Logger) Print(values ...interface{}) {
	fmt.Println(values)

	// if len(values) > 1 {
	// 	source := values[1].(string)
	//
	// 	if dirs := strings.Split(source, "/"); len(dirs) >= 3 {
	// 		source = strings.Join(dirs[len(dirs)-3:], "/")
	// 	}
	//
	// 	if values[0] == "sql" {
	// 		if len(values) > 5 {
	// 			sql := gorm.LogFormatter(values...)[3]
	// 			execTime := float64(values[2].(time.Duration).Nanoseconds()/1e4) / 100.0
	// 			rows := values[5].(int64)
	// 			l.logger.Debugf("query: <%s> | %.2fms | %d rows | %s", source, execTime, rows, sql)
	// 		}
	// 	} else {
	// 		l.logger.Debug(source, values[2:])
	// 	}
	// }
}

type Client struct {
	cli *mongo.Client
	c   *Config
}

type Pool struct {
	locker  sync.RWMutex
	clients map[string]*Client
	logger  *zap.SugaredLogger
}

func (p *Pool) Add(name string, c *Config) error {
	p.locker.Lock()
	defer p.locker.Unlock()

	var (
		err    error
		opts   *options.ClientOptions
		client *mongo.Client
	)

	if c.Timeout == 0 {
		c.Timeout = TimeOutDefault
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout*time.Second)
	defer cancel()

	if c.Url != "" {
		opts = options.Client().ApplyURI(c.Url)
	} else {
		opts = options.Client().
			SetHosts(c.Hosts).
			SetAuth(options.Credential{
				AuthMechanism:           "",
				AuthMechanismProperties: nil,
				AuthSource:              "",
				Username:                c.User,
				Password:                c.Password,
				PasswordSet:             true}).
			SetReplicaSet(c.ReplicaSet).
			SetDirect(c.DirectConnect).
			SetTimeout(c.Timeout * time.Second).
			SetMaxPoolSize(c.PoolMax).
			SetMinPoolSize(c.PoolMin).
			SetCompressors([]string{"snappy", "zlib", "zstd"})
	}

	client, err = mongo.Connect(ctx, opts)
	if err != nil {
		return err
	}

	p.clients[name] = &Client{
		cli: client,
		c:   c,
	}
	return nil
}

func (p *Pool) Get(name string) (cli *mongo.Client, err error) {
	defer func() {
		if e := recover(); e != nil {
			p.logger.Errorf("panic: %v", e)
			err = errors.New("get mongo failed")
			return
		}
	}()

	p.locker.RLock()
	defer p.locker.RUnlock()

	client, ok := p.clients[name]
	if ok {
		cli = client.cli
		return
	}

	return nil, errors.New("no mongo client")
}

func (p *Pool) GetDb(name string) (db *mongo.Database, err error) {
	defer func() {
		if e := recover(); e != nil {
			p.logger.Errorf("panic: %v", e)
			err = errors.New("get mongo failed")
			return
		}
	}()

	p.locker.RLock()
	defer p.locker.RUnlock()

	client, ok := p.clients[name]
	if ok {
		db = client.cli.Database(client.c.Database)
		return
	}

	return nil, errors.New("no mongo client")
}

func (p *Pool) Ping(name string) bool {
	p.locker.RLock()
	defer p.locker.RUnlock()

	client, ok := p.clients[name]
	if !ok {
		return false
	}

	if client != nil && client.cli.Ping(context.Background(), nil) == nil {
		return true
	}

	return false
}

func (p *Pool) Disconnect(name string) error {
	p.locker.RLock()
	defer p.locker.RUnlock()

	client, ok := p.clients[name]
	if !ok || client == nil {
		return nil
	}

	return client.cli.Disconnect(context.Background())
}

func NewPool(logger *zap.SugaredLogger) *Pool {
	return &Pool{
		clients: make(map[string]*Client, 64),
		logger:  logger,
	}
}
