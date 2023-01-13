// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ../amqpwrap/amqpwrap.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	amqpwrap "github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/amqpwrap"
	amqp "github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/go-amqp"
	gomock "github.com/golang/mock/gomock"
)

// MockAMQPReceiver is a mock of AMQPReceiver interface.
type MockAMQPReceiver struct {
	ctrl     *gomock.Controller
	recorder *MockAMQPReceiverMockRecorder
}

// MockAMQPReceiverMockRecorder is the mock recorder for MockAMQPReceiver.
type MockAMQPReceiverMockRecorder struct {
	mock *MockAMQPReceiver
}

// NewMockAMQPReceiver creates a new mock instance.
func NewMockAMQPReceiver(ctrl *gomock.Controller) *MockAMQPReceiver {
	mock := &MockAMQPReceiver{ctrl: ctrl}
	mock.recorder = &MockAMQPReceiverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAMQPReceiver) EXPECT() *MockAMQPReceiverMockRecorder {
	return m.recorder
}

// AcceptMessage mocks base method.
func (m *MockAMQPReceiver) AcceptMessage(ctx context.Context, msg *amqp.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptMessage", ctx, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// AcceptMessage indicates an expected call of AcceptMessage.
func (mr *MockAMQPReceiverMockRecorder) AcceptMessage(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptMessage", reflect.TypeOf((*MockAMQPReceiver)(nil).AcceptMessage), ctx, msg)
}

// Credits mocks base method.
func (m *MockAMQPReceiver) Credits() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Credits")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// Credits indicates an expected call of Credits.
func (mr *MockAMQPReceiverMockRecorder) Credits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Credits", reflect.TypeOf((*MockAMQPReceiver)(nil).Credits))
}

// IssueCredit mocks base method.
func (m *MockAMQPReceiver) IssueCredit(credit uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssueCredit", credit)
	ret0, _ := ret[0].(error)
	return ret0
}

// IssueCredit indicates an expected call of IssueCredit.
func (mr *MockAMQPReceiverMockRecorder) IssueCredit(credit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssueCredit", reflect.TypeOf((*MockAMQPReceiver)(nil).IssueCredit), credit)
}

// LinkName mocks base method.
func (m *MockAMQPReceiver) LinkName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkName")
	ret0, _ := ret[0].(string)
	return ret0
}

// LinkName indicates an expected call of LinkName.
func (mr *MockAMQPReceiverMockRecorder) LinkName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkName", reflect.TypeOf((*MockAMQPReceiver)(nil).LinkName))
}

// LinkSourceFilterValue mocks base method.
func (m *MockAMQPReceiver) LinkSourceFilterValue(name string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkSourceFilterValue", name)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// LinkSourceFilterValue indicates an expected call of LinkSourceFilterValue.
func (mr *MockAMQPReceiverMockRecorder) LinkSourceFilterValue(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSourceFilterValue", reflect.TypeOf((*MockAMQPReceiver)(nil).LinkSourceFilterValue), name)
}

// ModifyMessage mocks base method.
func (m *MockAMQPReceiver) ModifyMessage(ctx context.Context, msg *amqp.Message, options *amqp.ModifyMessageOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyMessage", ctx, msg, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyMessage indicates an expected call of ModifyMessage.
func (mr *MockAMQPReceiverMockRecorder) ModifyMessage(ctx, msg, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyMessage", reflect.TypeOf((*MockAMQPReceiver)(nil).ModifyMessage), ctx, msg, options)
}

// Prefetched mocks base method.
func (m *MockAMQPReceiver) Prefetched() *amqp.Message {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prefetched")
	ret0, _ := ret[0].(*amqp.Message)
	return ret0
}

// Prefetched indicates an expected call of Prefetched.
func (mr *MockAMQPReceiverMockRecorder) Prefetched() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prefetched", reflect.TypeOf((*MockAMQPReceiver)(nil).Prefetched))
}

// Receive mocks base method.
func (m *MockAMQPReceiver) Receive(ctx context.Context) (*amqp.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive", ctx)
	ret0, _ := ret[0].(*amqp.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Receive indicates an expected call of Receive.
func (mr *MockAMQPReceiverMockRecorder) Receive(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockAMQPReceiver)(nil).Receive), ctx)
}

// RejectMessage mocks base method.
func (m *MockAMQPReceiver) RejectMessage(ctx context.Context, msg *amqp.Message, e *amqp.Error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RejectMessage", ctx, msg, e)
	ret0, _ := ret[0].(error)
	return ret0
}

