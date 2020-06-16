package client

import (
	"context"
	"time"

	"github.com/pkg/errors"
	pb "github.com/webmalc/namecheap-dns/protos/changer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const timeout = 30

// Client is the server client
type Client struct {
	logger   logger
	config   *config
	ipGetter ipGetter
}

// getConnection set up a connection to the server.
func (c *Client) getConnection() *grpc.ClientConn {
	creds, err := credentials.NewClientTLSFromFile(c.config.crtPath, "")
	if err != nil {
		c.logger.Error(errors.Wrap(err, "client"))
	}
	conn, err := grpc.Dial(
		c.config.address, grpc.WithTransportCredentials(creds), grpc.WithBlock(),
	)
	if err != nil {
		c.logger.Error(errors.Wrap(err, "client"))
	}
	return conn
}

// Request makes the request
func (c *Client) Request() {
	conn := c.getConnection()
	defer conn.Close()
	cl := pb.NewChangerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*timeout)
	defer cancel()
	ip, err := c.ipGetter.GetIP()
	if err != nil {
		c.logger.Error(errors.Wrap(err, "client"))
		panic(err)
	}
	r, err := cl.Change(ctx, &pb.ChangeRequest{Ip: ip.String()})
	if err != nil {
		c.logger.Error(errors.Wrap(err, "client"))
	}
	c.logger.Infof("Got the response from the grpc server: %+v", r)
}

// NewClient creates a new client instance
func NewClient(l logger, i ipGetter) *Client {
	return &Client{logger: l, config: newConfig(), ipGetter: i}
}
