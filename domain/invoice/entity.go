package invoice

import (
	"time"
	"upsider-coding-test/domain/company"
)

type (
	Invoice struct {
		id             InvoiceID
		companyID      company.CompanyID
		partnerID      company.PartnerID
		issuedAt       time.Time
		paymentAmount  Amount
		fee            *Fee
		consumptionTax *ConsumptionTax
		invoiceAmount  Amount
		paymentDueAt   time.Time
		status         Status
	}
	ParseInvoiceInput struct {
		ID                 string
		CompanyID          string
		PartnerID          string
		IssuedAt           time.Time
		PaymentAmount      string
		Fee                string
		FeeRate            string
		ConsumptionTax     string
		ConsumptionTaxRate string
		InvoiceAmount      string
		PaymentDueAt       time.Time
		Status             int
	}
)

func ParseInvoice(i *ParseInvoiceInput) (*Invoice, error) {
	id, err := ParseInvoiceID(i.ID)
	if err != nil {
		return nil, err
	}
	companyID, err := company.ParseCompanyID(i.CompanyID)
	if err != nil {
		return nil, err
	}
	partnerID, err := company.ParsePartnerID(i.PartnerID)
	if err != nil {
		return nil, err
	}
	paymentAmount, err := NewAmount(i.PaymentAmount)
	if err != nil {
		return nil, err
	}
	fee, err := ParseFee(i.Fee, i.FeeRate)
	if err != nil {
		return nil, err
	}
	consumptionTax, err := ParseConsumptionTax(i.ConsumptionTax, i.ConsumptionTaxRate)
	if err != nil {
		return nil, err
	}
	invoiceAmount, err := NewAmount(i.InvoiceAmount)
	if err != nil {
		return nil, err
	}
	status := Status(i.Status)
	return &Invoice{
		id:             id,
		companyID:      companyID,
		partnerID:      partnerID,
		issuedAt:       i.IssuedAt,
		paymentAmount:  paymentAmount,
		fee:            fee,
		consumptionTax: consumptionTax,
		invoiceAmount:  invoiceAmount,
		paymentDueAt:   i.PaymentDueAt,
		status:         status,
	}, nil
}

func (i *Invoice) ID() InvoiceID {
	return i.id
}
func (i *Invoice) CompanyID() company.CompanyID {
	return i.companyID
}
func (i *Invoice) PartnerID() company.PartnerID {
	return i.partnerID
}
func (i *Invoice) IssuedAt() time.Time {
	return i.issuedAt
}
func (i *Invoice) PaymentAmount() Amount {
	return i.paymentAmount
}
func (i *Invoice) Fee() Fee {
	return *i.fee
}
func (i *Invoice) ConsumptionTax() ConsumptionTax {
	return *i.consumptionTax
}
func (i *Invoice) InvoiceAmount() Amount {
	return i.invoiceAmount
}
func (i *Invoice) PaymentDueAt() time.Time {
	return i.paymentDueAt
}
func (i *Invoice) Status() Status {
	return i.status
}

// calculateFee は手数料を計算します
// 手数料 = 支払い金額 * 手数料率
func (i *Invoice) calculateFee(rate string) error {
	feeRate, err := NewRate(rate)
	if err != nil {
		return err
	}
	feeValue := i.PaymentAmount().MulRate(feeRate)
	i.fee = NewFee(feeValue, feeRate)
	return nil
}

// calculateConsumptionTax は消費税を計算します
// 消費税 = 手数料 * 消費税率
func (i *Invoice) calculateConsumptionTax(rate string) error {
	consumptionTaxRate, err := NewRate(rate)
	if err != nil {
		return err
	}
	consumptionTaxValue := i.Fee().Value().MulRate(consumptionTaxRate)
	i.consumptionTax = NewConsumptionTax(consumptionTaxValue, consumptionTaxRate)
	return nil
}

// calculateInvoiceAmount は請求金額を計算します
// 請求金額 = 支払い金額 + 手数料 + 消費税
func (i *Invoice) calculateInvoiceAmount() error {
	i.invoiceAmount = i.PaymentAmount().Add(i.Fee().Value()).Add(i.ConsumptionTax().Value())
	return nil
}

// setDue は支払期限を設定します
// 請求書発行日の14日後の23:59:59までが支払い期限
func (i *Invoice) setDue() {
	// TODO: クライアントのTZはAsia/Tokyoに固定していますが、本来はユーザーのTZに合わせるべき
	issuedAtAsUserTZ := i.issuedAt.In(time.FixedZone("Asia/Tokyo", 9*60*60))
	endOfIssuedAtAsUserTZ := time.Date(issuedAtAsUserTZ.Year(), issuedAtAsUserTZ.Month(), issuedAtAsUserTZ.Day(), 23, 59, 59, 0, time.FixedZone("Asia/Tokyo", 9*60*60))
	dueAtAsUserTZ := endOfIssuedAtAsUserTZ.AddDate(0, 0, 14)
	i.paymentDueAt = dueAtAsUserTZ.UTC()
}
