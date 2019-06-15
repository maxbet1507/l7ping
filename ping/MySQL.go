package ping

import (
	"context"
	"database/sql/driver"
	"time"

	"github.com/go-sql-driver/mysql"
)

type devNull struct{}

func (s devNull) Print(...interface{}) {}

func init() {
	mysql.SetLogger(devNull{})
}

// MySQL -
type MySQL struct {
	DSN     string
	Timeout time.Duration
	Retries int
}

type runMySQL struct {
	dsn       string
	config    *mysql.Config
	connector driver.Connector
}

func (s *runMySQL) parseDSN(ctx context.Context) (err error) {
	s.config, err = mysql.ParseDSN(s.dsn)
	return
}

func (s *runMySQL) newConnector(ctx context.Context) (err error) {
	s.connector, err = mysql.NewConnector(s.config)
	return
}

func (s *runMySQL) connect(ctx context.Context) error {
	conn, err := s.connector.Connect(ctx)
	if err == nil {
		err = conn.Close()
	}
	return err
}

// Run -
func (s MySQL) Run(ctx context.Context) (err error) {
	run := runMySQL{
		dsn: s.DSN,
	}

	fn := []func(context.Context) error{
		run.parseDSN,
		run.newConnector,
		Wrap(run.connect, s.Timeout, s.Retries),
	}

	for i := 0; i < len(fn) && err == nil; i++ {
		err = fn[i](ctx)
	}
	return
}
