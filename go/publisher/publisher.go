package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/pubsub"
)

type Message struct {
	Origin  string `json:"origin"`
	Subject string `json:"subject"`
}

func publish(w io.Writer, projectID string, topicID string, msg Message) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub: NewClient: %v", err)
	}
	defer client.Close()

	t := client.Topic(topicID)

	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(b),
	})

	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("pubsub: result.Get: %v", err)
	}
	fmt.Fprintf(w, "Published a message; msg ID: %v\n", id)
	return nil
}

func main() {
	msg := Message{
		Origin:  "go",
		Subject: "Hello world",
	}
	if len(os.Args) > 1 && os.Args[1] != "" {
		msg.Subject = os.Args[1]
	}
	err := publish(os.Stdout, "gps-demo", "demo-topic", msg)
	if err != nil {
		fmt.Printf("ERROR: %s \n", err.Error())
	}
}
