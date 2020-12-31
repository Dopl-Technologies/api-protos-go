// Code generated by MockGen. DO NOT EDIT.
// Source: device_service.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	platform "github.com/pyrus/platform-public/pkg/protos/platform"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockDeviceServiceClient is a mock of DeviceServiceClient interface
type MockDeviceServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceServiceClientMockRecorder
}

// MockDeviceServiceClientMockRecorder is the mock recorder for MockDeviceServiceClient
type MockDeviceServiceClientMockRecorder struct {
	mock *MockDeviceServiceClient
}

// NewMockDeviceServiceClient creates a new mock instance
func NewMockDeviceServiceClient(ctrl *gomock.Controller) *MockDeviceServiceClient {
	mock := &MockDeviceServiceClient{ctrl: ctrl}
	mock.recorder = &MockDeviceServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeviceServiceClient) EXPECT() *MockDeviceServiceClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockDeviceServiceClient) Create(ctx context.Context, in *platform.CreateDeviceRequest, opts ...grpc.CallOption) (*platform.CreateDeviceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*platform.CreateDeviceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockDeviceServiceClientMockRecorder) Create(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDeviceServiceClient)(nil).Create), varargs...)
}

// Get mocks base method
func (m *MockDeviceServiceClient) Get(ctx context.Context, in *platform.GetDeviceRequest, opts ...grpc.CallOption) (*platform.GetDeviceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*platform.GetDeviceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockDeviceServiceClientMockRecorder) Get(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDeviceServiceClient)(nil).Get), varargs...)
}

// Update mocks base method
func (m *MockDeviceServiceClient) Update(ctx context.Context, in *platform.UpdateDeviceRequest, opts ...grpc.CallOption) (*platform.UpdateDeviceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*platform.UpdateDeviceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockDeviceServiceClientMockRecorder) Update(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDeviceServiceClient)(nil).Update), varargs...)
}

// List mocks base method
func (m *MockDeviceServiceClient) List(ctx context.Context, in *platform.ListDevicesRequest, opts ...grpc.CallOption) (platform.DeviceService_ListClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(platform.DeviceService_ListClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockDeviceServiceClientMockRecorder) List(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDeviceServiceClient)(nil).List), varargs...)
}

// Delete mocks base method
func (m *MockDeviceServiceClient) Delete(ctx context.Context, in *platform.DeleteDeviceRequest, opts ...grpc.CallOption) (*platform.DeleteDeviceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*platform.DeleteDeviceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockDeviceServiceClientMockRecorder) Delete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDeviceServiceClient)(nil).Delete), varargs...)
}

// MockDeviceService_ListClient is a mock of DeviceService_ListClient interface
type MockDeviceService_ListClient struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceService_ListClientMockRecorder
}

// MockDeviceService_ListClientMockRecorder is the mock recorder for MockDeviceService_ListClient
type MockDeviceService_ListClientMockRecorder struct {
	mock *MockDeviceService_ListClient
}

// NewMockDeviceService_ListClient creates a new mock instance
func NewMockDeviceService_ListClient(ctrl *gomock.Controller) *MockDeviceService_ListClient {
	mock := &MockDeviceService_ListClient{ctrl: ctrl}
	mock.recorder = &MockDeviceService_ListClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeviceService_ListClient) EXPECT() *MockDeviceService_ListClientMockRecorder {
	return m.recorder
}

// Recv mocks base method
func (m *MockDeviceService_ListClient) Recv() (*platform.ListDevicesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*platform.ListDevicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockDeviceService_ListClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockDeviceService_ListClient)(nil).Recv))
}

// Header mocks base method
func (m *MockDeviceService_ListClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockDeviceService_ListClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockDeviceService_ListClient)(nil).Header))
}

// Trailer mocks base method
func (m *MockDeviceService_ListClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockDeviceService_ListClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockDeviceService_ListClient)(nil).Trailer))
}

// CloseSend mocks base method
func (m *MockDeviceService_ListClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockDeviceService_ListClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockDeviceService_ListClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockDeviceService_ListClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockDeviceService_ListClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockDeviceService_ListClient)(nil).Context))
}

// SendMsg mocks base method
func (m_2 *MockDeviceService_ListClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockDeviceService_ListClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockDeviceService_ListClient)(nil).SendMsg), m)
}

