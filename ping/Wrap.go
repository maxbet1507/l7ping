package ping

import (
	"context"
	"time"
)

// Wrap -
func Wrap(fn func(context.Context) error, timeout time.Duration, retries int) func(context.Context) error {
	fnwt := func(ctx context.Context, last bool) (err error) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		if err = fn(ctx); err != nil && !last {
			<-ctx.Done()
		}
		return
	}

	r := func(ctx context.Context) (err error) {
		for n, cont := 0, true; cont; n++ {
			cont = (n < retries || retries < 0)
			select {
			case <-ctx.Done():
				err = ctx.Err()
				cont = false

			default:
				err = fnwt(ctx, !cont)
				cont = cont && (err != nil)
			}
		}
		return
	}
	return r
}
