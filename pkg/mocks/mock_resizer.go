// Code generated by MockGen. DO NOT EDIT.
// Source: resizer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	resizer "github.com/schepelin/imageresizer/pkg/resizer"
	image "image"
	reflect "reflect"
	time "time"
)

// MockImageService is a mock of ImageService interface
type MockImageService struct {
	ctrl     *gomock.Controller
	recorder *MockImageServiceMockRecorder
}

// MockImageServiceMockRecorder is the mock recorder for MockImageService
type MockImageServiceMockRecorder struct {
	mock *MockImageService
}

// NewMockImageService creates a new mock instance
func NewMockImageService(ctrl *gomock.Controller) *MockImageService {
	mock := &MockImageService{ctrl: ctrl}
	mock.recorder = &MockImageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImageService) EXPECT() *MockImageServiceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockImageService) Create(ctx context.Context, raw *[]byte) (*resizer.Image, error) {
	ret := m.ctrl.Call(m, "Create", ctx, raw)
	ret0, _ := ret[0].(*resizer.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockImageServiceMockRecorder) Create(ctx, raw interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockImageService)(nil).Create), ctx, raw)
}

// Read mocks base method
func (m *MockImageService) Read(ctx context.Context, imgId string) (*resizer.Image, error) {
	ret := m.ctrl.Call(m, "Read", ctx, imgId)
	ret0, _ := ret[0].(*resizer.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockImageServiceMockRecorder) Read(ctx, imgId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockImageService)(nil).Read), ctx, imgId)
}

// Delete mocks base method
func (m *MockImageService) Delete(ctx context.Context, imgId string) error {
	ret := m.ctrl.Call(m, "Delete", ctx, imgId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockImageServiceMockRecorder) Delete(ctx, imgId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockImageService)(nil).Delete), ctx, imgId)
}

