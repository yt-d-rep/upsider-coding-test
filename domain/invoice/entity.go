package invoice

import (
	"math/big"
	"time"
	"upsider-coding-test/domain/company"
)

type (
	Invoice struct {
		id                 InvoiceID
		companyID          company.CompanyID
		partnerID          company.PartnerID
		issuedAt           time.Time
		paymentAmount      Amount
		fee                Amount
		feeRate            FeeRate
		consumptionTax     Amount
		consumptionTaxRate ConsumptionTaxRate
		invoiceAmount      Amount
		paymentDueAt       time.Time
		status             Status
	}
	ParseInvoiceInput struct {
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
	fee, err := NewAmount(i.Fee)
	if err != nil {
		return nil, err
	}
	feeRate, err := ParseFeeRate(i.FeeRate)
	if err != nil {
		return nil, err
	}
	consumptionTax, err := NewAmount(i.ConsumptionTax)
	if err != nil {
		return nil, err
	}
	consumptionTaxRate, err := ParseConsumptionTaxRate(i.ConsumptionTaxRate)
	if err != nil {
		return nil, err
	}
	invoiceAmount, err := NewAmount(i.InvoiceAmount)
	if err != nil {
		return nil, err
	}
	status := Status(i.Status)
	return &Invoice{
		id:                 id,
		companyID:          companyID,
		partnerID:          partnerID,
		issuedAt:           i.IssuedAt,
		paymentAmount:      paymentAmount,
		fee:                fee,
		feeRate:            feeRate,
		consumptionTax:     consumptionTax,
		consumptionTaxRate: consumptionTaxRate,
		invoiceAmount:      invoiceAmount,
		paymentDueAt:       i.PaymentDueAt,
		status:             status,
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
func (i *Invoice) Fee() Amount {
	return i.fee
}
func (i *Invoice) FeeRate() FeeRate {
	return i.feeRate
}
func (i *Invoice) ConsumptionTax() Amount {
	return i.consumptionTax
}
func (i *Invoice) ConsumptionTaxRate() ConsumptionTaxRate {
	return i.consumptionTaxRate
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

func (i *Invoice) calculateFee() error {
	resultRat := new(big.Rat).Mul(i.paymentAmount.toRat(), i.feeRate.toRat())
	fee, err := NewAmount(resultRat.Num().Int64())
	if err != nil {
		return err
	}
	i.fee = fee
	return nil
}
func (i *Invoice) calculateConsumptionTax() error {
	resultRat := new(big.Rat).Mul(i.paymentAmount.toRat(), i.consumptionTaxRate.toRat())
	consumptionTax, err := NewAmount(resultRat.Num().Int64())
	if err != nil {
		return err
	}
	i.consumptionTax = consumptionTax
	return nil
}
func (i *Invoice) calculateInvoiceAmount() error {
	feeAdded := new(big.Rat).Add(i.paymentAmount.toRat(), i.fee.toRat())
	consumptionTaxAdded := new(big.Rat).Add(feeAdded, i.consumptionTax.toRat())
	invoiceAmount, err := NewAmount(consumptionTaxAdded.Num().Int64())
	if err != nil {
		return err
	}
	i.invoiceAmount = invoiceAmount
	return nil
}
func (i *Invoice) setDue() {
	// 請求書発行日の14日後の23:59:59までが支払い期限
	// TODO: クライアントのTZはAsia/Tokyoに固定しているが、本来はユーザーのTZに合わせる
	issuedAtAsUserTZ := i.issuedAt.In(time.FixedZone("Asia/Tokyo", 9*60*60))
	endOfIssuedAtAsUserTZ := time.Date(issuedAtAsUserTZ.Year(), issuedAtAsUserTZ.Month(), issuedAtAsUserTZ.Day(), 23, 59, 59, 0, time.FixedZone("Asia/Tokyo", 9*60*60))
	dueAtAsUserTZ := endOfIssuedAtAsUserTZ.AddDate(0, 0, 14)
	i.paymentDueAt = dueAtAsUserTZ.UTC()
}
