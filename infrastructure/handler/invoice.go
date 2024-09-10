package handler

import (
	"time"
	"upsider-coding-test/usecase"

	"github.com/gin-gonic/gin"
)

type (
	InvoiceHandler interface {
		Issue(ctx *gin.Context)
		ListBetween(ctx *gin.Context)
	}
	invoiceHandler struct {
		invoiceUsecase usecase.InvoiceUsecase
	}
	invoiceIssueParams struct {
		CompanyID     string `json:"company_id" binding:"required"`
		PartnerID     string `json:"partner_id" binding:"required"`
		PaymentAmount int64  `json:"payment_amount" binding:"required"`
	}
	invoiceListBetweenParams struct {
		From      string `form:"from" binding:"required"`
		To        string `form:"to" binding:"required"`
		CompanyID string `form:"company_id" binding:"required"`
	}
)

func (h *invoiceHandler) Issue(ctx *gin.Context) {
	var params invoiceIssueParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	invoice, err := h.invoiceUsecase.Issue(&usecase.IssueInput{
		CompanyID:     params.CompanyID,
		PartnerID:     params.PartnerID,
		PaymentAmount: params.PaymentAmount,
	})
	if err == nil {
		ctx.JSON(200, gin.H{
			"id":                   invoice.ID().String(),
			"company_id":           invoice.CompanyID().String(),
			"partner_id":           invoice.PartnerID().String(),
			"issued_at":            invoice.IssuedAt(),
			"payment_amount":       invoice.PaymentAmount(),
			"fee":                  invoice.Fee(),
			"fee_rate":             invoice.FeeRate().String(),
			"consumption_tax":      invoice.ConsumptionTax(),
			"consumption_tax_rate": invoice.ConsumptionTaxRate().String(),
			"invoice_amount":       invoice.InvoiceAmount(),
			"payment_due_at":       invoice.PaymentDueAt(),
			"status":               invoice.Status().String(),
		})
		return
	}
	handleError(ctx, err)
}

func (h *invoiceHandler) ListBetween(ctx *gin.Context) {
	var params invoiceListBetweenParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	from, err := time.Parse(time.RFC3339, params.From)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "invalid from"})
		return
	}
	to, err := time.Parse(time.RFC3339, params.To)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "invalid to"})
		return
	}
	invoice, err := h.invoiceUsecase.ListBetween(&usecase.ListBetweenInput{
		CompanyID: params.CompanyID,
		From:      from,
		To:        to,
	})
	if err == nil {
		res := make([]gin.H, 0, len(invoice))
		for _, invoice := range invoice {
			res = append(res, gin.H{
				"id":                   invoice.ID().String(),
				"company_id":           invoice.CompanyID().String(),
				"partner_id":           invoice.PartnerID().String(),
				"issued_at":            invoice.IssuedAt(),
				"payment_amount":       invoice.PaymentAmount(),
				"fee":                  invoice.Fee(),
				"fee_rate":             invoice.FeeRate().String(),
				"consumption_tax":      invoice.ConsumptionTax(),
				"consumption_tax_rate": invoice.ConsumptionTaxRate().String(),
				"invoice_amount":       invoice.InvoiceAmount(),
				"payment_due_at":       invoice.PaymentDueAt(),
				"status":               invoice.Status().String(),
			})
		}
		ctx.JSON(200, res)
		return
	}
	handleError(ctx, err)
}