// RejectMessage indicates an expected call of RejectMessage.
func (mr *MockAMQPReceiverMockRecorder) RejectMessage(ctx, msg, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RejectMessage", reflect.TypeOf((*MockAMQPReceiver)(nil).RejectMessage), ctx, msg, e)
}

// ReleaseMessage mocks base method.
func (m *MockAMQPReceiver) ReleaseMessage(ctx context.Context, msg *amqp.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseMessage", ctx, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReleaseMessage indicates an expected call of ReleaseMessage.
func (mr *MockAMQPReceiverMockRecorder) ReleaseMessage(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseMessage", reflect.TypeOf((*MockAMQPReceiver)(nil).ReleaseMessage), ctx, msg)
}

// MockAMQPReceiverCloser is a mock of AMQPReceiverCloser interface.
type MockAMQPReceiverCloser struct {
	ctrl     *gomock.Controller
	recorder *MockAMQPReceiverCloserMockRecorder
}

// MockAMQPReceiverCloserMockRecorder is the mock recorder for MockAMQPReceiverCloser.
type MockAMQPReceiverCloserMockRecorder struct {
	mock *MockAMQPReceiverCloser
}

// NewMockAMQPReceiverCloser creates a new mock instance.
func NewMockAMQPReceiverCloser(ctrl *gomock.Controller) *MockAMQPReceiverCloser {
	mock := &MockAMQPReceiverCloser{ctrl: ctrl}
	mock.recorder = &MockAMQPReceiverCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAMQPReceiverCloser) EXPECT() *MockAMQPReceiverCloserMockRecorder {
	return m.recorder
}

// AcceptMessage mocks base method.
func (m *MockAMQPReceiverCloser) AcceptMessage(ctx context.Context, msg *amqp.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptMessage", ctx, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// AcceptMessage indicates an expected call of AcceptMessage.
func (mr *MockAMQPReceiverCloserMockRecorder) AcceptMessage(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptMessage", reflect.TypeOf((*MockAMQPReceiverCloser)(nil).AcceptMessage), ctx, msg)
}

// Close mocks base method.
func (m *MockAMQPReceiverCloser) Close(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAMQPReceiverCloserMockRecorder) Close(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAMQPReceiverCloser)(nil).Close), ctx)
}

// Credits mocks base method.
func (m *MockAMQPReceiverCloser) Credits() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Credits")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// Credits indicates an expected call of Credits.
func (mr *MockAMQPReceiverCloserMockRecorder) Credits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Credits", reflect.TypeOf((*MockAMQPReceiverCloser)(nil).Credits))
}

// IssueCredit mocks base method.
func (m *MockAMQPReceiverCloser) IssueCredit(credit uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssueCredit", credit)
	ret0, _ := ret[0].(error)
	return ret0
}

// IssueCredit indicates an expected call of IssueCredit.
func (mr *MockAMQPReceiverCloserMockRecorder) IssueCredit(credit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssueCredit", reflect.TypeOf((*MockAMQPReceiverCloser)(nil).IssueCredit), credit)
}

// LinkName mocks base method.
func (m *MockAMQPReceiverCloser) LinkName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkName")
	ret0, _ := ret[0].(string)
	return ret0
}

// LinkName indicates an expected call of LinkName.
func (mr *MockAMQPReceiverCloserMockRecorder) LinkName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkName", reflect.TypeOf((*MockAMQPReceiverCloser)(nil).LinkName))
}

