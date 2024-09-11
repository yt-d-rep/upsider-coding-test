package handler

import (
	"time"
	"upsider-coding-test/domain/invoice"
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
	InvoiceIssueParams struct {
		CompanyID     string `json:"company_id" binding:"required"`
		PartnerID     string `json:"partner_id" binding:"required"`
		PaymentAmount string `json:"payment_amount" binding:"required"`
	}
	InvoiceListBetweenParams struct {
		From      string `form:"from" binding:"required"`
		To        string `form:"to" binding:"required"`
		CompanyID string `form:"company_id" binding:"required"`
	}
)

func (h *invoiceHandler) Issue(ctx *gin.Context) {
	var params InvoiceIssueParams
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
		ctx.JSON(200, convertInvoice(invoice))
		return
	}
	handleError(ctx, err)
}

func (h *invoiceHandler) ListBetween(ctx *gin.Context) {
	var params InvoiceListBetweenParams
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
			res = append(res, convertInvoice(invoice))
		}
		ctx.JSON(200, res)
		return
	}
	handleError(ctx, err)
}

func convertInvoice(invoice *invoice.Invoice) gin.H {
	return gin.H{
		"id":                   invoice.ID().String(),
		"company_id":           invoice.CompanyID().String(),
		"partner_id":           invoice.PartnerID().String(),
		"issued_at":            invoice.IssuedAt(),
		"payment_amount":       invoice.PaymentAmount().String(),
		"fee":                  invoice.Fee().Value().String(),
		"fee_rate":             invoice.Fee().Rate().String(),
		"consumption_tax":      invoice.ConsumptionTax().Value().String(),
		"consumption_tax_rate": invoice.ConsumptionTax().Rate().String(),
		"invoice_amount":       invoice.InvoiceAmount().String(),
		"payment_due_at":       invoice.PaymentDueAt(),
		"status":               invoice.Status().String(),
	}
}
