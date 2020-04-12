package main

import (
	"fmt"
	"os"

	"github.com/gosdk-example/blockchain"
	"github.com/gosdk-example/web"
	"github.com/gosdk-example/web/controllers"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Network parameters
		OrdererID: "orderer.hf.sample.io",

		// Channel parameters
		ChannelID:     "sample",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/gosdk-example/fixtures/artifacts/sample.channel.tx",

		// Chaincode parameters
		ChainCodeID:     "gosdk-example",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/gosdk-example/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "org1",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		return
	}
	// Close SDK
	defer fSetup.CloseSDK()

	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC()
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
		return
	}

	// Launch the web application listening
	app := &controllers.Application{
		Fabric: &fSetup,
	}
	web.Serve(app)
}
