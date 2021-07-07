// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/kt/channel/types.go

// Package channel is a generated GoMock package.
package channel

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockChannel is a mock of Channel interface.
type MockChannel struct {
	ctrl     *gomock.Controller
	recorder *MockChannelMockRecorder
}

// MockChannelMockRecorder is the mock recorder for MockChannel.
type MockChannelMockRecorder struct {
	mock *MockChannel
}

// NewMockChannel creates a new mock instance.
func NewMockChannel(ctrl *gomock.Controller) *MockChannel {
	mock := &MockChannel{ctrl: ctrl}
	mock.recorder = &MockChannelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannel) EXPECT() *MockChannelMockRecorder {
	return m.recorder
}

// ForwardRemoteToLocal mocks base method.
func (m *MockChannel) ForwardRemoteToLocal(certificate *Certificate, sshAddress, remoteEndpoint, localEndpoint string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForwardRemoteToLocal", certificate, sshAddress, remoteEndpoint, localEndpoint)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForwardRemoteToLocal indicates an expected call of ForwardRemoteToLocal.
func (mr *MockChannelMockRecorder) ForwardRemoteToLocal(certificate, sshAddress, remoteEndpoint, localEndpoint interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForwardRemoteToLocal", reflect.TypeOf((*MockChannel)(nil).ForwardRemoteToLocal), certificate, sshAddress, remoteEndpoint, localEndpoint)
}

// StartSocks5Proxy mocks base method.
func (m *MockChannel) StartSocks5Proxy(certificate *Certificate, sshAddress, socks5Address string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartSocks5Proxy", certificate, sshAddress, socks5Address)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartSocks5Proxy indicates an expected call of StartSocks5Proxy.
func (mr *MockChannelMockRecorder) StartSocks5Proxy(certificate, sshAddress, socks5Address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartSocks5Proxy", reflect.TypeOf((*MockChannel)(nil).StartSocks5Proxy), certificate, sshAddress, socks5Address)
}
