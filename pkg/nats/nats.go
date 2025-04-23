package nats

import (
	"context"
	"github.com/nats-io/nats.go"
)

type NatsClient struct {
	conn *nats.Conn
}

func New(ctx context.Context, hosts []string, nkey string, isTest bool) (*NatsClient, error) {
	opts, err := setOptions(ctx, hosts, nkey, isTest)
	if err != nil {
		panic(err)
	}

	conn, err := opts.Connect()
	if err != nil {
		panic(err)
	}

	return &NatsClient{conn: conn}, nil
}

func (nc *NatsClient) Close() {
	nc.conn.Close()
}
