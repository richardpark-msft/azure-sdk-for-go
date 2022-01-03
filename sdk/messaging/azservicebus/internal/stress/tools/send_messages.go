// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package tools

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/stress/shared"
)

func SendMessages(remainingArgs []string) int {
	log.SetFlags(0)

	fs := flag.NewFlagSet("sender", flag.ExitOnError)

	entityName := fs.String("entity", "", "Name of a queue or topic to send messages to")
	clientCreator := shared.AddAuthFlags(fs)

	if err := fs.Parse(remainingArgs); err != nil {
		log.Printf("ERROR: %s", err.Error())
		fs.PrintDefaults()
		return 1
	}

	if *entityName == "" {
		log.Printf("ERROR: must pass a queue or topic using -entity")
		fs.PrintDefaults()
		return 1
	}

	client, _, err := clientCreator()

	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		return 1
	}

	sender, err := client.NewSender(*entityName, nil)

	if err != nil {
		log.Printf("ERROR: failed to create sender: %s", err.Error())
		return 1
	}

	ctx, cancel := shared.NewCtrlCContext()
	defer cancel()

	streamingBatch, err := shared.NewStreamingMessageBatch(ctx, sender, nil)

	if err != nil {
		log.Printf("ERROR: failed to create streaming batch: %s", err.Error())
		return 1
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		err = streamingBatch.Add(ctx, &azservicebus.Message{
			Body: scanner.Bytes(),
		})

		if err != nil {
			log.Printf("Failed to send batch: %s", err.Error())
			return 1
		}
	}

	if err := streamingBatch.Close(ctx); err != nil {
		log.Printf("Failed to send final batch: %s", err.Error())
		return 1
	}

	return 0
}
