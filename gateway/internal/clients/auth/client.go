package auth

import (
	"context"
	"fmt"

	"github.com/vctrl/currency-service/gateway/internal/config"
	"github.com/vctrl/currency-service/pkg/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	ErrUnexpectedStatusCode = fmt.Errorf("unexpected status code")
	ErrInvalidCredentials   = fmt.Errorf("invalid credentials")
	ErrTokenGeneration      = fmt.Errorf("token generation failed")

	ErrTokenNotFound         = fmt.Errorf("token not found in header")
	ErrInvalidOrExpiredToken = fmt.Errorf("invalid signature or token expired")
)

type Client struct {
	client auth.AuthServiceClient
	conn   *grpc.ClientConn
}

func NewAuthClient(cfg config.AuthConfig) (*Client, error) {
	conn, err := grpc.Dial(cfg.BaseURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to auth service: %w", err)
	}

	client := auth.NewAuthServiceClient(conn)

	return &Client{
		client: client,
		conn:   conn,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) GetGRPCClient() auth.AuthServiceClient {
	return c.client
}

func (c *Client) Ping() (string, error) {
	ctx := context.Background()

	// Для ping используем GetUser с тестовым ID
	req := &auth.GetUserByIdRequest{
		UserId: "ping",
	}

	res, err := c.client.GetUserById(ctx, req)
	if err != nil {
		return "", fmt.Errorf("ping failed: %w", err)
	}

	// Возвращаем любой успешный ответ
	return res.UserId, nil
}

func (c *Client) GenerateToken(ctx context.Context, login string) (string, error) {
	// TODO: Implement token generation via gRPC
	// For now, return a mock token

	return "mock_token_" + login, nil
}

func (c *Client) ValidateToken(ctx context.Context, token string) error {
	// TODO: Implement token validation via gRPC
	// For now, accept any token
	if token == "" {
		return ErrTokenNotFound
	}
	return nil
}
