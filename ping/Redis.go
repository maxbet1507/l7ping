package ping

import (
	"context"
	"time"

	"github.com/gomodule/redigo/redis"
)

// Redis -
type Redis struct {
	URL     string
	Timeout time.Duration
	Retries int
}

func (s Redis) connect(ctx context.Context) error {
	conn, err := redis.DialURL(s.URL)
	if err == nil {
		err = conn.Close()
	}
	return err
}

// Run -
func (s Redis) Run(ctx context.Context) (err error) {
	return Wrap(s.connect, s.Timeout, s.Retries)(ctx)
}
