package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hapoon/go-pubsub/pkg/pubsub"
)

func main() {
	fmt.Println("start go-pubsub")
	ps, err := pubsub.New("local", "local1")
	if err != nil {
		log.Fatalf("pubsub.New: %v", err)
		return
	}
	ctx := context.Background()
	if err = ps.Publish(ctx, "hello"); err != nil {
		log.Fatalf("pubsub.Publish: %v", err)
		return
	}
}
