package gorm

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

const (
	DriverMysql      = "mysql"
	DriverClickhouse = "clickhouse"
)

type Config struct {
	Host            string `yaml:"host" toml:"host" json:"host"`
	Port            int    `yaml:"port" toml:"port" json:"port"`
	User            string `yaml:"user" toml:"user" json:"user"`
	Password        string `yaml:"password" toml:"password" json:"password"`
	Charset         string `yaml:"charset" toml:"charset" json:"charset"`
	Database        string `yaml:"database" toml:"database" json:"database"`
	Timeout         int    `yaml:"timeout" toml:"timeout" json:"timeout"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime" toml:"conn_max_lifetime" json:"conn_max_lifetime"`
	MaxOpenConns    int    `yaml:"max_open_conns" toml:"max_open_conns" json:"max_open_conns"`
	MaxIdleConns    int    `yaml:"max_idle_conns" toml:"max_idle_conns" json:"max_idle_conns"`
	Debug           bool   `yaml:"debug" toml:"debug" json:"debug"`
	Driver          string `yaml:"driver" toml:"driver" json:"driver"`
}

func (c *Config) GetDsn() string {
	switch c.Driver {
	case DriverClickhouse:
		return fmt.Sprintf("tcp://%s:%d?username=%s&password=%s&debug=%t",
			c.Host, c.Port, c.User, c.Password, false)

	}

	if c.Timeout <= 0 {
		c.Timeout = 3
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&timeout=%ds",
		c.User, c.Password, c.Host, c.Port, c.Database, c.Charset, c.Timeout)
}

type Pool struct {
	locker  sync.RWMutex
	clients map[string]*gorm.DB
	logger  *zap.Logger
}

func (p *Pool) Add(name string, c *Config) error {
	p.locker.Lock()
	defer p.locker.Unlock()

	orm, err := New(c, p.logger)
	if err != nil {
		return err
	}

	p.clients[name] = orm
	return nil
}

func (p *Pool) Get(name string) (*gorm.DB, error) {
	p.locker.RLock()
	defer p.locker.RUnlock()

	client, ok := p.clients[name]
	if ok {
		return client, nil
	}

	return nil, errors.New("no gorm client")
}

func (p *Pool) Ping(name string) bool {
	client, ok := p.clients[name]
	if !ok || client == nil {
		return false
	}

	db, err := client.DB()
	if err != nil {
		p.logger.Error("gorm ping failed", zap.Error(err))
		return false
	}

	if db.Ping() == nil {
		return true
	}

	return false
}

func NewPool(logger *zap.Logger) *Pool {
	return &Pool{clients: make(map[string]*gorm.DB, 64), logger: logger}
}

func New(c *Config, logger *zap.Logger) (*gorm.DB, error) {
	if logger == nil {
		logger = zap.L()
	}

	var logLevel gormLogger.LogLevel
	if c.Debug {
		logLevel = gormLogger.Info
	} else {
		logLevel = gormLogger.Error
	}

	var (
		dsn  = c.GetDsn()
		dial gorm.Dialector
	)
	switch c.Driver {
	case DriverMysql:
		dial = mysql.Open(dsn)
	case DriverClickhouse:
		dial = clickhouse.Open(dsn)
	default: // 默认mysql
		dial = mysql.Open(dsn)
	}

	orm, err := gorm.Open(dial, &gorm.Config{Logger: NewLogger(logger, logLevel)})
	if err != nil {
		return nil, err
	}

	db, err := orm.DB()
	if err != nil {
		return nil, err
	}

	if c.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(time.Duration(c.ConnMaxLifetime) * time.Second)
	}

	if c.MaxIdleConns > 0 {
		db.SetMaxIdleConns(c.MaxIdleConns)
	}

	if c.MaxOpenConns > 0 {
		db.SetMaxOpenConns(c.MaxOpenConns)
	}

	if c.Debug {
		orm.Logger.LogMode(gormLogger.Info)
	}

	return orm, nil
}

func Ping(client *gorm.DB) bool {
	if client == nil {
		return false
	}

	db, err := client.DB()
	if err != nil {
		log.Printf("[ping] gorm get db failed | err: %s", err.Error())
		return false
	}

	if err = db.Ping(); err != nil {
		log.Printf("[ping] gorm ping failed | err: %s", err.Error())
		return false
	}

	return true
}
