package ping_test

import (
	"context"
	"testing"
	"time"

	"github.com/maxbet1507/l7ping/ping"
)

func TestMySQL(t *testing.T) {
	patterns := map[string]struct {
		v       ping.MySQL
		success bool
	}{
		"default": {
			v: ping.MySQL{
				DSN:     "user:password@tcp(mysql:3306)/dbname",
				Timeout: 1 * time.Second,
			},
			success: true,
		},
		"badhost": {
			v: ping.MySQL{
				DSN:     "user:password@tcp(badhost:3306)/dbname",
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
