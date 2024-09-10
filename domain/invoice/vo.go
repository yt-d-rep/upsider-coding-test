package invoice

import (
	"errors"
	"math/big"

	"github.com/google/uuid"

	"upsider-base/shared"
)

type (
	InvoiceID          string
	Amount             int64
	FeeRate            big.Rat
	ConsumptionTaxRate big.Rat
	Status             int
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

func NewAmount(amount int64) (Amount, error) {
	if amount < 0 {
		return 0, &shared.ValidationError{Field: "amount", Err: "amount must be greater than or equal to 0"}
	}
	return Amount(amount), nil
}
func (a Amount) Int64() int64 {
	return int64(a)
}
func (a Amount) toRat() *big.Rat {
	return big.NewRat(a.Int64(), 1)
}
func (a Amount) IsZero() bool {
	return a == 0
}

func NewFeeRate() (FeeRate, error) {
	r := new(big.Rat)
	rat, ok := r.SetString("0.03")
	if !ok {
		return FeeRate(*new(big.Rat)), errors.New("invalid fee rate")
	}
	return FeeRate(*rat), nil
}
func ParseFeeRate(rate string) (FeeRate, error) {
	r := new(big.Rat)
	rat, ok := r.SetString(rate)
	if !ok {
		return FeeRate(*rat), &shared.ArgumentError{Field: "fee_rate", Err: "invalid fee rate"}
	}
	return FeeRate(*r), nil
}
func (f FeeRate) toRat() *big.Rat {
	return (*big.Rat)(&f)
}
func (f FeeRate) String() string {
	return f.toRat().FloatString(2)
}

func NewConsumptionTaxRate() (ConsumptionTaxRate, error) {
	r := new(big.Rat)
	rat, ok := r.SetString("0.10")
	if !ok {
		return ConsumptionTaxRate(*new(big.Rat)), errors.New("invalid consumption tax rate")
	}
	return ConsumptionTaxRate(*rat), nil
}
func ParseConsumptionTaxRate(rate string) (ConsumptionTaxRate, error) {
	r := new(big.Rat)
	rat, ok := r.SetString(rate)
	if !ok {
		return ConsumptionTaxRate(*new(big.Rat)), &shared.ArgumentError{Field: "consumption_tax_rate", Err: "invalid consumption tax rate"}
	}
	return ConsumptionTaxRate(*rat), nil
}
func (c ConsumptionTaxRate) toRat() *big.Rat {
	return (*big.Rat)(&c)
}
func (c ConsumptionTaxRate) String() string {
	return c.toRat().FloatString(2)
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
