package azservicebus

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"
	"github.com/stretchr/testify/require"
)

func TestAutoLockRenewer(t *testing.T) {
	ctx := context.TODO()

	client, cleanup, queueName := setupLiveTest(t)
	defer cleanup()

	sender, err := client.NewSender(queueName)
	require.NoError(t, err)

	defer sender.Close(ctx)

	sender.SendMessage(ctx, &Message{
		Body: []byte("lock renewal test"),
	})

	receiver, err := client.NewReceiver(ReceiverWithQueue(queueName))
	require.NoError(t, err)

	defer receiver.Close(ctx)

	messages, err := receiver.ReceiveMessages(ctx, 1)
	require.NoError(t, err)

	_, _, mgmt, _, err := receiver.amqpLinks.Get(ctx)
	require.NoError(t, err)

	// the associated-link-name is an option (for efficiency). We'll need to pass it
	// when we get to actually doing this for real.
	lockRenewer := internal.NewLockRenewer("", mgmt.RenewLocks, func(err error) {
		require.NoError(t, err)
	})

	cancel, err := lockRenewer.Renew(ctx, messages[0].RawAMQPMessage)
	require.NoError(t, err)

	defer cancel()
}
