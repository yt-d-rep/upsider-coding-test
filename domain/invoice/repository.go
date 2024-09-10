package invoice

import (
	"upsider-base/domain/company"
	"upsider-base/shared"
)

type (
	InvoiceRepository interface {
		Save(invoice *Invoice) error
		ListBetween(timeRange *shared.TimeRange, companyID company.CompanyID) ([]*Invoice, error)
	}
)
