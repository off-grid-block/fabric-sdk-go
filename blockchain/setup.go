package blockchain

import (
	"fmt"

	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	caMsp "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/pkg/errors"
)

// FabricSetup implementation
type FabricSetup struct {
	ConfigFile      string
	OrgID           string
	OrdererID       string
	ChannelID       string
	ChainCodeID     string
	initialized     bool
	ChannelConfig   string
	ChaincodeGoPath string
	ChaincodePath   string
	OrgAdmin        string
	OrgName         string
	UserName        string
	CaClient        *caMsp.Client
	client          *channel.Client
	admin           *resmgmt.Client
	sdk             *fabsdk.FabricSDK
	event           *event.Client
}

// Initialize reads the configuration file and sets up the client, chain and event hub
func (setup *FabricSetup) Initialize() error {

	// Add parameters for the initialization
	if setup.initialized {
		return errors.New("sdk already initialized")
	}

	// Initialize the SDK with the configuration file
	sdk, err := fabsdk.New(config.FromFile(setup.ConfigFile))
	if err != nil {
		return errors.WithMessage(err, "failed to create SDK")
	}
	setup.sdk = sdk
	fmt.Println("SDK created")

	// The resource management client is responsible for managing channels (create/update channel)
	resourceManagerClientContext := setup.sdk.Context(fabsdk.WithUser(setup.OrgAdmin), fabsdk.WithOrg(setup.OrgName))
	if err != nil {
		return errors.WithMessage(err, "failed to load Admin identity")
	}
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return errors.WithMessage(err, "failed to create channel management client from Admin identity")
	}
	setup.admin = resMgmtClient
	fmt.Println("Ressource management client created")

	// The MSP client allow us to retrieve user information from their identity, like its signing identity which we will need to save the channel
	mspClient, err := mspclient.New(sdk.Context(), mspclient.WithOrg(setup.OrgName))
	if err != nil {
		return errors.WithMessage(err, "failed to create MSP client")
	}

	adminIdentity, err := mspClient.GetSigningIdentity(setup.OrgAdmin)
	if err != nil {
		return errors.WithMessage(err, "failed to get admin signing identity")
	}
	req := resmgmt.SaveChannelRequest{ChannelID: setup.ChannelID, ChannelConfigPath: setup.ChannelConfig, SigningIdentities: []msp.SigningIdentity{adminIdentity}}
	txID, err := setup.admin.SaveChannel(req, resmgmt.WithOrdererEndpoint(setup.OrdererID))
	if err != nil || txID.TransactionID == "" {
		return errors.WithMessage(err, "failed to save channel")
	}
	fmt.Println("Channel created")

	// Make admin user join the previously created channel
	if err = setup.admin.JoinChannel(setup.ChannelID, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(setup.OrdererID)); err != nil {
		return errors.WithMessage(err, "failed to make admin join channel")
	}
	fmt.Println("Channel joined")

	fmt.Println("Initialization Successful")
	setup.initialized = true
	return nil
}

// Create collection config to for chaincode instantiation
func newCollectionConfig(colName, policy string, reqPeerCount, maxPeerCount int32, blockToLive uint64) (*cb.CollectionConfig, error) {
	p, err := cauthdsl.FromString(policy)
	if err != nil {
		return nil, err
	}
	cpc := &cb.CollectionPolicyConfig{
		Payload: &cb.CollectionPolicyConfig_SignaturePolicy{
			SignaturePolicy: p,
		},
	}
	return &cb.CollectionConfig{
		Payload: &cb.CollectionConfig_StaticCollectionConfig{
			StaticCollectionConfig: &cb.StaticCollectionConfig{
				Name:              colName,
				MemberOrgsPolicy:  cpc,
				RequiredPeerCount: reqPeerCount,
				MaximumPeerCount:  maxPeerCount,
				BlockToLive:       blockToLive,
			},
		},
	}, nil
}

