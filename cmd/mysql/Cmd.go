package mysql

import (
	"context"
	"time"

	"github.com/maxbet1507/l7ping/ping"
	"github.com/spf13/cobra"
)

// -
var (
	Cmd = &cobra.Command{
		Use:          "mysql",
		Short:        `Ping for MySQL server`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			return instance.Run(context.Background())
		},
	}
	instance ping.MySQL
)

func init() {
	Cmd.Flags().StringVar(&instance.DSN, "dsn", "user:password@tcp(127.0.0.1:3306)/dbname", "data source name")
	Cmd.Flags().DurationVar(&instance.Timeout, "timeout", 1*time.Second, "timeout")
	Cmd.Flags().IntVar(&instance.Retries, "retries", 0, "retries")
}