// LinkSourceFilterValue mocks base method.
func (m *MockAMQPReceiverCloser) LinkSourceFilterValue(name string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkSourceFilterValue", name)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// LinkSourceFilterValue indicates an expected call of LinkSourceFilterValue.
func (mr *MockAMQPReceiverCloserMockRecorder) LinkSourceFilterValue(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSourceFilterValue", reflect.TypeOf((*MockAMQPReceiverCloser)(nil).LinkSourceFilterValue), name)
}

// ModifyMessage mocks base method.
func (m *MockAMQPReceiverCloser) ModifyMessage(ctx context.Context, msg *amqp.Message, options *amqp.ModifyMessageOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyMessage", ctx, msg, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyMessage indicates an expected call of ModifyMessage.
func (mr *MockAMQPReceiverCloserMockRecorder) ModifyMessage(ctx, msg, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyMessage", reflect.TypeOf((*MockAMQPReceiverCloser)(nil).ModifyMessage), ctx, msg, options)
}

// Prefetched mocks base method.
func (m *MockAMQPReceiverCloser) Prefetched() *amqp.Message {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prefetched")
	ret0, _ := ret[0].(*amqp.Message)
	return ret0
}

// Prefetched indicates an expected call of Prefetched.
func (mr *MockAMQPReceiverCloserMockRecorder) Prefetched() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prefetched", reflect.TypeOf((*MockAMQPReceiverCloser)(nil).Prefetched))
}

// Receive mocks base method.
func (m *MockAMQPReceiverCloser) Receive(ctx context.Context) (*amqp.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive", ctx)
	ret0, _ := ret[0].(*amqp.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Receive indicates an expected call of Receive.
func (mr *MockAMQPReceiverCloserMockRecorder) Receive(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockAMQPReceiverCloser)(nil).Receive), ctx)
}

// RejectMessage mocks base method.
func (m *MockAMQPReceiverCloser) RejectMessage(ctx context.Context, msg *amqp.Message, e *amqp.Error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RejectMessage", ctx, msg, e)
	ret0, _ := ret[0].(error)
	return ret0
}

// RejectMessage indicates an expected call of RejectMessage.
func (mr *MockAMQPReceiverCloserMockRecorder) RejectMessage(ctx, msg, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RejectMessage", reflect.TypeOf((*MockAMQPReceiverCloser)(nil).RejectMessage), ctx, msg, e)
}

// ReleaseMessage mocks base method.
func (m *MockAMQPReceiverCloser) ReleaseMessage(ctx context.Context, msg *amqp.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseMessage", ctx, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReleaseMessage indicates an expected call of ReleaseMessage.
func (mr *MockAMQPReceiverCloserMockRecorder) ReleaseMessage(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseMessage", reflect.TypeOf((*MockAMQPReceiverCloser)(nil).ReleaseMessage), ctx, msg)
}

// MockAMQPSender is a mock of AMQPSender interface.
type MockAMQPSender struct {
	ctrl     *gomock.Controller
	recorder *MockAMQPSenderMockRecorder
}

// MockAMQPSenderMockRecorder is the mock recorder for MockAMQPSender.
type MockAMQPSenderMockRecorder struct {
	mock *MockAMQPSender
}

// NewMockAMQPSender creates a new mock instance.
func NewMockAMQPSender(ctrl *gomock.Controller) *MockAMQPSender {
	mock := &MockAMQPSender{ctrl: ctrl}
	mock.recorder = &MockAMQPSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAMQPSender) EXPECT() *MockAMQPSenderMockRecorder {
	return m.recorder
}

// LinkName mocks base method.
func (m *MockAMQPSender) LinkName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkName")
	ret0, _ := ret[0].(string)
	return ret0
}

// LinkName indicates an expected call of LinkName.
func (mr *MockAMQPSenderMockRecorder) LinkName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkName", reflect.TypeOf((*MockAMQPSender)(nil).LinkName))
}

// MaxMessageSize mocks base method.
func (m *MockAMQPSender) MaxMessageSize() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxMessageSize")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// MaxMessageSize indicates an expected call of MaxMessageSize.
func (mr *MockAMQPSenderMockRecorder) MaxMessageSize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxMessageSize", reflect.TypeOf((*MockAMQPSender)(nil).MaxMessageSize))
}

// Send mocks base method.
func (m *MockAMQPSender) Send(ctx context.Context, msg *amqp.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", ctx, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockAMQPSenderMockRecorder) Send(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockAMQPSender)(nil).Send), ctx, msg)
}

// MockAMQPSenderCloser is a mock of AMQPSenderCloser interface.
type MockAMQPSenderCloser struct {
	ctrl     *gomock.Controller
	recorder *MockAMQPSenderCloserMockRecorder
}

// MockAMQPSenderCloserMockRecorder is the mock recorder for MockAMQPSenderCloser.
type MockAMQPSenderCloserMockRecorder struct {
	mock *MockAMQPSenderCloser
}

// NewMockAMQPSenderCloser creates a new mock instance.
func NewMockAMQPSenderCloser(ctrl *gomock.Controller) *MockAMQPSenderCloser {
	mock := &MockAMQPSenderCloser{ctrl: ctrl}
	mock.recorder = &MockAMQPSenderCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAMQPSenderCloser) EXPECT() *MockAMQPSenderCloserMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockAMQPSenderCloser) Close(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAMQPSenderCloserMockRecorder) Close(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAMQPSenderCloser)(nil).Close), ctx)
}

