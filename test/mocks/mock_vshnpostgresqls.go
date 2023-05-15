// Code generated by MockGen. DO NOT EDIT.
// Source: ../pkg/apiserver/vshn/postgres/vshnpostgresql.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/vshn/appcat-apiserver/apis/vshn/v1"
)

// MockvshnPostgresqlProvider is a mock of vshnPostgresqlProvider interface.
type MockvshnPostgresqlProvider struct {
	ctrl     *gomock.Controller
	recorder *MockvshnPostgresqlProviderMockRecorder
}

// MockvshnPostgresqlProviderMockRecorder is the mock recorder for MockvshnPostgresqlProvider.
type MockvshnPostgresqlProviderMockRecorder struct {
	mock *MockvshnPostgresqlProvider
}

// NewMockvshnPostgresqlProvider creates a new mock instance.
func NewMockvshnPostgresqlProvider(ctrl *gomock.Controller) *MockvshnPostgresqlProvider {
	mock := &MockvshnPostgresqlProvider{ctrl: ctrl}
	mock.recorder = &MockvshnPostgresqlProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockvshnPostgresqlProvider) EXPECT() *MockvshnPostgresqlProviderMockRecorder {
	return m.recorder
}

// ListXVSHNPostgreSQL mocks base method.
func (m *MockvshnPostgresqlProvider) ListXVSHNPostgreSQL(ctx context.Context, namespace string) (*v1.XVSHNPostgreSQLList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListXVSHNPostgreSQL", ctx, namespace)
	ret0, _ := ret[0].(*v1.XVSHNPostgreSQLList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListXVSHNPostgreSQL indicates an expected call of ListXVSHNPostgreSQL.
func (mr *MockvshnPostgresqlProviderMockRecorder) ListXVSHNPostgreSQL(ctx, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListXVSHNPostgreSQL", reflect.TypeOf((*MockvshnPostgresqlProvider)(nil).ListXVSHNPostgreSQL), ctx, namespace)
}
