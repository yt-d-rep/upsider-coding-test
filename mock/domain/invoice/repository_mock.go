// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/invoice/repository.go
//
// Generated by this command:
//
//	mockgen -source ./domain/invoice/repository.go -destination ./mock/domain/invoice/repository_mock.go -package invoice_mock
//

// Package invoice_mock is a generated GoMock package.
package invoice_mock

import (
	reflect "reflect"
	company "upsider-coding-test/domain/company"
	invoice "upsider-coding-test/domain/invoice"
	shared "upsider-coding-test/shared"

	gomock "go.uber.org/mock/gomock"
)

// MockInvoiceRepository is a mock of InvoiceRepository interface.
type MockInvoiceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockInvoiceRepositoryMockRecorder
}

// MockInvoiceRepositoryMockRecorder is the mock recorder for MockInvoiceRepository.
type MockInvoiceRepositoryMockRecorder struct {
	mock *MockInvoiceRepository
}

// NewMockInvoiceRepository creates a new mock instance.
func NewMockInvoiceRepository(ctrl *gomock.Controller) *MockInvoiceRepository {
	mock := &MockInvoiceRepository{ctrl: ctrl}
	mock.recorder = &MockInvoiceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInvoiceRepository) EXPECT() *MockInvoiceRepositoryMockRecorder {
	return m.recorder
}

// ListBetween mocks base method.
func (m *MockInvoiceRepository) ListBetween(timeRange *shared.TimeRange, companyID company.CompanyID) ([]*invoice.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBetween", timeRange, companyID)
	ret0, _ := ret[0].([]*invoice.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBetween indicates an expected call of ListBetween.
func (mr *MockInvoiceRepositoryMockRecorder) ListBetween(timeRange, companyID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBetween", reflect.TypeOf((*MockInvoiceRepository)(nil).ListBetween), timeRange, companyID)
}

// Save mocks base method.
func (m *MockInvoiceRepository) Save(invoice *invoice.Invoice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", invoice)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockInvoiceRepositoryMockRecorder) Save(invoice any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockInvoiceRepository)(nil).Save), invoice)
}
