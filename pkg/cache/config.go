package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Client struct {
	Options *redis.Options
	ctx     context.Context
}

func NewRedisConfig(Host, Port, Password string, DB int) *Client {
	return &Client{
		Options: &redis.Options{
			Addr:     Host + ":" + Port,
			Password: Password,
			DB:       DB,
		},
	}
}

func (c *Client) Connect() *redis.Client {
	return redis.NewClient(c.Options)
}
