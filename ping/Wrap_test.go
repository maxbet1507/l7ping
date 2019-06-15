package ping_test

import (
	"context"
	"errors"
	"math"
	"testing"
	"time"

	"github.com/maxbet1507/l7ping/ping"
)

func TestWrap(t *testing.T) {

	patterns := map[string]struct {
		fn      func(context.Context) error
		timeout time.Duration
		retries int
		newctx  func() (context.Context, context.CancelFunc)
		success bool
		delta   time.Duration
	}{
		"default": {
			fn: func(context.Context) error {
				return nil
			},
			timeout: 100 * time.Millisecond,
			retries: 1,
			newctx: func() (context.Context, context.CancelFunc) {
				return context.WithCancel(context.Background())
			},
			success: true,
			delta:   10 * time.Millisecond,
		},
		"retries": {
			fn: func(context.Context) error {
				return errors.New("")
			},
			timeout: 100 * time.Millisecond,
			retries: 1,
			newctx: func() (context.Context, context.CancelFunc) {
				return context.WithCancel(context.Background())
			},
			success: false,
			delta:   100 * time.Millisecond,
		},
		"timeout": {
			fn: func(context.Context) error {
				return errors.New("")
			},
			timeout: 100 * time.Millisecond,
			retries: 1,
			newctx: func() (context.Context, context.CancelFunc) {
				return context.WithTimeout(context.Background(), 10*time.Millisecond)
			},
			success: false,
			delta:   10 * time.Millisecond,
		},
	}

	for n, p := range patterns {
		fn := ping.Wrap(p.fn, p.timeout, p.retries)
		t0 := time.Now()

		ctx, cancel := p.newctx()
		err := fn(ctx)
		cancel()

		t1 := time.Now()
		td := t1.Sub(t0)

		if (err == nil) != p.success || math.Abs(td.Seconds()-p.delta.Seconds()) > 0.01 {
			t.Fatal(n, err, td)
		}
	}
}
