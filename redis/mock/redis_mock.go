// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/redis/redis.go
//
// Generated by this command:
//
//	mockgen -source=internal/pkg/redis/redis.go -destination=internal/pkg/redis/mock/redis_mock.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockRedis is a mock of Redis interface.
type MockRedis struct {
	ctrl     *gomock.Controller
	recorder *MockRedisMockRecorder
}

// MockRedisMockRecorder is the mock recorder for MockRedis.
type MockRedisMockRecorder struct {
	mock *MockRedis
}

// NewMockRedis creates a new mock instance.
func NewMockRedis(ctrl *gomock.Controller) *MockRedis {
	mock := &MockRedis{ctrl: ctrl}
	mock.recorder = &MockRedisMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedis) EXPECT() *MockRedisMockRecorder {
	return m.recorder
}

// DeleteIfExist mocks base method.
func (m *MockRedis) DeleteIfExist(ctx context.Context, lockKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteIfExist", ctx, lockKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteIfExist indicates an expected call of DeleteIfExist.
func (mr *MockRedisMockRecorder) DeleteIfExist(ctx, lockKey any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIfExist", reflect.TypeOf((*MockRedis)(nil).DeleteIfExist), ctx, lockKey)
}

// Get mocks base method.
func (m *MockRedis) Get(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRedisMockRecorder) Get(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRedis)(nil).Get), ctx, key)
}

// Set mocks base method.
func (m *MockRedis) Set(ctx context.Context, key string, value any, expirationTime time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, value, expirationTime)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockRedisMockRecorder) Set(ctx, key, value, expirationTime any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRedis)(nil).Set), ctx, key, value, expirationTime)
}

// SetIfNotExist mocks base method.
func (m *MockRedis) SetIfNotExist(ctx context.Context, key string, value any, expirationTime time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetIfNotExist", ctx, key, value, expirationTime)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetIfNotExist indicates an expected call of SetIfNotExist.
func (mr *MockRedisMockRecorder) SetIfNotExist(ctx, key, value, expirationTime any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIfNotExist", reflect.TypeOf((*MockRedis)(nil).SetIfNotExist), ctx, key, value, expirationTime)
}
