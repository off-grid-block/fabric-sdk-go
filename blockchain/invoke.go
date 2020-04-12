package blockchain

import (
	"fmt"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/pkg/errors"
)

// InvokeHello
func (setup *FabricSetup) Invoke(fcn string, key string, value string) (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "invoke")
	args = append(args, fcn)
	args = append(args, key)
	args = append(args, value)

	eventID := "eventInvoke"

	// Add data that will be visible in the proposal, like a description of the invoke request
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in hello invoke")
	// client context created
	clientContext := setup.sdk.ChannelContext(setup.ChannelID, fabsdk.WithUser("Voting"))
	client, err := channel.New(clientContext)
	if err != nil {
		return "client error", errors.WithMessage(err, "failed to create new channel client")
	}
	//Event Creation
	event, err := event.New(clientContext)
	if err != nil {
		return "", errors.WithMessage(err, "failed to create new event client")
	}
	fmt.Println("Event client created")

	reg, notifier, err := event.RegisterChaincodeEvent(setup.ChainCodeID, eventID)
	if err != nil {
		return "", err
	}
	defer event.Unregister(reg)

	// Create a request (proposal) and send it
	response, err := client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3])}, TransientMap: transientDataMap})
	if err != nil {
		return "", fmt.Errorf("failed to move funds: %v", err)
	}

	// Wait for the result of the submission
	select {
	case ccEvent := <-notifier:
		fmt.Printf("Received CC 2 event: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return "", fmt.Errorf("did NOT receive CC event for eventId(%s)", eventID)
	}

	return string(response.TransactionID), nil
}
