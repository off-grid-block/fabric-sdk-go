package blockchain

import (
	"fmt"
	"time"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/pkg/errors"
)

// Add a new user
func (setup *FabricSetup) Adduser(username string, userID string, userAddress string, userPhotoLocation string, createdBy string, createdDate string, updatedBy string, updatedDate string) (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "AddUserDetails")
	args = append(args, userID)
	args = append(args, userAddress)
	args = append(args, userPhotoLocation)
	args = append(args, createdBy)
	args = append(args, createdDate)
	args = append(args, updatedBy)
	args = append(args, updatedDate)

	eventID := "eventAddUser"

	// Add data that will be visible in the proposal, like a description of the invoke request
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in hello invoke")
	// client context created
	clientContext := setup.sdk.ChannelContext(setup.ChannelID, fabsdk.WithUser(username))
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
	response, err := client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3]), []byte(args[4]), []byte(args[5]), []byte(args[6]), []byte(args[7])}, TransientMap: transientDataMap})
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

// update a new user
func (setup *FabricSetup) Updateuser(username string, userID string, userAddress string, userPhotoLocation string, createdBy string, createdDate string, updatedBy string, updatedDate string) (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "UpdateUserDetails")
	args = append(args, userID)
	args = append(args, userAddress)
	args = append(args, userPhotoLocation)
	args = append(args, createdBy)
	args = append(args, createdDate)
	args = append(args, updatedBy)
	args = append(args, updatedDate)

	eventID := "eventUpdateUser"

	// Add data that will be visible in the proposal, like a description of the invoke request
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in hello invoke")
	// client context created
	clientContext := setup.sdk.ChannelContext(setup.ChannelID, fabsdk.WithUser(username))
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
	response, err := client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3]), []byte(args[4]), []byte(args[5]), []byte(args[6]), []byte(args[7])}, TransientMap: transientDataMap})
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

//User Query
func (setup *FabricSetup) UserQuery(username string, key string) (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, key)
	fun := "QueryUserDetails"

	//Channel client is used to query and execute transactions
	clientContext := setup.sdk.ChannelContext(setup.ChannelID, fabsdk.WithUser(username))
	fmt.Println("Query Client context", clientContext)
	client, err := channel.New(clientContext)
	if err != nil {
		return "", errors.WithMessage(err, "failed to create new channel client")
	}
	response, err := client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: fun, Args: [][]byte{[]byte(args[0])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}
