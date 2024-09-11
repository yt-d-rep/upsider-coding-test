package e2e_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
	"upsider-coding-test/infrastructure/handler"
	"upsider-coding-test/infrastructure/server"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIssueAndList(t *testing.T) {
	// Setup
	router := gin.Default()
	server.Route(router)
	now := time.Now()

	// Issue#Setup
	w := httptest.NewRecorder()
	issueInvoice := handler.InvoiceIssueParams{
		CompanyID:     "b8e7fce5-77a5-4e64-9e3c-90e0c5b4c17d",
		PartnerID:     "a5b6f8d4-9b44-4f9e-919f-d5cb2d7b8e9f",
		PaymentAmount: "10000",
	}
	payload, err := json.Marshal(issueInvoice)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/api/invoices", strings.NewReader(string(payload)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+"TOKEN")
	// Issue#Execise
	router.ServeHTTP(w, req)
	// Issue#Verify
	assert.Equal(t, 200, w.Code)
	respBody := make(map[string]interface{})
	if err = json.Unmarshal(w.Body.Bytes(), &respBody); err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, respBody["id"])
	assert.Equal(t, issueInvoice.CompanyID, respBody["company_id"])
	assert.Equal(t, issueInvoice.PartnerID, respBody["partner_id"])
	assert.Equal(t, "10000", respBody["payment_amount"])
	issuedAt, err := time.Parse(time.RFC3339, respBody["issued_at"].(string))
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, now.Before(issuedAt))
	assert.Equal(t, "400", respBody["fee"])
	assert.Equal(t, "0.04", respBody["fee_rate"])
	assert.Equal(t, "40", respBody["consumption_tax"])
	assert.Equal(t, "0.1", respBody["consumption_tax_rate"])
	assert.Equal(t, "10440", respBody["invoice_amount"])
	paymentDueAt, err := time.Parse(time.RFC3339, respBody["payment_due_at"].(string))
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, now.AddDate(0, 0, 14).Before(paymentDueAt))
	assert.Equal(t, "未払い", respBody["status"])

	// List#Setup
	w = httptest.NewRecorder()
	params := url.Values{}
	params.Add("company_id", issueInvoice.CompanyID)
	params.Add("from", now.Format(time.RFC3339))
	params.Add("to", now.AddDate(0, 0, 1).Format(time.RFC3339))
	req, err = http.NewRequest("GET", "/api/invoices?"+params.Encode(), nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+"TOKEN")
	// List#Execise
	router.ServeHTTP(w, req)
	// List#Verify
	assert.Equal(t, 200, w.Code)
	listRespBody := make([]map[string]interface{}, 0)
	if err = json.Unmarshal(w.Body.Bytes(), &listRespBody); err != nil {
		t.Fatal(err)
	}
	for _, invoice := range listRespBody {
		assert.Equal(t, issueInvoice.CompanyID, invoice["company_id"])
		assert.Equal(t, issueInvoice.PartnerID, invoice["partner_id"])
		assert.Equal(t, "10000", invoice["payment_amount"])
		issuedAt, err := time.Parse(time.RFC3339, invoice["issued_at"].(string))
		if err != nil {
			t.Fatal(err)
		}
		assert.True(t, now.Before(issuedAt))
		assert.Equal(t, "400", invoice["fee"])
		assert.Equal(t, "0.04", invoice["fee_rate"])
		assert.Equal(t, "40", invoice["consumption_tax"])
		assert.Equal(t, "0.1", invoice["consumption_tax_rate"])
		assert.Equal(t, "10440", invoice["invoice_amount"])
		paymentDueAt, err := time.Parse(time.RFC3339, invoice["payment_due_at"].(string))
		if err != nil {
			t.Fatal(err)
		}
		assert.True(t, now.AddDate(0, 0, 14).Before(paymentDueAt))
		assert.Equal(t, "未払い", invoice["status"])
	}
}
