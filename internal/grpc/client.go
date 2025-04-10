package grpc

import (
	"context"
	"fmt"
	ssov1 "github.com/bwjson/kolesa_proto/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"time"
)

type Client struct {
	api ssov1.AuthClient
	log *slog.Logger
}

func New(ctx context.Context, log *slog.Logger, addr string, timeout time.Duration, retriesCount int) (*Client, error) {
	cc, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	grpcClient := ssov1.NewAuthClient(cc)

	return &Client{api: grpcClient}, nil
}

func (c *Client) SendVerificationCode(phoneNumber string) error {
	_, err := c.api.SendVerificationCode(context.Background(), &ssov1.SendVerificationCodeRequest{PhoneNumber: phoneNumber})
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (c *Client) VerifyCode(phoneNumber, code string) (accessToken, refreshToken string, err error) {
	tokens, err := c.api.VerifyCode(context.Background(), &ssov1.VerifyCodeRequest{
		PhoneNumber:      phoneNumber,
		VerificationCode: code,
	})

	if err != nil {
		return "", "", fmt.Errorf("%w", err)
	}

	return tokens.AccessToken, tokens.RefreshToken, nil
}
