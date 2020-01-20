package fire

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/thang14/footballnotify/types"
)

// Fire ...
type Fire struct {
	app *firebase.App
}

// New ...
func New() *Fire {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return &Fire{
		app: app,
	}
}

// SendMsgs ...
func (fire Fire) SendMsgs(msgs []types.Message) {
	for _, msg := range msgs {
		log.Printf("topic %v title: %s", msg.Topics, msg.Title)
		if err := fire.SendMsg(msg); err != nil {
			log.Printf("send msg err: %s", err)
			continue
		}
	}
}

// SendMsg ...
func (fire Fire) SendMsg(msg types.Message) error {

	// [START send_to_token_golang]
	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := fire.app.Messaging(ctx)
	if err != nil {
		return err
	}

	for _, topic := range msg.Topics {
		// [START send_to_topic_golang]
		// The topic name can be optionally prefixed with "/topics/".
		topic := topic

		// See documentation on defining a message payload.
		message := &messaging.Message{
			Notification: &messaging.Notification{
				Title: msg.Title,
			},
			Topic: topic,
		}

		// Send a message to the devices subscribed to the provided topic.
		response, err := client.Send(ctx, message)
		if err != nil {
			return err
		}
		// Response is a message ID string.
		fmt.Println("Successfully sent message:", response)
	}

	// [END send_to_topic_golang]
	return nil
}
