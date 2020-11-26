/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package signingmgr

import (
	"fmt"

	"github.com/off-grid-block/controller"

	"github.com/off-grid-block/fabric-sdk-go/pkg/common/providers/core"

	"github.com/off-grid-block/fabric-sdk-go/pkg/core/cryptosuite"
	"github.com/pkg/errors"
)

// SigningManager is used for signing objects with private key
type SigningManager struct {
	cryptoProvider core.CryptoSuite
	hashOpts       core.HashOpts
	signerOpts     core.SignerOpts
}

// New Constructor for a signing manager.
// @param {BCCSP} cryptoProvider - crypto provider
// @param {Config} config - configuration provider
// @returns {SigningManager} new signing manager
func New(cryptoProvider core.CryptoSuite) (*SigningManager, error) {
	return &SigningManager{cryptoProvider: cryptoProvider, hashOpts: cryptosuite.GetSHAOpts()}, nil
}

// Sign will sign the given object using provided key
func (mgr *SigningManager) Sign(object []byte, key core.Key, indyFlag bool, did string) ([]byte, error) {

	if len(object) == 0 {
		return nil, errors.New("object (to sign) required")
	}

	digest, err := mgr.cryptoProvider.Hash(object, mgr.hashOpts)
	if err != nil {
		return nil, err
	}

	// Signature for transaction
	var signature []byte

	// Calling the package for Signing using Indy
	if indyFlag == true {
		// create new client controller
		cc, _ := controller.NewClientController()
		cc.SigningDid = did
		// sign digest and return signature
		sigResp, err := cc.SignMessage(digest)
		if err != nil {
			return nil, fmt.Errorf("Error while signing message: %v\n", err)
		}
		// new signature
		signature = []byte(sigResp)

		//fmt.Println("Indy sign flag on")
		//indySig, _ := sigindy.IndySign(digest, did)
		//sig := fmt.Sprintf("%v", indySig["signature"])
		//fmt.Println("SIGNATURE INSIDE: ", sig)
		//signature = []byte(sig)
	}

	if indyFlag == false {

		if key == nil {
			return nil, errors.New("key (for signing) required")
		}
		signature, err = mgr.cryptoProvider.Sign(key, digest, mgr.signerOpts)
		if err != nil {
			return nil, err
		}
	}
	return signature, nil
}
