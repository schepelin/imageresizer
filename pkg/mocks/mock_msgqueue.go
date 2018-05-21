// Code generated by MockGen. DO NOT EDIT.
// Source: msgqueue.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPublisher is a mock of Publisher interface
type MockPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockPublisherMockRecorder
}

// MockPublisherMockRecorder is the mock recorder for MockPublisher
type MockPublisherMockRecorder struct {
	mock *MockPublisher
}

// NewMockPublisher creates a new mock instance
func NewMockPublisher(ctrl *gomock.Controller) *MockPublisher {
	mock := &MockPublisher{ctrl: ctrl}
	mock.recorder = &MockPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPublisher) EXPECT() *MockPublisherMockRecorder {
	return m.recorder
}

// PublishResizeJob mocks base method
func (m *MockPublisher) PublishResizeJob(ctx context.Context, jobId uint64) error {
	ret := m.ctrl.Call(m, "PublishResizeJob", ctx, jobId)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishResizeJob indicates an expected call of PublishResizeJob
func (mr *MockPublisherMockRecorder) PublishResizeJob(ctx, jobId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishResizeJob", reflect.TypeOf((*MockPublisher)(nil).PublishResizeJob), ctx, jobId)
}

// MockConsumer is a mock of Consumer interface
type MockConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerMockRecorder
}

// MockConsumerMockRecorder is the mock recorder for MockConsumer
type MockConsumerMockRecorder struct {
	mock *MockConsumer
}

// NewMockConsumer creates a new mock instance
func NewMockConsumer(ctrl *gomock.Controller) *MockConsumer {
	mock := &MockConsumer{ctrl: ctrl}
	mock.recorder = &MockConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConsumer) EXPECT() *MockConsumerMockRecorder {
	return m.recorder
}

// ConsumeResizeJobs mocks base method
func (m *MockConsumer) ConsumeResizeJobs(ctx context.Context, ch chan<- uint64) error {
	ret := m.ctrl.Call(m, "ConsumeResizeJobs", ctx, ch)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConsumeResizeJobs indicates an expected call of ConsumeResizeJobs
func (mr *MockConsumerMockRecorder) ConsumeResizeJobs(ctx, ch interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumeResizeJobs", reflect.TypeOf((*MockConsumer)(nil).ConsumeResizeJobs), ctx, ch)
}

// MockPublisherConsumer is a mock of PublisherConsumer interface
type MockPublisherConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockPublisherConsumerMockRecorder
}

// MockPublisherConsumerMockRecorder is the mock recorder for MockPublisherConsumer
type MockPublisherConsumerMockRecorder struct {
	mock *MockPublisherConsumer
}

// NewMockPublisherConsumer creates a new mock instance
func NewMockPublisherConsumer(ctrl *gomock.Controller) *MockPublisherConsumer {
	mock := &MockPublisherConsumer{ctrl: ctrl}
	mock.recorder = &MockPublisherConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPublisherConsumer) EXPECT() *MockPublisherConsumerMockRecorder {
	return m.recorder
}

// PublishResizeJob mocks base method
func (m *MockPublisherConsumer) PublishResizeJob(ctx context.Context, jobId uint64) error {
	ret := m.ctrl.Call(m, "PublishResizeJob", ctx, jobId)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishResizeJob indicates an expected call of PublishResizeJob
func (mr *MockPublisherConsumerMockRecorder) PublishResizeJob(ctx, jobId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishResizeJob", reflect.TypeOf((*MockPublisherConsumer)(nil).PublishResizeJob), ctx, jobId)
}

// ConsumeResizeJobs mocks base method
func (m *MockPublisherConsumer) ConsumeResizeJobs(ctx context.Context, ch chan<- uint64) error {
	ret := m.ctrl.Call(m, "ConsumeResizeJobs", ctx, ch)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConsumeResizeJobs indicates an expected call of ConsumeResizeJobs
func (mr *MockPublisherConsumerMockRecorder) ConsumeResizeJobs(ctx, ch interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumeResizeJobs", reflect.TypeOf((*MockPublisherConsumer)(nil).ConsumeResizeJobs), ctx, ch)
}