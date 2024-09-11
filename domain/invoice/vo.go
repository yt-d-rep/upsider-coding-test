package invoice

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"upsider-coding-test/shared"
)

type (
	InvoiceID string
	Amount    decimal.Decimal
	Rate      decimal.Decimal
	Status    int
	Fee       struct {
		value Amount
		rate  Rate
	}
	ConsumptionTax struct {
		value Amount
		rate  Rate
	}
)

const (
	CurrentFeeRate            string = "0.04"
	CurrentConsumptionTaxRate string = "0.10"
)

func NewInvoiceID() InvoiceID {
	return InvoiceID(uuid.New().String())
}
func ParseInvoiceID(id string) (InvoiceID, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return "", &shared.ArgumentError{Field: "invoiceID", Err: "invalid invoice id"}
	}
	return InvoiceID(id), nil
}
func (id InvoiceID) String() string {
	return string(id)
}

func NewAmount(value string) (Amount, error) {
	amount, err := decimal.NewFromString(value)
	if err != nil {
		return Amount{}, &shared.ArgumentError{Field: "amount", Err: "invalid amount"}
	}
	if amount.LessThan(decimal.Zero) {
		return Amount{}, &shared.ArgumentError{Field: "amount", Err: "amount must be greater than or equal to 0"}
	}
	return Amount(amount), nil
}
func (a Amount) String() string {
	return a.asDecimal().String()
}
func (a Amount) asDecimal() decimal.Decimal {
	return decimal.Decimal(a)
}
func (a Amount) MulRate(rate Rate) Amount {
	amount := a.asDecimal().Mul(rate.asDecimal())
	return Amount(amount)
}
func (a Amount) Add(amount Amount) Amount {
	return Amount(a.asDecimal().Add(amount.asDecimal()))
}

func NewRate(value string) (Rate, error) {
	rate, err := decimal.NewFromString(value)
	if err != nil {
		return Rate{}, &shared.ArgumentError{Field: "rate", Err: "invalid rate"}
	}
	return Rate(rate), nil
}
func (r Rate) String() string {
	return r.asDecimal().String()
}
func (r Rate) asDecimal() decimal.Decimal {
	return decimal.Decimal(r)
}

func NewFee(value Amount, rate Rate) *Fee {
	return &Fee{value: value, rate: rate}
}
func ParseFee(value string, rate string) (*Fee, error) {
	amount, err := NewAmount(value)
	if err != nil {
		return &Fee{}, err
	}
	rateValue, err := NewRate(rate)
	if err != nil {
		return &Fee{}, err
	}
	return &Fee{value: amount, rate: rateValue}, nil
}
func (f Fee) Value() Amount {
	return f.value
}
func (f Fee) Rate() Rate {
	return f.rate
}

func NewConsumptionTax(value Amount, rate Rate) *ConsumptionTax {
	return &ConsumptionTax{value: value, rate: rate}
}
func ParseConsumptionTax(value string, rate string) (*ConsumptionTax, error) {
	amount, err := NewAmount(value)
	if err != nil {
		return &ConsumptionTax{}, err
	}
	rateValue, err := NewRate(rate)
	if err != nil {
		return &ConsumptionTax{}, err
	}
	return &ConsumptionTax{value: amount, rate: rateValue}, nil
}
func (c ConsumptionTax) Value() Amount {
	return c.value
}
func (c ConsumptionTax) Rate() Rate {
	return c.rate
}

const (
	Unpaid Status = iota
	Processing
	Paid
	Error
)

func (s Status) String() string {
	switch s {
	case Unpaid:
		return "未払い"
	case Processing:
		return "処理中"
	case Paid:
		return "支払い済み"
	case Error:
		return "エラー"
	}
	return ""
}

func NewStatus(status int) Status {
	return Status(status)
}