func (setup *FabricSetup) InstallAndInstantiateCC() error {

	// Create the chaincode package that will be sent to the peers
	ccPkg, err := packager.NewCCPackage(setup.ChaincodePath, setup.ChaincodeGoPath)
	if err != nil {
		return errors.WithMessage(err, "failed to create chaincode package")
	}
	fmt.Println("ccPkg created")

	// Install example cc to org peers
	installCCReq := resmgmt.InstallCCRequest{Name: setup.ChainCodeID, Path: setup.ChaincodePath, Version: "0", Package: ccPkg}
	_, err = setup.admin.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return errors.WithMessage(err, "failed to install chaincode")
	}
	fmt.Println("Chaincode installed")

	// Set up chaincode policy
	ccPolicy := cauthdsl.SignedByAnyMember([]string{"org1.hf.sample.io"})

	//resp, err := setup.admin.InstantiateCC(setup.ChannelID, resmgmt.InstantiateCCRequest{Name: setup.ChainCodeID, Path: setup.ChaincodeGoPath, Version: "0", Args: [][]byte{[]byte("init")}, Policy: ccPolicy})

	// Create collection config 1 for collectionVote
	var collCfg1RequiredPeerCount, collCfg1MaximumPeerCount int32
	var collCfg1BlockToLive uint64

	collCfg1Name := "collectionVote"
	collCfg1BlockToLive = 1000000
	collCfg1RequiredPeerCount = 0
	collCfg1MaximumPeerCount = 2
	collCfg1Policy := "OR('org1.hf.sample.io.member')"

	collCfg1, err := newCollectionConfig(collCfg1Name, collCfg1Policy, collCfg1RequiredPeerCount, collCfg1MaximumPeerCount, collCfg1BlockToLive)
	if err != nil {
		return errors.WithMessage(err, "failed to create collection config 1")
	}

	// Create collection config 1 for collectionVote
	var collCfg2RequiredPeerCount, collCfg2MaximumPeerCount int32
	var collCfg2BlockToLive uint64

	collCfg2Name := "collectionVotePrivateDetails"
	collCfg2BlockToLive = 1000000
	collCfg2RequiredPeerCount = 0
	collCfg2MaximumPeerCount = 2
	collCfg2Policy := "OR('org1.hf.sample.io.member')"

	collCfg2, err := newCollectionConfig(collCfg2Name, collCfg2Policy, collCfg2RequiredPeerCount, collCfg2MaximumPeerCount, collCfg2BlockToLive)
	if err != nil {
		return errors.WithMessage(err, "failed to create collection config 1")
	}

	cfg := []*cb.CollectionConfig{collCfg1, collCfg2}

	// instantiate chaincode with cc policy and collection configs
	resp, err := setup.admin.InstantiateCC(
		// Channel ID
		setup.ChannelID,
		// InstantiateCCRequest struct
		resmgmt.InstantiateCCRequest{
			Name:       setup.ChainCodeID,
			Path:       setup.ChaincodeGoPath,
			Version:    "0",
			Args:       [][]byte{[]byte("init")},
			Policy:     ccPolicy,
			CollConfig: cfg,
		},
		// options
		resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil || resp.TransactionID == "" {
		return errors.WithMessage(err, "failed to instantiate the chaincode")
	}
	fmt.Println("Chaincode instantiated")

	// Channel client is used to query and execute transactions
	// clientContext := setup.sdk.ChannelContext(setup.ChannelID, fabsdk.WithUser(setup.UserName))
	// setup.client, err = channel.New(clientContext)
	// if err != nil {
	// 	return errors.WithMessage(err, "failed to create new channel client")
	// }
	// fmt.Println("Channel client created")

	// // Creation of the client which will enables access to our channel events
	// setup.event, err = event.New(clientContext)
	// if err != nil {
	// 	return errors.WithMessage(err, "failed to create new event client")
	// }
	// fmt.Println("Event client created")

	fmt.Println("Chaincode Installation & Instantiation Successful")

	return nil
}

func (setup *FabricSetup) CloseSDK() {
	setup.sdk.Close()
}
