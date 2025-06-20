package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

type Client interface {
	Publish(ctx context.Context, msg string) error
}

type client struct {
	projectID string
	topicID   string
	pool      *pubsub.Client
}

func New(projectID string, topicID string, opts ...option.ClientOption) (Client, error) {
	ctx := context.Background()
	cli, err := pubsub.NewClient(ctx, projectID, opts...)
	if err != nil {
		return nil, fmt.Errorf("pubsub: NewClient: %w", err)
	}
	return &client{
		projectID: projectID,
		topicID:   topicID,
		pool:      cli,
	}, nil
}

func (c *client) Publish(ctx context.Context, msg string) error {
	t := c.pool.Topic(c.topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	serverID, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("pubsub: result.Get: %w", err)
	}
	fmt.Printf("pubsub: serverID: %s", serverID)
	return nil
}
