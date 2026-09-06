package feereceiver

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/SapphireDAOO/contract-api/proto/feereceiverpb"
)

const callTimeout = 25 * time.Second

type Client struct {
	conn   *grpc.ClientConn
	client pb.FeeReceiverClient
	target string
}

func NewClient(target string) (*Client, error) {
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("fee receiver at %s: %w", target, err)
	}

	return &Client{conn: conn, client: pb.NewFeeReceiverClient(conn), target: target}, nil
}

func (c *Client) Generate(ctx context.Context, in *pb.PrepareAddressRequest) (*pb.PrepareAddressResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, callTimeout)
	defer cancel()

	return c.client.Generateaddresses(ctx, in)
}

func (c *Client) Authorize(ctx context.Context, in *pb.VerifyAddressesRequest) (*pb.VerifyAddressesResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, callTimeout)
	defer cancel()

	return c.client.Send(ctx, in)
}

func (c *Client) Target() string { return c.target }

func (c *Client) Close() error { return c.conn.Close() }
