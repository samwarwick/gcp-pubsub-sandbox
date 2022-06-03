package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync/atomic"
	"time"

	"cloud.google.com/go/pubsub"
)

type Message struct {
	Origin  string `json:"origin"`
	Subject string `json:"subject"`
}

func pullMsgs(w io.Writer, projectID, subID string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	sub := client.Subscription(subID)

	ctx, cancel := context.WithTimeout(ctx, 300*time.Second)
	defer cancel()

	var received int32
	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		messageData := string(msg.Data)
		fmt.Fprintf(w, "Got message (ID=%s): %q\n", msg.ID, messageData)
		var pubsubMessage Message
		json.Unmarshal([]byte(messageData), &pubsubMessage)
		fmt.Println("Subject: " + pubsubMessage.Subject)

		atomic.AddInt32(&received, 1)
		msg.Ack()
		if pubsubMessage.Subject == "quit" {
			fmt.Println("quiting")
			cancel()
		}
	})
	if err != nil {
		return fmt.Errorf("sub.Receive: %v", err)
	}
	fmt.Fprintf(w, "Received %d messages\n", received)

	return nil
}

func main() {
	fmt.Println("Go subscriber")
	err := pullMsgs(os.Stdout, "gps-demo", "demo-sub")
	if err != nil {
		fmt.Printf("ERROR: %s \n", err.Error())
	}
}
