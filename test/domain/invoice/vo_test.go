package invoice_test

import (
	"math/big"
	"reflect"
	"testing"
	"upsider-base/domain/invoice"
)

func Test_NewFeeRate(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		want := big.NewRat(3, 100)
		wantString := "0.03"
		got, err := invoice.NewFeeRate()
		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
		}
		if reflect.DeepEqual(got, want) {
			t.Errorf("手数料率が異なります: %s", got.String())
		}
		if got.String() != wantString {
			t.Errorf("手数料率が異なります: %s", got.String())
		}
	})
}

func Test_NewConsumptionTaxRate(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		want := big.NewRat(10, 100)
		wantString := "0.10"
		got, err := invoice.NewConsumptionTaxRate()
		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
		}
		if reflect.DeepEqual(got, want) {
			t.Errorf("消費税率が異なります: %s", got.String())
		}
		if got.String() != wantString {
			t.Errorf("消費税率が異なります: %s", got.String())
		}
	})
}
