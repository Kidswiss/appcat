// Code generated by MockGen. DO NOT EDIT.
// Source: sgbackups.go

// Package mock_postgres is a generated GoMock package.
package mock_postgres

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/vshn/appcat-apiserver/apis/appcat/v1"
	internalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	watch "k8s.io/apimachinery/pkg/watch"
)

// MocksgbackupProvider is a mock of sgbackupProvider interface.
type MocksgbackupProvider struct {
	ctrl     *gomock.Controller
	recorder *MocksgbackupProviderMockRecorder
}

// MocksgbackupProviderMockRecorder is the mock recorder for MocksgbackupProvider.
type MocksgbackupProviderMockRecorder struct {
	mock *MocksgbackupProvider
}

// NewMocksgbackupProvider creates a new mock instance.
func NewMocksgbackupProvider(ctrl *gomock.Controller) *MocksgbackupProvider {
	mock := &MocksgbackupProvider{ctrl: ctrl}
	mock.recorder = &MocksgbackupProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksgbackupProvider) EXPECT() *MocksgbackupProviderMockRecorder {
	return m.recorder
}

// GetSGBackup mocks base method.
func (m *MocksgbackupProvider) GetSGBackup(ctx context.Context, name, namespace string) (*v1.SGBackupInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSGBackup", ctx, name, namespace)
	ret0, _ := ret[0].(*v1.SGBackupInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSGBackup indicates an expected call of GetSGBackup.
func (mr *MocksgbackupProviderMockRecorder) GetSGBackup(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSGBackup", reflect.TypeOf((*MocksgbackupProvider)(nil).GetSGBackup), ctx, name, namespace)
}

// ListSGBackup mocks base method.
func (m *MocksgbackupProvider) ListSGBackup(ctx context.Context, namespace string, options *internalversion.ListOptions) (*[]v1.SGBackupInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSGBackup", ctx, namespace, options)
	ret0, _ := ret[0].(*[]v1.SGBackupInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSGBackup indicates an expected call of ListSGBackup.
func (mr *MocksgbackupProviderMockRecorder) ListSGBackup(ctx, namespace, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSGBackup", reflect.TypeOf((*MocksgbackupProvider)(nil).ListSGBackup), ctx, namespace, options)
}

// WatchSGBackup mocks base method.
func (m *MocksgbackupProvider) WatchSGBackup(ctx context.Context, namespace string, options *internalversion.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchSGBackup", ctx, namespace, options)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchSGBackup indicates an expected call of WatchSGBackup.
func (mr *MocksgbackupProviderMockRecorder) WatchSGBackup(ctx, namespace, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchSGBackup", reflect.TypeOf((*MocksgbackupProvider)(nil).WatchSGBackup), ctx, namespace, options)
}
