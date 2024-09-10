// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/invoice/factory.go
//
// Generated by this command:
//
//	mockgen -source ./domain/invoice/factory.go -destination ./mock/domain/invoice/factory_mock.go -package invoice_mock
//

// Package invoice_mock is a generated GoMock package.
package invoice_mock

import (
	reflect "reflect"
	invoice "upsider-coding-test/domain/invoice"

	gomock "go.uber.org/mock/gomock"
)

// MockInvoiceFactory is a mock of InvoiceFactory interface.
type MockInvoiceFactory struct {
	ctrl     *gomock.Controller
	recorder *MockInvoiceFactoryMockRecorder
}

// MockInvoiceFactoryMockRecorder is the mock recorder for MockInvoiceFactory.
type MockInvoiceFactoryMockRecorder struct {
	mock *MockInvoiceFactory
}

// NewMockInvoiceFactory creates a new mock instance.
func NewMockInvoiceFactory(ctrl *gomock.Controller) *MockInvoiceFactory {
	mock := &MockInvoiceFactory{ctrl: ctrl}
	mock.recorder = &MockInvoiceFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInvoiceFactory) EXPECT() *MockInvoiceFactoryMockRecorder {
	return m.recorder
}

// Issue mocks base method.
func (m *MockInvoiceFactory) Issue(input *invoice.IssueInput) (*invoice.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Issue", input)
	ret0, _ := ret[0].(*invoice.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Issue indicates an expected call of Issue.
func (mr *MockInvoiceFactoryMockRecorder) Issue(input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Issue", reflect.TypeOf((*MockInvoiceFactory)(nil).Issue), input)
}
