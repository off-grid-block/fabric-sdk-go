/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package txn

import (
	reqContext "context"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"

	"github.com/off-grid-block/fabric-protos-go/common"
	pb "github.com/off-grid-block/fabric-protos-go/peer"
	"github.com/off-grid-block/fabric-sdk-go/internal/github.com/hyperledger/fabric/protoutil"
	"github.com/off-grid-block/fabric-sdk-go/pkg/common/errors/multi"
	contextApi "github.com/off-grid-block/fabric-sdk-go/pkg/common/providers/context"
	"github.com/off-grid-block/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/off-grid-block/fabric-sdk-go/pkg/context"
)

// CreateChaincodeInvokeProposal creates a proposal for transaction.
func CreateChaincodeInvokeProposal(txh fab.TransactionHeader, request fab.ChaincodeInvokeRequest, indyflag bool) (*fab.TransactionProposal, error) {
	if request.ChaincodeID == "" {
		return nil, errors.New("ChaincodeID is required")
	}

	if request.Fcn == "" {
		return nil, errors.New("Fcn is required")
	}

	// Add function name to arguments
	argsArray := make([][]byte, len(request.Args)+1)
	argsArray[0] = []byte(request.Fcn)
	for i, arg := range request.Args {
		argsArray[i+1] = arg
	}

	// create invocation spec to target a chaincode with arguments
	ccis := &pb.ChaincodeInvocationSpec{ChaincodeSpec: &pb.ChaincodeSpec{
		Type: pb.ChaincodeSpec_GOLANG, ChaincodeId: &pb.ChaincodeID{Name: request.ChaincodeID},
		Input: &pb.ChaincodeInput{Args: argsArray}}}

	proposal, txid, err := protoutil.CreateChaincodeProposalWithTxIDNonceAndTransient(string(txh.TransactionID()), common.HeaderType_ENDORSER_TRANSACTION, txh.ChannelID(), ccis, txh.Nonce(), txh.Creator(), request.TransientMap, indyflag, txh.Did())
	if err != nil {
		return nil, errors.Wrap(err, "failed to create chaincode proposal")
	}

	tp := fab.TransactionProposal{
		TxnID:    fab.TransactionID(txid), //TxnID:    txh.TransactionID(),
		Proposal: proposal,
	}
	return &tp, nil
}

// signProposal creates a SignedProposal based on the current context.
func signProposal(ctx contextApi.Client, proposal *pb.Proposal, indyFlag bool, did string) (*pb.SignedProposal, error) {
	proposalBytes, err := proto.Marshal(proposal)
	if err != nil {
		return nil, errors.Wrap(err, "mashal proposal failed")
	}

	signingMgr := ctx.SigningManager()
	if signingMgr == nil {
		return nil, errors.New("signing manager is nil")
	}
    // Creating signature using aca-py agent
	signature, err := signingMgr.Sign(proposalBytes, ctx.PrivateKey(), indyFlag, did)
	if err != nil {
		return nil, errors.WithMessage(err, "sign failed")
	}

	return &pb.SignedProposal{ProposalBytes: proposalBytes, Signature: signature}, nil
}

// SendProposal sends a TransactionProposal to ProposalProcessor.
func SendProposal(reqCtx reqContext.Context, proposal *fab.TransactionProposal, targets []fab.ProposalProcessor, indyFlag bool, did string) ([]*fab.TransactionProposalResponse, error) {

	if proposal == nil {
		return nil, errors.New("proposal is required")
	}
	if len(targets) < 1 {
		return nil, errors.New("targets is required")
	}

	for _, p := range targets {
		if p == nil {
			return nil, errors.New("target is nil")
		}
	}

	targets = getTargetsWithoutDuplicates(targets)

	ctx, ok := context.RequestClientContext(reqCtx)
	if !ok {
		return nil, errors.New("failed get client context from reqContext for signProposal")
	}
	signedProposal, err := signProposal(ctx, proposal.Proposal, indyFlag, did)
	if err != nil {
		return nil, errors.WithMessage(err, "sign proposal failed")
	}
	request := fab.ProcessProposalRequest{SignedProposal: signedProposal}

	var responseMtx sync.Mutex
	var transactionProposalResponses []*fab.TransactionProposalResponse
	var wg sync.WaitGroup
	errs := multi.Errors{}
	count := 0
	for _, p := range targets {
		wg.Add(1)
		count = count + 1
		go func(processor fab.ProposalProcessor) {
			defer wg.Done()

			// TODO: The RPC should be timed-out.
			//resp, err := processor.ProcessTransactionProposal(context.NewRequestOLD(ctx), request)
			resp, err := processor.ProcessTransactionProposal(reqCtx, request)
			if err != nil {
				logger.Debugf("Received error response from txn proposal processing: %s", err)
				responseMtx.Lock()
				errs = append(errs, err)
				responseMtx.Unlock()
				return
			}

			responseMtx.Lock()
			transactionProposalResponses = append(transactionProposalResponses, resp)
			responseMtx.Unlock()
		}(p)
	}
	wg.Wait()
	// fmt.Println("Proposal Resposnse")
	// fmt.Println(transactionProposalResponses)

	return transactionProposalResponses, errs.ToError()
}

// getTargetsWithoutDuplicates returns a list of targets without duplicates
func getTargetsWithoutDuplicates(targets []fab.ProposalProcessor) []fab.ProposalProcessor {
	peerUrlsToTargets := map[string]fab.ProposalProcessor{}
	var uniqueTargets []fab.ProposalProcessor

	for i := range targets {
		peer, ok := targets[i].(fab.Peer)
		if !ok {
			// ProposalProcessor is not a fab.Peer... cannot remove duplicates
			return targets
		}
		if _, present := peerUrlsToTargets[peer.URL()]; !present {
			uniqueTargets = append(uniqueTargets, targets[i])
			peerUrlsToTargets[peer.URL()] = targets[i]
		}
	}

	if len(uniqueTargets) != len(targets) {
		logger.Warn("Duplicate target peers in configuration")
	}

	return uniqueTargets
}
