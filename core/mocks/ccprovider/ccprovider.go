/*
Copyright IBM Corp. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ccprovider

import (
	"context"

	"github.com/hyperledger/fabric/core/common/ccprovider"
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/protos/peer"
)

// MockCcProviderFactory is a factory that returns
// mock implementations of the ccprovider.ChaincodeProvider interface
type MockCcProviderFactory struct {
}

// NewChaincodeProvider returns a mock implementation of the ccprovider.ChaincodeProvider interface
func (c *MockCcProviderFactory) NewChaincodeProvider() ccprovider.ChaincodeProvider {
	return &mockCcProviderImpl{}
}

// mockCcProviderImpl is a mock implementation of the chaincode provider
type mockCcProviderImpl struct {
}

type mockCcProviderContextImpl struct {
}

// GetContext does nothing
func (c *mockCcProviderImpl) GetContext(ledger ledger.PeerLedger) (context.Context, error) {
	return nil, nil
}

// GetCCContext does nothing
func (c *mockCcProviderImpl) GetCCContext(cid, name, version, txid string, syscc bool, prop *peer.Proposal) interface{} {
	return &mockCcProviderContextImpl{}
}

// GetCCValidationInfoFromLCCC does nothing
func (c *mockCcProviderImpl) GetCCValidationInfoFromLCCC(ctxt context.Context, txid string, prop *peer.Proposal, chainID string, chaincodeID string) (string, []byte, error) {
	return "vscc", nil, nil
}

// ExecuteChaincode does nothing
func (c *mockCcProviderImpl) ExecuteChaincode(ctxt context.Context, cccid interface{}, args [][]byte) (*peer.Response, *peer.ChaincodeEvent, error) {
	return nil, nil, nil
}

// ReleaseContext does nothing
func (c *mockCcProviderImpl) ReleaseContext() {
}
