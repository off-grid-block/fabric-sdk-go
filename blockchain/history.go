package blockchain

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func (setup *FabricSetup) GettxbyID(username string,id string) (*peer.ProcessedTransaction, error) {
	clientContext := setup.sdk.ChannelContext(setup.ChannelID, fabsdk.WithUser(username))
	ledgerclient, err := ledger.New(clientContext)
	if err != nil {
		fmt.Println("error creating ledger instance")
	}
	t, err := ledgerclient.QueryTransaction(fab.TransactionID(id))
	return t, nil
}
func (setup *FabricSetup) GetBlockbytxID(id string) (*common.Block, error) {
	clientContext := setup.sdk.ChannelContext(setup.ChannelID, fabsdk.WithUser(setup.UserName))
	ledgerclient, err := ledger.New(clientContext)
	if err != nil {
		fmt.Println("error creating ledger instance")
	}
	t, err := ledgerclient.QueryBlockByTxID(fab.TransactionID(id))
	return t, nil
}
func (setup *FabricSetup) GetBlockbyHash(id string) (*common.Block, error) {
	clientContext := setup.sdk.ChannelContext(setup.ChannelID, fabsdk.WithUser(setup.UserName))
	ledgerclient, err := ledger.New(clientContext)
	if err != nil {
		fmt.Println("error creating ledger instance")
	}
	t, err := ledgerclient.QueryBlockByHash([]byte(id))
	return t, nil
}
func (setup *FabricSetup) GetBlockbyID(id string) (*common.Block, error) {
	clientContext := setup.sdk.ChannelContext(setup.ChannelID, fabsdk.WithUser(setup.UserName))
	ledgerclient, err := ledger.New(clientContext)
	if err != nil {
		fmt.Println("error creating ledger instance")
	}
	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	t, err := ledgerclient.QueryBlock(u64)
	return t, nil
}
