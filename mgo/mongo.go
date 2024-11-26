package mgo

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
	mgo "gopkg.in/mgo.v2"
)

const (
	PoolLimitDefault = 20
	TimeOutDefault   = 10 // second
)

type Config struct {
	User          string   `yaml:"user" json:"user"`
	Password      string   `yaml:"password" json:"password"`
	Hosts         []string `yaml:"hosts" json:"hosts"`
	Database      string   `yaml:"database" json:"database"`
	ReplicaSet    string   `yaml:"replica_set" json:"replica_set"`
	DirectConnect bool     `yaml:"direct_connect" json:"direct_connect"`
	PoolLimit     int      `yaml:"pool_limit" json:"pool_limit"`
	Timeout       int64    `yaml:"timeout" json:"timeout"`
	Url           string   `yaml:"url" json:"url"`
}

type Logger struct {
	logger *zap.SugaredLogger
}

func (l *Logger) Print(values ...interface{}) {
	fmt.Println("==================== mongo =========================")
	fmt.Println(values)
	fmt.Println("==================== mongo =========================")

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

type Pool struct {
	locker   sync.RWMutex
	clients  map[string]*mgo.Database
	dialInfo map[string]*mgo.DialInfo
	logger   *zap.SugaredLogger
}

func (p *Pool) Add(name string, c *Config) error {
	p.locker.Lock()
	defer p.locker.Unlock()

	var (
		err      error
		session  *mgo.Session
		dialInfo *mgo.DialInfo
	)

	if c.Timeout == 0 {
		c.Timeout = TimeOutDefault
	}

	if c.PoolLimit == 0 {
		c.PoolLimit = PoolLimitDefault
	}

	if c.Url != "" {
		if dialInfo, err = mgo.ParseURL(c.Url); err != nil {
			return err
		}
	} else {
		dialInfo = &mgo.DialInfo{
			Addrs:          c.Hosts,
			Timeout:        time.Duration(c.Timeout) * time.Second,
			ReplicaSetName: c.ReplicaSet,
			Direct:         c.DirectConnect,
			Database:       c.Database,
			Username:       c.User,
			Password:       c.Password,
			PoolLimit:      c.PoolLimit,
		}
	}

	if session, err = mgo.DialWithInfo(dialInfo); err != nil {
		return err
	}

	curSession := session.Copy()
	curSession.SetMode(mgo.Monotonic, true)
	p.clients[name] = curSession.DB(c.Database)
	p.dialInfo[name] = dialInfo

	return nil
}

func (p *Pool) GetDb(name string) (db *mgo.Database, err error) {
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
		db = client.Session.Copy().DB(p.dialInfo[name].Database)
		return
	}

	return nil, errors.New("no mongo client")
}

func (p *Pool) Ping(name string) bool {
	client, ok := p.clients[name]

	if !ok {
		return false
	}

	if client != nil && client.Session.Ping() == nil {
		return true
	}

	return false
}

func NewPool(logger *zap.SugaredLogger) *Pool {
	return &Pool{clients: make(map[string]*mgo.Database, 64), dialInfo: make(map[string]*mgo.DialInfo, 64), logger: logger}
}
