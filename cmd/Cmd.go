package cmd

import (
	"github.com/maxbet1507/l7ping/cmd/mysql"
	"github.com/maxbet1507/l7ping/cmd/redis"
	"github.com/spf13/cobra"
)

// -
var (
	Cmd = &cobra.Command{
		Use: "l7ping",
	}
)

func init() {
	Cmd.AddCommand(mysql.Cmd)
	Cmd.AddCommand(redis.Cmd)
}
