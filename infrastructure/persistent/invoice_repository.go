package persistent

import (
	"database/sql"
	"fmt"
	"time"
	"upsider-base/domain/company"
	"upsider-base/domain/invoice"
	"upsider-base/shared"
)

type (
	invoiceRepository struct {
		db *sql.DB
	}
)

func (r *invoiceRepository) Save(invoice *invoice.Invoice) error {
	query := `
		INSERT INTO invoices (
			invoice_id,
			company_id,
			partner_id,
			issued_at,
			payment_amount,
			fee,
			fee_rate,
			consumption_tax,
			consumption_tax_rate,
			invoice_amount,
			payment_due_at,
			status
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	;`
	_, err := r.db.Exec(query,
		invoice.ID().String(),
		invoice.CompanyID().String(),
		invoice.PartnerID().String(),
		invoice.IssuedAt(),
		invoice.PaymentAmount(),
		invoice.Fee(),
		invoice.FeeRate().String(),
		invoice.ConsumptionTax(),
		invoice.ConsumptionTaxRate().String(),
		invoice.InvoiceAmount(),
		invoice.PaymentDueAt(),
		invoice.Status(),
	)
	if err != nil {
		return fmt.Errorf("failed to save invoice in Save: %w", err)
	}
	return nil
}

func (r *invoiceRepository) ListBetween(timeRange *shared.TimeRange, companyID company.CompanyID) ([]*invoice.Invoice, error) {
	query := `
		SELECT
			invoice_id,
			company_id,
			partner_id,
			issued_at,
			payment_amount,
			fee,
			fee_rate,
			consumption_tax,
			consumption_tax_rate,
			invoice_amount,
			payment_due_at,
			status
		FROM invoices
		WHERE company_id = $1
		AND issued_at BETWEEN $2 AND $3
		ORDER BY issued_at ASC
	;`
	rows, err := r.db.Query(query, companyID.String(), timeRange.From(), timeRange.To())
	if err != nil {
		return nil, fmt.Errorf("failed to list invoices in ListBetween: %w", err)
	}
	defer rows.Close()

	type invoiceRow struct {
		ID                 string
		CompanyID          string
		PartnerID          string
		IssuedAt           time.Time
		PaymentAmount      int64
		Fee                int64
		FeeRate            string
		ConsumptionTax     int64
		ConsumptionTaxRate string
		InvoiceAmount      int64
		PaymentDueAt       time.Time
		Status             int
	}
	var invoices []*invoice.Invoice
	for rows.Next() {
		var row invoiceRow
		err := rows.Scan(
			&row.ID,
			&row.CompanyID,
			&row.PartnerID,
			&row.IssuedAt,
			&row.PaymentAmount,
			&row.Fee,
			&row.FeeRate,
			&row.ConsumptionTax,
			&row.ConsumptionTaxRate,
			&row.InvoiceAmount,
			&row.PaymentDueAt,
			&row.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan invoice row in ListBetween: %w", err)
		}
		invoice, err := invoice.ParseInvoice(&invoice.ParseInvoiceInput{
			ID:                 row.ID,
			CompanyID:          row.CompanyID,
			PartnerID:          row.PartnerID,
			IssuedAt:           row.IssuedAt,
			PaymentAmount:      row.PaymentAmount,
			Fee:                row.Fee,
			FeeRate:            row.FeeRate,
			ConsumptionTax:     row.ConsumptionTax,
			ConsumptionTaxRate: row.ConsumptionTaxRate,
			InvoiceAmount:      row.InvoiceAmount,
			PaymentDueAt:       row.PaymentDueAt,
			Status:             row.Status,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to parse invoice in ListBetween: %w", err)
		}
		invoices = append(invoices, invoice)
	}
	return invoices, nil
}
