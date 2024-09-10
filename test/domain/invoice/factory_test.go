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
		paymentAmount      int64
		fee                int64
		feeRate            string
		consumptionTax     int64
		consumptionTaxRate string
		invoiceAmount      int64
		paymentDueAt       time.Time
		Status             invoice.Status
	}

	t.Run("支払い金額が1000円で発行できる", func(t *testing.T) {
		t.Parallel()
		// Setup
		now := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		due := time.Date(2021, 1, 15, 14, 59, 59, 0, time.UTC)
		clock.EXPECT().Now().Return(now)
		want := expected{
			companyID:          "company_id",
			partnerID:          "partner_id",
			issuedAt:           now,
			paymentAmount:      1000,
			fee:                30,
			feeRate:            "0.03",
			consumptionTax:     100,
			consumptionTaxRate: "0.10",
			invoiceAmount:      1130,
			paymentDueAt:       due,
			Status:             invoice.Unpaid,
		}
		// Exercise
		got, err := factory.Issue(&invoice.IssueInput{
			CompanyID:     "company_id",
			PartnerID:     "partner_id",
			PaymentAmount: 1000,
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
		if got.PaymentAmount().Int64() != want.paymentAmount {
			t.Errorf("支払い金額が異なります: %d", got.PaymentAmount().Int64())
		}
		if got.Fee().Int64() != want.fee {
			t.Errorf("手数料が異なります: %d", got.Fee().Int64())
		}
		if got.FeeRate().String() != want.feeRate {
			t.Errorf("手数料率が異なります: %s", got.FeeRate().String())
		}
		if got.ConsumptionTax().Int64() != want.consumptionTax {
			t.Errorf("消費税が異なります: %d", got.ConsumptionTax().Int64())
		}
		if got.ConsumptionTaxRate().String() != want.consumptionTaxRate {
			t.Errorf("消費税率が異なります: %s", got.ConsumptionTaxRate().String())
		}
		if got.InvoiceAmount().Int64() != want.invoiceAmount {
			t.Errorf("請求金額が異なります: %d", got.InvoiceAmount().Int64())
		}
		if got.PaymentDueAt() != want.paymentDueAt {
			t.Errorf("支払期限が異なります: %s", got.PaymentDueAt())
		}
		if got.Status() != want.Status {
			t.Errorf("ステータスが異なります: %d", got.Status())
		}
	})
	t.Run("支払い金額が0円の場合エラー", func(t *testing.T) {
		t.Parallel()
		// Setup
		now := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		clock.EXPECT().Now().Return(now)
		// Exercise
		_, err := factory.Issue(&invoice.IssueInput{
			CompanyID:     "company_id",
			PartnerID:     "partner_id",
			PaymentAmount: 0,
		})
		// Verify
		if err == nil {
			t.Error("エラーが発生しませんでした")
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
			PaymentAmount: -1,
		})
		// Verify
		if err == nil {
			t.Error("エラーが発生しませんでした")
		}
	})
}
