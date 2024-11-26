package redis

import (
	"context"
	"log"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type PubSub struct {
	client      *redis.Client
	pubSub      *redis.PubSub
	channelSize int
	isClose     bool
	done        chan struct{}
}

func (ps *PubSub) Publish(ctx context.Context, channel string, message interface{}) error {
	err := ps.client.Publish(ctx, channel, message).Err()
	if err != nil {
		log.Printf("redis publish failed | err: %s", err.Error())
		return err
	}
	return nil
}

func (ps *PubSub) Subscribe(ctx context.Context, handler func(payload []byte), channels ...string) {
	pubSub := ps.client.Subscribe(ctx, channels...)
	ps.pubSub = pubSub

	for {
		select {
		case msg := <-pubSub.Channel(redis.WithChannelSize(ps.channelSize)):
			if ps.isClose {
				return
			}

			if handler != nil {
				handler([]byte(msg.Payload))
			}

		case <-ps.done:
			return

		case <-time.After(60 * time.Second):
			if ps == nil || ps.pubSub == nil || ps.isClose {
				return
			}
			if err := ps.pubSub.Ping(ctx); err != nil {
				log.Printf("redis: PubSub ping error  | error: %s", err.Error())
				return
			}
		}
	}
}

func (ps *PubSub) Unsubscribe(ctx context.Context, channels ...string) error {
	ps.Close()
	return ps.pubSub.Unsubscribe(ctx, channels...)
}

func (ps *PubSub) Close() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("error: PubSub close failed | err: %v", err)
		}
	}()

	ps.isClose = true
	close(ps.done)
	if ps.pubSub != nil {
		ps.pubSub.Close()
	}
}

func (ps *PubSub) IsClose() bool {
	return ps.isClose
}

func (ps *PubSub) SetChannelSize(size int) *PubSub {
	ps.channelSize = size
	return ps
}

func (ps *PubSub) SetClient(client *redis.Client) *PubSub {
	ps.client = client
	return ps
}

func (ps *PubSub) GetClient() *redis.Client {
	return ps.client
}

func NewPubSub(client *redis.Client) *PubSub {
	if client == nil {
		return nil
	}

	return &PubSub{client: client, channelSize: 100, done: make(chan struct{})}
}
