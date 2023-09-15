package handler

import (
	"app-api-natulius/domain"
	"app-api-natulius/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InvoiceHandler struct {
	InvoiceService domain.InvoiceService
}

func NewInvoiceHandler(invoiceService domain.InvoiceService) InvoiceHandler {
	return InvoiceHandler{InvoiceService: invoiceService}
}

func (ih *InvoiceHandler) GetData(ctx *gin.Context) {
	id := ctx.Param("id")
	data, err := ih.InvoiceService.GetData(context.Background(), &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": data})

}

func (ih *InvoiceHandler) CreateData(ctx *gin.Context) {
	var data models.Invoices
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := ih.InvoiceService.CreateData(context.Background(), &data)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "created successfully", "data": data})
}

func (ih *InvoiceHandler) GetAll(ctx *gin.Context) {
	data, err := ih.InvoiceService.GetAll(context.Background())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Success", "data": data})
}

func (ih *InvoiceHandler) Delete(ctx *gin.Context) {
	idInvoice := ctx.Param("id")
	err := ih.InvoiceService.Delete(context.Background(), &idInvoice)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Success Delete"})
}

func (ih *InvoiceHandler) InvoiceRoutes(rg *gin.RouterGroup) {
	rg.GET("/invoice/:id", ih.GetData)
	rg.POST("/invoice", ih.CreateData)
	rg.GET("/invoice", ih.GetAll)
	rg.DELETE("/invoice/:id", ih.Delete)
}