// LinkName mocks base method.
func (m *MockAMQPSenderCloser) LinkName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkName")
	ret0, _ := ret[0].(string)
	return ret0
}

// LinkName indicates an expected call of LinkName.
func (mr *MockAMQPSenderCloserMockRecorder) LinkName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkName", reflect.TypeOf((*MockAMQPSenderCloser)(nil).LinkName))
}

// MaxMessageSize mocks base method.
func (m *MockAMQPSenderCloser) MaxMessageSize() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxMessageSize")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// MaxMessageSize indicates an expected call of MaxMessageSize.
func (mr *MockAMQPSenderCloserMockRecorder) MaxMessageSize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxMessageSize", reflect.TypeOf((*MockAMQPSenderCloser)(nil).MaxMessageSize))
}

// Send mocks base method.
func (m *MockAMQPSenderCloser) Send(ctx context.Context, msg *amqp.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", ctx, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockAMQPSenderCloserMockRecorder) Send(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockAMQPSenderCloser)(nil).Send), ctx, msg)
}

// MockAMQPSession is a mock of AMQPSession interface.
type MockAMQPSession struct {
	ctrl     *gomock.Controller
	recorder *MockAMQPSessionMockRecorder
}

// MockAMQPSessionMockRecorder is the mock recorder for MockAMQPSession.
type MockAMQPSessionMockRecorder struct {
	mock *MockAMQPSession
}

// NewMockAMQPSession creates a new mock instance.
func NewMockAMQPSession(ctrl *gomock.Controller) *MockAMQPSession {
	mock := &MockAMQPSession{ctrl: ctrl}
	mock.recorder = &MockAMQPSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAMQPSession) EXPECT() *MockAMQPSessionMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockAMQPSession) Close(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAMQPSessionMockRecorder) Close(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAMQPSession)(nil).Close), ctx)
}

// NewReceiver mocks base method.
func (m *MockAMQPSession) NewReceiver(ctx context.Context, source string, opts *amqp.ReceiverOptions) (amqpwrap.AMQPReceiverCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewReceiver", ctx, source, opts)
	ret0, _ := ret[0].(amqpwrap.AMQPReceiverCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewReceiver indicates an expected call of NewReceiver.
func (mr *MockAMQPSessionMockRecorder) NewReceiver(ctx, source, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewReceiver", reflect.TypeOf((*MockAMQPSession)(nil).NewReceiver), ctx, source, opts)
}

// NewSender mocks base method.
func (m *MockAMQPSession) NewSender(ctx context.Context, target string, opts *amqp.SenderOptions) (amqpwrap.AMQPSenderCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewSender", ctx, target, opts)
	ret0, _ := ret[0].(amqpwrap.AMQPSenderCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewSender indicates an expected call of NewSender.
func (mr *MockAMQPSessionMockRecorder) NewSender(ctx, target, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSender", reflect.TypeOf((*MockAMQPSession)(nil).NewSender), ctx, target, opts)
}

// MockAMQPClient is a mock of AMQPClient interface.
type MockAMQPClient struct {
	ctrl     *gomock.Controller
	recorder *MockAMQPClientMockRecorder
}

// MockAMQPClientMockRecorder is the mock recorder for MockAMQPClient.
type MockAMQPClientMockRecorder struct {
	mock *MockAMQPClient
}

// NewMockAMQPClient creates a new mock instance.
func NewMockAMQPClient(ctrl *gomock.Controller) *MockAMQPClient {
	mock := &MockAMQPClient{ctrl: ctrl}
	mock.recorder = &MockAMQPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAMQPClient) EXPECT() *MockAMQPClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockAMQPClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAMQPClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAMQPClient)(nil).Close))
}

// NewSession mocks base method.
func (m *MockAMQPClient) NewSession(ctx context.Context, opts *amqp.SessionOptions) (amqpwrap.AMQPSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewSession", ctx, opts)
	ret0, _ := ret[0].(amqpwrap.AMQPSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewSession indicates an expected call of NewSession.
func (mr *MockAMQPClientMockRecorder) NewSession(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSession", reflect.TypeOf((*MockAMQPClient)(nil).NewSession), ctx, opts)
}
