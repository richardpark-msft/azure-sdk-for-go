// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"context"
	"crypto/tls"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/amqpwrap"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/test"
	"github.com/Azure/go-amqp"
	"github.com/stretchr/testify/require"
)

func TestAMQPTransactions(t *testing.T) {
	entityPath := "testqueue"
	keyLogPath := os.Getenv("SSLKEYLOGFILE")

	t.Logf("Writing keylogfile to %s", keyLogPath)
	writer, err := os.OpenFile(keyLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	require.NoError(t, err)

	defer writer.Close()

	cs := test.GetConnectionString(t)
	ns, err := NewNamespace(NamespaceWithConnectionString(cs), NamespaceWithTLSConfig(&tls.Config{
		KeyLogWriter: writer,
	}))
	require.NoError(t, err)

	defer func() { _ = ns.Close(false) }()

	links := NewAMQPLinks(NewAMQPLinksArgs{
		NS:         ns,
		EntityPath: entityPath,
		CreateLinkFunc: func(ctx context.Context, session amqpwrap.AMQPSession) (amqpwrap.AMQPSenderCloser, amqpwrap.AMQPReceiverCloser, error) {
			return newLinksForAMQPLinksTest(test.BuiltInTestQueue, session)
		},
		GetRecoveryKindFunc: GetRecoveryKind,
	})

	defer test.RequireLinksClose(t, links)

	// NOTE: If the control link is closed while there exist non-discharged transactions it created, then all
	// such transactions are immediately rolled back, and attempts to perform further transactional work
	// on them will lead to failure.

	// NOTE:  Control link is outside the boundary of an entity, that is, same control link can be used to
	// initiate and discharge transactions for multiple entities.

	lwid, err := links.Get(context.TODO())
	require.NoError(t, err)

	func(ctx context.Context) {
		session, _, err := ns.NewAMQPSession(ctx)
		require.NoError(t, err)

		transactionController, err := session.NewSender(ctx, "", &amqp.SenderOptions{
			TransactionController: true,
		})
		require.NoError(t, err)

		// http://docs.oasis-open.org/amqp/core/v1.0/os/amqp-core-transactions-v1.0-os.html#type-declare
		id, err := transactionController.Declare(ctx, amqp.TransactionDeclare{}, nil)
		require.NoError(t, err)

		err = lwid.Sender.Send(context.Background(), &amqp.Message{
			TransactionID: id,
			Value:         "hello",
		}, nil)
		require.NoError(t, err)

		// this comes back as a string, and I'll need to use it later when discharging the transaction later as well.
		require.NotEmpty(t, id)

		err = transactionController.Discharge(ctx, amqp.TransactionDischarge{
			TransactionID: id,
			Fail:          false,
			//Fail: true,
		}, nil)
		require.NoError(t, err)
	}(context.Background())
}