// RecvMsg mocks base method
func (m_2 *MockDeviceService_ListClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockDeviceService_ListClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockDeviceService_ListClient)(nil).RecvMsg), m)
}

// MockDeviceServiceServer is a mock of DeviceServiceServer interface
type MockDeviceServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceServiceServerMockRecorder
}

// MockDeviceServiceServerMockRecorder is the mock recorder for MockDeviceServiceServer
type MockDeviceServiceServerMockRecorder struct {
	mock *MockDeviceServiceServer
}

// NewMockDeviceServiceServer creates a new mock instance
func NewMockDeviceServiceServer(ctrl *gomock.Controller) *MockDeviceServiceServer {
	mock := &MockDeviceServiceServer{ctrl: ctrl}
	mock.recorder = &MockDeviceServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeviceServiceServer) EXPECT() *MockDeviceServiceServerMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockDeviceServiceServer) Create(arg0 context.Context, arg1 *platform.CreateDeviceRequest) (*platform.CreateDeviceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*platform.CreateDeviceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockDeviceServiceServerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDeviceServiceServer)(nil).Create), arg0, arg1)
}

// Get mocks base method
func (m *MockDeviceServiceServer) Get(arg0 context.Context, arg1 *platform.GetDeviceRequest) (*platform.GetDeviceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*platform.GetDeviceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockDeviceServiceServerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDeviceServiceServer)(nil).Get), arg0, arg1)
}

// Update mocks base method
func (m *MockDeviceServiceServer) Update(arg0 context.Context, arg1 *platform.UpdateDeviceRequest) (*platform.UpdateDeviceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*platform.UpdateDeviceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockDeviceServiceServerMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDeviceServiceServer)(nil).Update), arg0, arg1)
}

// List mocks base method
func (m *MockDeviceServiceServer) List(arg0 *platform.ListDevicesRequest, arg1 platform.DeviceService_ListServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// List indicates an expected call of List
func (mr *MockDeviceServiceServerMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDeviceServiceServer)(nil).List), arg0, arg1)
}

// Delete mocks base method
func (m *MockDeviceServiceServer) Delete(arg0 context.Context, arg1 *platform.DeleteDeviceRequest) (*platform.DeleteDeviceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*platform.DeleteDeviceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockDeviceServiceServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDeviceServiceServer)(nil).Delete), arg0, arg1)
}

// MockDeviceService_ListServer is a mock of DeviceService_ListServer interface
type MockDeviceService_ListServer struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceService_ListServerMockRecorder
}

// MockDeviceService_ListServerMockRecorder is the mock recorder for MockDeviceService_ListServer
type MockDeviceService_ListServerMockRecorder struct {
	mock *MockDeviceService_ListServer
}

// NewMockDeviceService_ListServer creates a new mock instance
func NewMockDeviceService_ListServer(ctrl *gomock.Controller) *MockDeviceService_ListServer {
	mock := &MockDeviceService_ListServer{ctrl: ctrl}
	mock.recorder = &MockDeviceService_ListServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeviceService_ListServer) EXPECT() *MockDeviceService_ListServerMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockDeviceService_ListServer) Send(arg0 *platform.ListDevicesResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockDeviceService_ListServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockDeviceService_ListServer)(nil).Send), arg0)
}

// SetHeader mocks base method
func (m *MockDeviceService_ListServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockDeviceService_ListServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockDeviceService_ListServer)(nil).SetHeader), arg0)
}

// SendHeader mocks base method
func (m *MockDeviceService_ListServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockDeviceService_ListServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockDeviceService_ListServer)(nil).SendHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockDeviceService_ListServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockDeviceService_ListServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockDeviceService_ListServer)(nil).SetTrailer), arg0)
}

// Context mocks base method
func (m *MockDeviceService_ListServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockDeviceService_ListServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockDeviceService_ListServer)(nil).Context))
}

// SendMsg mocks base method
func (m_2 *MockDeviceService_ListServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockDeviceService_ListServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockDeviceService_ListServer)(nil).SendMsg), m)
}

// RecvMsg mocks base method
func (m_2 *MockDeviceService_ListServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockDeviceService_ListServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockDeviceService_ListServer)(nil).RecvMsg), m)
}
