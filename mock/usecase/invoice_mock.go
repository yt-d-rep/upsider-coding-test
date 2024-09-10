// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase/invoice.go
//
// Generated by this command:
//
//	mockgen -source ./usecase/invoice.go -destination ./mock/usecase/invoice_mock.go -package usecase_mock
//

// Package usecase_mock is a generated GoMock package.
package usecase_mock

import (
	reflect "reflect"
	invoice "upsider-base/domain/invoice"
	usecase "upsider-base/usecase"

	gomock "go.uber.org/mock/gomock"
)

// MockInvoiceUsecase is a mock of InvoiceUsecase interface.
type MockInvoiceUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockInvoiceUsecaseMockRecorder
}

// MockInvoiceUsecaseMockRecorder is the mock recorder for MockInvoiceUsecase.
type MockInvoiceUsecaseMockRecorder struct {
	mock *MockInvoiceUsecase
}

// NewMockInvoiceUsecase creates a new mock instance.
func NewMockInvoiceUsecase(ctrl *gomock.Controller) *MockInvoiceUsecase {
	mock := &MockInvoiceUsecase{ctrl: ctrl}
	mock.recorder = &MockInvoiceUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInvoiceUsecase) EXPECT() *MockInvoiceUsecaseMockRecorder {
	return m.recorder
}

// Issue mocks base method.
func (m *MockInvoiceUsecase) Issue(input *usecase.IssueInput) (*invoice.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Issue", input)
	ret0, _ := ret[0].(*invoice.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Issue indicates an expected call of Issue.
func (mr *MockInvoiceUsecaseMockRecorder) Issue(input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Issue", reflect.TypeOf((*MockInvoiceUsecase)(nil).Issue), input)
}

// ListBetween mocks base method.
func (m *MockInvoiceUsecase) ListBetween(input *usecase.ListBetweenInput) ([]*invoice.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBetween", input)
	ret0, _ := ret[0].([]*invoice.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBetween indicates an expected call of ListBetween.
func (mr *MockInvoiceUsecaseMockRecorder) ListBetween(input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBetween", reflect.TypeOf((*MockInvoiceUsecase)(nil).ListBetween), input)
}
