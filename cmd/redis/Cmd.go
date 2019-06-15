package redis

import (
	"context"
	"time"

	"github.com/maxbet1507/l7ping/ping"
	"github.com/spf13/cobra"
)

// -
var (
	Cmd = &cobra.Command{
		Use:          "redis",
		Short:        `Ping for Redis server`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			return instance.Run(context.Background())
		},
	}
	instance ping.Redis
)

func init() {
	Cmd.Flags().StringVar(&instance.URL, "url", "redis://user:password@127.0.0.1:6379/0", "redis url")
	Cmd.Flags().DurationVar(&instance.Timeout, "timeout", 1*time.Second, "timeout")
	Cmd.Flags().IntVar(&instance.Retries, "retries", 0, "retries")
}
