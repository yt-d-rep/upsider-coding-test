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
		PaymentAmount string
	}
)

func (f *invoiceFactory) Issue(input *IssueInput) (*Invoice, error) {
	id := NewInvoiceID()
	issuedAt := f.clock.Now()
	paymentAmount, err := NewAmount(input.PaymentAmount)
	if err != nil {
		return &Invoice{}, err
	}
	// メタデータと金額計算に必要な情報を設定してinvoice生成
	invoice := &Invoice{
		id:            id,
		companyID:     input.CompanyID,
		partnerID:     input.PartnerID,
		issuedAt:      issuedAt,
		paymentAmount: paymentAmount,
		status:        Unpaid,
	}
	err = invoice.calculateFee(CurrentFeeRate)
	if err != nil {
		return &Invoice{}, err
	}
	err = invoice.calculateConsumptionTax(CurrentConsumptionTaxRate)
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
