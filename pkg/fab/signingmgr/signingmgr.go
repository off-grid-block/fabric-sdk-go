/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package signingmgr

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/fab/sigindy"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/cryptosuite"
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
	var signature []byte

	// Calling the package for Signing using Indy
	if indyFlag == true {
		indySig, _ := sigindy.IndySign(digest, did)
		sig := fmt.Sprintf("%v", indySig["signature"])
		signature = []byte(sig)
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