// ScheduleResizeJob mocks base method
func (m *MockImageService) ScheduleResizeJob(ctx context.Context, imgId string, width, height uint) (*resizer.ResizeJob, error) {
	ret := m.ctrl.Call(m, "ScheduleResizeJob", ctx, imgId, width, height)
	ret0, _ := ret[0].(*resizer.ResizeJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScheduleResizeJob indicates an expected call of ScheduleResizeJob
func (mr *MockImageServiceMockRecorder) ScheduleResizeJob(ctx, imgId, width, height interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleResizeJob", reflect.TypeOf((*MockImageService)(nil).ScheduleResizeJob), ctx, imgId, width, height)
}

// GetResizeJob mocks base method
func (m *MockImageService) GetResizeJob(ctx context.Context, jobId uint64) (*resizer.ResizeJob, error) {
	ret := m.ctrl.Call(m, "GetResizeJob", ctx, jobId)
	ret0, _ := ret[0].(*resizer.ResizeJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResizeJob indicates an expected call of GetResizeJob
func (mr *MockImageServiceMockRecorder) GetResizeJob(ctx, jobId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResizeJob", reflect.TypeOf((*MockImageService)(nil).GetResizeJob), ctx, jobId)
}

// MockClocker is a mock of Clocker interface
type MockClocker struct {
	ctrl     *gomock.Controller
	recorder *MockClockerMockRecorder
}

// MockClockerMockRecorder is the mock recorder for MockClocker
type MockClockerMockRecorder struct {
	mock *MockClocker
}

// NewMockClocker creates a new mock instance
func NewMockClocker(ctrl *gomock.Controller) *MockClocker {
	mock := &MockClocker{ctrl: ctrl}
	mock.recorder = &MockClockerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClocker) EXPECT() *MockClockerMockRecorder {
	return m.recorder
}

// Now mocks base method
func (m *MockClocker) Now() time.Time {
	ret := m.ctrl.Call(m, "Now")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Now indicates an expected call of Now
func (mr *MockClockerMockRecorder) Now() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Now", reflect.TypeOf((*MockClocker)(nil).Now))
}

// MockHasher is a mock of Hasher interface
type MockHasher struct {
	ctrl     *gomock.Controller
	recorder *MockHasherMockRecorder
}

// MockHasherMockRecorder is the mock recorder for MockHasher
type MockHasherMockRecorder struct {
	mock *MockHasher
}

// NewMockHasher creates a new mock instance
func NewMockHasher(ctrl *gomock.Controller) *MockHasher {
	mock := &MockHasher{ctrl: ctrl}
	mock.recorder = &MockHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHasher) EXPECT() *MockHasherMockRecorder {
	return m.recorder
}

// Gen mocks base method
func (m *MockHasher) Gen(raw *[]byte) string {
	ret := m.ctrl.Call(m, "Gen", raw)
	ret0, _ := ret[0].(string)
	return ret0
}

// Gen indicates an expected call of Gen
func (mr *MockHasherMockRecorder) Gen(raw interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gen", reflect.TypeOf((*MockHasher)(nil).Gen), raw)
}

// MockConverter is a mock of Converter interface
type MockConverter struct {
	ctrl     *gomock.Controller
	recorder *MockConverterMockRecorder
}

// MockConverterMockRecorder is the mock recorder for MockConverter
type MockConverterMockRecorder struct {
	mock *MockConverter
}

// NewMockConverter creates a new mock instance
func NewMockConverter(ctrl *gomock.Controller) *MockConverter {
	mock := &MockConverter{ctrl: ctrl}
	mock.recorder = &MockConverterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConverter) EXPECT() *MockConverterMockRecorder {
	return m.recorder
}

// Transform mocks base method
func (m *MockConverter) Transform(raw *[]byte) (image.Image, error) {
	ret := m.ctrl.Call(m, "Transform", raw)
	ret0, _ := ret[0].(image.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Transform indicates an expected call of Transform
func (mr *MockConverterMockRecorder) Transform(raw interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transform", reflect.TypeOf((*MockConverter)(nil).Transform), raw)
}

// Resize mocks base method
func (m *MockConverter) Resize(img *image.Image, width, height uint) ([]byte, error) {
	ret := m.ctrl.Call(m, "Resize", img, width, height)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resize indicates an expected call of Resize
func (mr *MockConverterMockRecorder) Resize(img, width, height interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resize", reflect.TypeOf((*MockConverter)(nil).Resize), img, width, height)
}

// MockResizeService is a mock of ResizeService interface
type MockResizeService struct {
	ctrl     *gomock.Controller
	recorder *MockResizeServiceMockRecorder
}

// MockResizeServiceMockRecorder is the mock recorder for MockResizeService
type MockResizeServiceMockRecorder struct {
	mock *MockResizeService
}

// NewMockResizeService creates a new mock instance
func NewMockResizeService(ctrl *gomock.Controller) *MockResizeService {
	mock := &MockResizeService{ctrl: ctrl}
	mock.recorder = &MockResizeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResizeService) EXPECT() *MockResizeServiceMockRecorder {
	return m.recorder
}

// SendResizeJob mocks base method
func (m *MockResizeService) SendResizeJob(ctx context.Context, req *resizer.ResizeServiceRequest) error {
	ret := m.ctrl.Call(m, "SendResizeJob", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendResizeJob indicates an expected call of SendResizeJob
func (mr *MockResizeServiceMockRecorder) SendResizeJob(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendResizeJob", reflect.TypeOf((*MockResizeService)(nil).SendResizeJob), ctx, req)
}

// RunResizeWorker mocks base method
func (m *MockResizeService) RunResizeWorker(ctx context.Context, ch chan uint64) error {
	ret := m.ctrl.Call(m, "RunResizeWorker", ctx, ch)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunResizeWorker indicates an expected call of RunResizeWorker
func (mr *MockResizeServiceMockRecorder) RunResizeWorker(ctx, ch interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunResizeWorker", reflect.TypeOf((*MockResizeService)(nil).RunResizeWorker), ctx, ch)
}
