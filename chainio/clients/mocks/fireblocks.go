// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Layr-Labs/eigensdk-go/chainio/clients/fireblocks (interfaces: FireblocksClient)
//
// Generated by this command:
//
//	mockgen -destination=./chainio/clients/mocks/fireblocks.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/fireblocks FireblocksClient
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	fireblocks "github.com/Layr-Labs/eigensdk-go/chainio/clients/fireblocks"
	gomock "go.uber.org/mock/gomock"
)

// MockFireblocksClient is a mock of FireblocksClient interface.
type MockFireblocksClient struct {
	ctrl     *gomock.Controller
	recorder *MockFireblocksClientMockRecorder
}

// MockFireblocksClientMockRecorder is the mock recorder for MockFireblocksClient.
type MockFireblocksClientMockRecorder struct {
	mock *MockFireblocksClient
}

// NewMockFireblocksClient creates a new mock instance.
func NewMockFireblocksClient(ctrl *gomock.Controller) *MockFireblocksClient {
	mock := &MockFireblocksClient{ctrl: ctrl}
	mock.recorder = &MockFireblocksClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFireblocksClient) EXPECT() *MockFireblocksClientMockRecorder {
	return m.recorder
}

// ContractCall mocks base method.
func (m *MockFireblocksClient) ContractCall(arg0 context.Context, arg1 *fireblocks.ContractCallRequest) (*fireblocks.ContractCallResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContractCall", arg0, arg1)
	ret0, _ := ret[0].(*fireblocks.ContractCallResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContractCall indicates an expected call of ContractCall.
func (mr *MockFireblocksClientMockRecorder) ContractCall(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContractCall", reflect.TypeOf((*MockFireblocksClient)(nil).ContractCall), arg0, arg1)
}

// ListContracts mocks base method.
func (m *MockFireblocksClient) ListContracts(arg0 context.Context) ([]fireblocks.WhitelistedContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListContracts", arg0)
	ret0, _ := ret[0].([]fireblocks.WhitelistedContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContracts indicates an expected call of ListContracts.
func (mr *MockFireblocksClientMockRecorder) ListContracts(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContracts", reflect.TypeOf((*MockFireblocksClient)(nil).ListContracts), arg0)
}
