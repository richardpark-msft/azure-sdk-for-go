package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/joho/godotenv"
)

func purge(queueName string, client *azservicebus.Client) {
	fmt.Printf("[BEGIN] Purging messages\n")
	defer fmt.Printf("[END] Purging messages\n")
	receiver, err := client.NewReceiverForQueue(queueName, &azservicebus.ReceiverOptions{
		ReceiveMode: azservicebus.ReceiveAndDelete,
	})

	if err != nil {
		panic(err)
	}

	for {
		messages, err := receiver.ReceiveMessages(context.Background(), 100, &azservicebus.ReceiveOptions{
			MaxWaitTime: 5 * time.Second,
		})

		if err != nil {
			panic(err)
		}

		if len(messages) == 0 {
			break
		}
	}
}

func main() {
	_ = godotenv.Load()

	queueName := "samples"

	writer, err := os.Create("C:\\dev\\temp\\gotlskey.txt")

	if err != nil {
		panic(err)
	}

	defer writer.Close()

	client, err := azservicebus.NewClientFromConnectionString(os.Getenv("SERVICEBUS_CONNECTION_STRING"), &azservicebus.ClientOptions{
		TLSConfig: &tls.Config{
			KeyLogWriter: writer,
		},
	})

	if err != nil {
		panic(err)
	}

	purge(queueName, client)

	sender, err := client.NewSender(queueName)

	if err != nil {
		panic(err)
	}

	receiver, err := client.NewReceiverForQueue(queueName, nil)

	if err != nil {
		panic(err)
	}

	_, amqpReceiver, _, _, err := receiver.AmqpLinks.Get(context.Background())

	if err != nil {
		panic(err)
	}

	total := 0

	for {
		log.Printf("Starting to receive messages (so far: %d)", total)

		// no messages here yet, so we'll issue credit only to drain it.
		if err := amqpReceiver.IssueCredit(10); err != nil {
			panic(err)
		}

		if err := amqpReceiver.DrainCredit(context.Background()); err != nil {
			panic(err)
		}

		err := sender.SendMessage(context.Background(), &azservicebus.Message{
			Body: []byte("hello world"),
		})

		if err != nil {
			panic(err)
		}

		if err := amqpReceiver.IssueCredit(10); err != nil {
			panic(err)
		}

		// there are plenty of messages to receive at this point
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		message, err := amqpReceiver.Receive(ctx)
		cancel()

		if err != nil {
			panic(err)
		}

		if message == nil {
			fmt.Printf("No message received!")
		}

		total++
	}
}

func sendMessages(sender *azservicebus.Sender) {
	const numBatches = 10
	const batchSize = 100

	fmt.Printf("[BEGIN] sendMessages(%d)\n", numBatches*batchSize)

	buffer := [1024]byte{}

	for i := 0; i < numBatches; i++ {
		batch, err := sender.NewMessageBatch(context.Background(), nil)

		if err != nil {
			panic(err)
		}

		for j := 0; j < batchSize; j++ {
			added, err := batch.Add(&azservicebus.Message{
				Body: buffer[:],
			})

			if !added || err != nil {
				panic(err)
			}
		}

		if err := sender.SendMessageBatch(context.Background(), batch); err != nil {
			panic(err)
		}
	}

	fmt.Printf("[END] sendMessages\n")
}
