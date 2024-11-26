package cron

import (
	"context"

	cron "github.com/robfig/cron/v3"
)

type Cron struct {
	*cron.Cron
}

func New(opt ...cron.Option) *Cron {
	c := cron.New(opt...)
	return &Cron{c}
}

func (c *Cron) Start() {
	c.Cron.Start()
}

func (c *Cron) Stop() context.Context {
	return c.Cron.Stop()
}

func (c *Cron) everyStr(every string) string {
	return "@every " + every
}

func (c *Cron) EveryJob(every string, cmd cron.Job) (cron.EntryID, error) {
	return c.AddJob(c.everyStr(every), cmd)
}

func (c *Cron) Remove(id cron.EntryID) {
	if c == nil {
		return
	}
	c.Cron.Remove(id)
}
