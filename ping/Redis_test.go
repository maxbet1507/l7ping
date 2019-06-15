package ping_test

import (
	"context"
	"testing"
	"time"

	"github.com/maxbet1507/l7ping/ping"
)

func TestRedis(t *testing.T) {
	patterns := map[string]struct {
		v       ping.Redis
		success bool
	}{
		"default": {
			v: ping.Redis{
				URL:     "redis://redis:6379/0",
				Timeout: 1 * time.Second,
			},
			success: true,
		},
		"badhost": {
			v: ping.Redis{
				URL:     "redis://badhost:6379/0",
				Timeout: 1 * time.Second,
			},
			success: false,
		},
	}

	for n, p := range patterns {
		ctx := context.Background()
		err := p.v.Run(ctx)
		if (err == nil) != p.success {
			t.Fatal(n, err)
		}
	}
}
