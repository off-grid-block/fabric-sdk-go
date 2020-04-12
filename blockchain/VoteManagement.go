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
func (setup *FabricSetup) InitVoteHandler(username string, PollID string, VoterID string, VoterSex string, VoterAge string, Salt string, VoteHash string) (string, error) {

	text := fmt.Sprintf(
		"{\"PollID\":\"%s\",\"VoterID\":\"%s\",\"VoterSex\":\"%s\",\"VoterAge\":%s,\"Salt\":\"%s\",\"VoteHash\":\"%s\"}",
		PollID,
		VoterID,
		VoterSex,
		VoterAge,
		Salt,
		VoteHash,
	)

	// Add data to transient map (because we are using private data, all of the data will be in the transient map)
	transientDataMap := make(map[string][]byte)
	transientDataMap["vote"] = []byte(text)

	eventID := "eventInitVote"

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
	response, err := client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: "initVote", Args: [][]byte{}, TransientMap: transientDataMap})
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

//read entry on chaincode using SDK
func (setup *FabricSetup) GetVoteSDK(username string, pollID string, voterID string) (string, error) {

	// concatenate poll ID and voter ID to get vote key
	voteKey := pollID + voterID
	clientContext := setup.sdk.ChannelContext(setup.ChannelID, fabsdk.WithUser(username))
	fmt.Println("Query Client context", clientContext)
	client, err := channel.New(clientContext)
	if err != nil {
		return "", errors.WithMessage(err, "failed to create new channel client")
	}
	// create and send request for reading an entry
	response, err := client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: "getVote", Args: [][]byte{[]byte(voteKey)}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}
