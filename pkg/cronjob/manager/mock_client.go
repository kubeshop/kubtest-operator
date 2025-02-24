// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeshop/testkube-operator/pkg/cronjob/manager (interfaces: Interface)

// Package manager is a generated GoMock package.
package manager

import (
	context "context"
	reflect "reflect"

	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// CleanForNewArchitecture mocks base method.
func (m *MockInterface) CleanForNewArchitecture(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanForNewArchitecture", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanForNewArchitecture indicates an expected call of CleanForNewArchitecture.
func (mr *MockInterfaceMockRecorder) CleanForNewArchitecture(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanForNewArchitecture", reflect.TypeOf((*MockInterface)(nil).CleanForNewArchitecture), arg0)
}

// IsNamespaceForNewArchitecture mocks base method.
func (m *MockInterface) IsNamespaceForNewArchitecture(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNamespaceForNewArchitecture", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsNamespaceForNewArchitecture indicates an expected call of IsNamespaceForNewArchitecture.
func (mr *MockInterfaceMockRecorder) IsNamespaceForNewArchitecture(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNamespaceForNewArchitecture", reflect.TypeOf((*MockInterface)(nil).IsNamespaceForNewArchitecture), arg0, arg1)
}

// Reconcile mocks base method.
func (m *MockInterface) Reconcile(arg0 context.Context, arg1 logr.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reconcile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reconcile indicates an expected call of Reconcile.
func (mr *MockInterfaceMockRecorder) Reconcile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconcile", reflect.TypeOf((*MockInterface)(nil).Reconcile), arg0, arg1)
}
