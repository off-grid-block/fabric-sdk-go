package blockchain

import (
	caMsp "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/pkg/errors"
	//"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

// InvokeHello
func (setup *FabricSetup) RegUser(data caMsp.RegistrationRequest) (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "invoke")

	// new User information

	caClient, err := caMsp.New(setup.sdk.Context())
	enrollSecret, err := caClient.Register(&data)
	if err != nil {
		return "", errors.WithMessage(err, "Unable to register user with CA")
	}

	return string(enrollSecret), nil
}
