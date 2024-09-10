package invoice

import (
	"upsider-coding-test/domain/company"
	"upsider-coding-test/shared"
)

type (
	InvoiceRepository interface {
		Save(invoice *Invoice) error
		ListBetween(timeRange *shared.TimeRange, companyID company.CompanyID) ([]*Invoice, error)
	}
)
