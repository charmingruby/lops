package postgres

import (
	"context"

	"github.com/charmingruby/pack/pkg/telemetry/logger"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

const driver = "postgres"

type Client struct {
	Conn *sqlx.DB

	log *logger.Logger
}

func New(log *logger.Logger, url string) (*Client, error) {
	db, err := sqlx.Connect(driver, url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Client{Conn: db, log: log}, nil
}

func (c *Client) Ping(ctx context.Context) error {
	return c.Conn.PingContext(ctx)
}

func (c *Client) Close() error {
	if err := c.Conn.Close(); err != nil {
		c.log.Error("failed to close postgres connection", "error", err)
		return err
	}

	c.log.Info("postgres connection closed")

	return nil
}
