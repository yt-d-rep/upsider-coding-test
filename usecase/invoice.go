package usecase

import (
	"time"
	"upsider-coding-test/domain/company"
	"upsider-coding-test/domain/invoice"
	"upsider-coding-test/shared"
)

type (
	InvoiceUsecase interface {
		Issue(input *IssueInput) (*invoice.Invoice, error)
		ListBetween(input *ListBetweenInput) ([]*invoice.Invoice, error)
	}
	invoiceUsecase struct {
		invoiceFactory    invoice.InvoiceFactory
		invoiceRepository invoice.InvoiceRepository
	}
	IssueInput struct {
		CompanyID     string
		PartnerID     string
		PaymentAmount string
	}
	ListBetweenInput struct {
		From      time.Time
		To        time.Time
		CompanyID string
	}
)

func (u *invoiceUsecase) Issue(input *IssueInput) (*invoice.Invoice, error) {
	companyID, err := company.ParseCompanyID(input.CompanyID)
	if err != nil {
		return nil, err
	}
	partnerID, err := company.ParsePartnerID(input.PartnerID)
	if err != nil {
		return nil, err
	}
	invoice, err := u.invoiceFactory.Issue(&invoice.IssueInput{
		CompanyID:     companyID,
		PartnerID:     partnerID,
		PaymentAmount: input.PaymentAmount,
	})
	if err != nil {
		return nil, err
	}
	if err := u.invoiceRepository.Save(invoice); err != nil {
		return nil, err
	}
	return invoice, nil
}

func (u *invoiceUsecase) ListBetween(input *ListBetweenInput) ([]*invoice.Invoice, error) {
	companyID, err := company.ParseCompanyID(input.CompanyID)
	if err != nil {
		return nil, err
	}
	timeRange, err := shared.NewTimeRange(input.From, input.To)
	if err != nil {
		return nil, err
	}
	return u.invoiceRepository.ListBetween(timeRange, companyID)
}
