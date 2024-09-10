package invoice

import (
	"upsider-coding-test/domain/company"
	"upsider-coding-test/shared"
)

type (
	InvoiceFactory interface {
		Issue(input *IssueInput) (*Invoice, error)
	}
	invoiceFactory struct {
		clock shared.Clock
	}
	IssueInput struct {
		CompanyID     company.CompanyID
		PartnerID     company.PartnerID
		PaymentAmount int64
	}
)

func (f *invoiceFactory) Issue(input *IssueInput) (*Invoice, error) {
	id := NewInvoiceID()
	issuedAt := f.clock.Now()
	paymentAmount, err := NewAmount(input.PaymentAmount)
	if err != nil {
		return &Invoice{}, err
	}
	if paymentAmount.IsZero() {
		return &Invoice{}, &shared.ValidationError{Field: "payment_amount", Err: "payment amount must be greater than 0"}
	}
	feeRate, err := NewFeeRate()
	if err != nil {
		return &Invoice{}, err
	}
	consumptionTaxRate, err := NewConsumptionTaxRate()
	if err != nil {
		return &Invoice{}, err
	}
	invoice := &Invoice{
		id:                 id,
		companyID:          input.CompanyID,
		partnerID:          input.PartnerID,
		feeRate:            feeRate,
		consumptionTaxRate: consumptionTaxRate,
		issuedAt:           issuedAt,
		paymentAmount:      paymentAmount,
		status:             Unpaid,
	}
	err = invoice.calculateFee()
	if err != nil {
		return &Invoice{}, err
	}
	err = invoice.calculateConsumptionTax()
	if err != nil {
		return &Invoice{}, err
	}
	err = invoice.calculateInvoiceAmount()
	if err != nil {
		return &Invoice{}, err
	}
	invoice.setDue()
	return invoice, nil
}
