package invoice_test

import (
	"testing"
	"time"
	"upsider-coding-test/domain/invoice"
	shared_mock "upsider-coding-test/mock/shared"

	"go.uber.org/mock/gomock"
)

func Test_InvoiceFactory_Issue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	clock := shared_mock.NewMockClock(ctrl)
	factory := invoice.ProvideInvoiceFactory(clock)

	type expected struct {
		companyID          string
		partnerID          string
		issuedAt           time.Time
		paymentAmount      string
		fee                string
		feeRate            string
		consumptionTax     string
		consumptionTaxRate string
		invoiceAmount      string
		paymentDueAt       time.Time
		Status             invoice.Status
	}

	t.Run("支払い金額が10000で発行できる", func(t *testing.T) {
		t.Parallel()
		// Setup
		now := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		due := time.Date(2021, 1, 15, 14, 59, 59, 0, time.UTC)
		clock.EXPECT().Now().Return(now)
		want := expected{
			companyID:          "company_id",
			partnerID:          "partner_id",
			issuedAt:           now,
			paymentAmount:      "10000",
			fee:                "400",
			feeRate:            "0.04",
			consumptionTax:     "40",
			consumptionTaxRate: "0.1",
			invoiceAmount:      "10440",
			paymentDueAt:       due,
			Status:             invoice.Unpaid,
		}
		// Exercise
		got, err := factory.Issue(&invoice.IssueInput{
			CompanyID:     "company_id",
			PartnerID:     "partner_id",
			PaymentAmount: "10000",
		})
		// Verify
		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
		}
		if got.CompanyID().String() != want.companyID {
			t.Errorf("会社IDが異なります: %s", got.CompanyID())
		}
		if got.PartnerID().String() != want.partnerID {
			t.Errorf("取引先IDが異なります: %s", got.PartnerID())
		}
		if got.IssuedAt() != want.issuedAt {
			t.Errorf("発行日時が異なります: %s", got.IssuedAt())
		}
		if got.PaymentAmount().String() != want.paymentAmount {
			t.Errorf("支払い金額が異なります: %s", got.PaymentAmount().String())
		}
		if got.Fee().Value().String() != want.fee {
			t.Errorf("手数料が異なります: %s", got.Fee().Value().String())
		}
		if got.Fee().Rate().String() != want.feeRate {
			t.Errorf("手数料率が異なります: %s ", got.Fee().Rate().String())
		}
		if got.ConsumptionTax().Value().String() != want.consumptionTax {
			t.Errorf("消費税が異なります: %s", got.ConsumptionTax().Value().String())
		}
		if got.ConsumptionTax().Rate().String() != want.consumptionTaxRate {
			t.Errorf("消費税率が異なります: %s ", got.ConsumptionTax().Rate().String())
		}
		if got.InvoiceAmount().String() != want.invoiceAmount {
			t.Errorf("請求金額が異なります: %s", got.InvoiceAmount().String())
		}
		if got.PaymentDueAt() != want.paymentDueAt {
			t.Errorf("支払期限が異なります: %s", got.PaymentDueAt())
		}
		if got.Status() != want.Status {
			t.Errorf("ステータスが異なります: %d", got.Status())
		}
	})
	t.Run("支払い金額が10の場合: 手数料0.4,消費税0.04,請求額10.44", func(t *testing.T) {
		t.Parallel()
		// Setup
		now := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		clock.EXPECT().Now().Return(now)
		want := expected{
			fee:            "0.4",
			consumptionTax: "0.04",
			invoiceAmount:  "10.44",
		}
		// Exercise
		got, err := factory.Issue(&invoice.IssueInput{
			CompanyID:     "company_id",
			PartnerID:     "partner_id",
			PaymentAmount: "10",
		})
		// Verify
		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
		}
		if got.Fee().Value().String() != want.fee {
			t.Errorf("手数料が異なります: %s", got.Fee().Value().String())
		}
		if got.ConsumptionTax().Value().String() != want.consumptionTax {
			t.Errorf("消費税が異なります: %s", got.ConsumptionTax().Value().String())
		}
		if got.InvoiceAmount().String() != want.invoiceAmount {
			t.Errorf("請求金額が異なります: %s", got.InvoiceAmount().String())
		}
	})
	t.Run("支払い金額が123456.78の場合: 手数料4938.2712,消費税493.82712,請求額128888.87832", func(t *testing.T) {
		t.Parallel()
		// Setup
		now := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		clock.EXPECT().Now().Return(now)
		want := expected{
			fee:            "4938.2712",
			consumptionTax: "493.82712",
			invoiceAmount:  "128888.87832",
		}
		// Exercise
		got, err := factory.Issue(&invoice.IssueInput{
			CompanyID:     "company_id",
			PartnerID:     "partner_id",
			PaymentAmount: "123456.78",
		})
		// Verify
		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
		}
		if got.Fee().Value().String() != want.fee {
			t.Errorf("手数料が異なります: %s", got.Fee().Value().String())
		}
		if got.ConsumptionTax().Value().String() != want.consumptionTax {
			t.Errorf("消費税が異なります: %s", got.ConsumptionTax().Value().String())
		}
		if got.InvoiceAmount().String() != want.invoiceAmount {
			t.Errorf("請求金額が異なります: %s", got.InvoiceAmount().String())
		}
	})
	t.Run("支払い金額が-1円の場合エラー", func(t *testing.T) {
		t.Parallel()
		// Setup
		now := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		clock.EXPECT().Now().Return(now)
		// Exercise
		_, err := factory.Issue(&invoice.IssueInput{
			CompanyID:     "company_id",
			PartnerID:     "partner_id",
			PaymentAmount: "-1",
		})
		// Verify
		if err == nil {
			t.Error("エラーが発生しませんでした")
		}
	})
}
