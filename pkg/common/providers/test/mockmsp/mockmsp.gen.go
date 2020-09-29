// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/off-grid-block/fabric-sdk-go/pkg/common/providers/msp (interfaces: IdentityConfig,IdentityManager,Providers)

// Package mockmsp is a generated GoMock package.
package mockmsp

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	msp "github.com/off-grid-block/fabric-sdk-go/pkg/common/providers/msp"
)

// MockIdentityConfig is a mock of IdentityConfig interface
type MockIdentityConfig struct {
	ctrl     *gomock.Controller
	recorder *MockIdentityConfigMockRecorder
}

// MockIdentityConfigMockRecorder is the mock recorder for MockIdentityConfig
type MockIdentityConfigMockRecorder struct {
	mock *MockIdentityConfig
}

// NewMockIdentityConfig creates a new mock instance
func NewMockIdentityConfig(ctrl *gomock.Controller) *MockIdentityConfig {
	mock := &MockIdentityConfig{ctrl: ctrl}
	mock.recorder = &MockIdentityConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIdentityConfig) EXPECT() *MockIdentityConfigMockRecorder {
	return m.recorder
}

// CAClientCert mocks base method
func (m *MockIdentityConfig) CAClientCert(arg0 string) ([]byte, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CAClientCert", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CAClientCert indicates an expected call of CAClientCert
func (mr *MockIdentityConfigMockRecorder) CAClientCert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CAClientCert", reflect.TypeOf((*MockIdentityConfig)(nil).CAClientCert), arg0)
}

// CAClientKey mocks base method
func (m *MockIdentityConfig) CAClientKey(arg0 string) ([]byte, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CAClientKey", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CAClientKey indicates an expected call of CAClientKey
func (mr *MockIdentityConfigMockRecorder) CAClientKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CAClientKey", reflect.TypeOf((*MockIdentityConfig)(nil).CAClientKey), arg0)
}

// CAConfig mocks base method
func (m *MockIdentityConfig) CAConfig(arg0 string) (*msp.CAConfig, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CAConfig", arg0)
	ret0, _ := ret[0].(*msp.CAConfig)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CAConfig indicates an expected call of CAConfig
func (mr *MockIdentityConfigMockRecorder) CAConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CAConfig", reflect.TypeOf((*MockIdentityConfig)(nil).CAConfig), arg0)
}

// CAKeyStorePath mocks base method
func (m *MockIdentityConfig) CAKeyStorePath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CAKeyStorePath")
	ret0, _ := ret[0].(string)
	return ret0
}

// CAKeyStorePath indicates an expected call of CAKeyStorePath
func (mr *MockIdentityConfigMockRecorder) CAKeyStorePath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CAKeyStorePath", reflect.TypeOf((*MockIdentityConfig)(nil).CAKeyStorePath))
}

// CAServerCerts mocks base method
func (m *MockIdentityConfig) CAServerCerts(arg0 string) ([][]byte, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CAServerCerts", arg0)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CAServerCerts indicates an expected call of CAServerCerts
func (mr *MockIdentityConfigMockRecorder) CAServerCerts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CAServerCerts", reflect.TypeOf((*MockIdentityConfig)(nil).CAServerCerts), arg0)
}

// Client mocks base method
func (m *MockIdentityConfig) Client() *msp.ClientConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(*msp.ClientConfig)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockIdentityConfigMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockIdentityConfig)(nil).Client))
}

// CredentialStorePath mocks base method
func (m *MockIdentityConfig) CredentialStorePath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CredentialStorePath")
	ret0, _ := ret[0].(string)
	return ret0
}

// CredentialStorePath indicates an expected call of CredentialStorePath
func (mr *MockIdentityConfigMockRecorder) CredentialStorePath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredentialStorePath", reflect.TypeOf((*MockIdentityConfig)(nil).CredentialStorePath))
}

// MockIdentityManager is a mock of IdentityManager interface
type MockIdentityManager struct {
	ctrl     *gomock.Controller
	recorder *MockIdentityManagerMockRecorder
}

// MockIdentityManagerMockRecorder is the mock recorder for MockIdentityManager
type MockIdentityManagerMockRecorder struct {
	mock *MockIdentityManager
}

// NewMockIdentityManager creates a new mock instance
func NewMockIdentityManager(ctrl *gomock.Controller) *MockIdentityManager {
	mock := &MockIdentityManager{ctrl: ctrl}
	mock.recorder = &MockIdentityManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIdentityManager) EXPECT() *MockIdentityManagerMockRecorder {
	return m.recorder
}

// CreateSigningIdentity mocks base method
func (m *MockIdentityManager) CreateSigningIdentity(arg0 ...msp.SigningIdentityOption) (msp.SigningIdentity, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateSigningIdentity", varargs...)
	ret0, _ := ret[0].(msp.SigningIdentity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSigningIdentity indicates an expected call of CreateSigningIdentity
func (mr *MockIdentityManagerMockRecorder) CreateSigningIdentity(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSigningIdentity", reflect.TypeOf((*MockIdentityManager)(nil).CreateSigningIdentity), arg0...)
}

// GetSigningIdentity mocks base method
func (m *MockIdentityManager) GetSigningIdentity(arg0 string) (msp.SigningIdentity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSigningIdentity", arg0)
	ret0, _ := ret[0].(msp.SigningIdentity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSigningIdentity indicates an expected call of GetSigningIdentity
func (mr *MockIdentityManagerMockRecorder) GetSigningIdentity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSigningIdentity", reflect.TypeOf((*MockIdentityManager)(nil).GetSigningIdentity), arg0)
}

// MockProviders is a mock of Providers interface
type MockProviders struct {
	ctrl     *gomock.Controller
	recorder *MockProvidersMockRecorder
}

// MockProvidersMockRecorder is the mock recorder for MockProviders
type MockProvidersMockRecorder struct {
	mock *MockProviders
}

// NewMockProviders creates a new mock instance
func NewMockProviders(ctrl *gomock.Controller) *MockProviders {
	mock := &MockProviders{ctrl: ctrl}
	mock.recorder = &MockProvidersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProviders) EXPECT() *MockProvidersMockRecorder {
	return m.recorder
}

// IdentityConfig mocks base method
func (m *MockProviders) IdentityConfig() msp.IdentityConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IdentityConfig")
	ret0, _ := ret[0].(msp.IdentityConfig)
	return ret0
}

// IdentityConfig indicates an expected call of IdentityConfig
func (mr *MockProvidersMockRecorder) IdentityConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IdentityConfig", reflect.TypeOf((*MockProviders)(nil).IdentityConfig))
}

// IdentityManager mocks base method
func (m *MockProviders) IdentityManager(arg0 string) (msp.IdentityManager, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IdentityManager", arg0)
	ret0, _ := ret[0].(msp.IdentityManager)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// IdentityManager indicates an expected call of IdentityManager
func (mr *MockProvidersMockRecorder) IdentityManager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IdentityManager", reflect.TypeOf((*MockProviders)(nil).IdentityManager), arg0)
}

// UserStore mocks base method
func (m *MockProviders) UserStore() msp.UserStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserStore")
	ret0, _ := ret[0].(msp.UserStore)
	return ret0
}

// UserStore indicates an expected call of UserStore
func (mr *MockProvidersMockRecorder) UserStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserStore", reflect.TypeOf((*MockProviders)(nil).UserStore))
}
