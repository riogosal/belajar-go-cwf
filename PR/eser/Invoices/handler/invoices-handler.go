package handler

import (
	"chenwoo-simple-invoices/domain"
	"chenwoo-simple-invoices/models"
	"context"

	"github.com/gin-gonic/gin"
)

type InvoicesHandler struct {
	Invoices domain.Invoices
}

func NewInvoicesHandler(invoices domain.Invoices) InvoicesHandler {
	return InvoicesHandler{
		Invoices: invoices,
	}
}

func (i *InvoicesHandler) GetData(ctx *gin.Context) {
	dataInvoices, err := i.Invoices.GetData(ctx)

	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal server Error"})
		return
	}

	ctx.JSON(200, gin.H{"data": dataInvoices})
}

func (i *InvoicesHandler) GetDataByID(ctx *gin.Context) {
	id := ctx.Param("id")

	idInvoices, err := i.Invoices.GetDataID(ctx, id)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	ctx.JSON(200, idInvoices)
}

func (i *InvoicesHandler) CreateData(ctx *gin.Context) {
	var dataInvoices models.Invoices

	if err := ctx.ShouldBindJSON(&dataInvoices); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	err := i.Invoices.CreateData(context.Background(), &dataInvoices)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
	}

	ctx.JSON(200, gin.H{"message": "Success", "Data": dataInvoices})
}

func (i *InvoicesHandler) DeleteDataID(ctx *gin.Context) {
	id := ctx.Param("id")

	err := i.Invoices.DeleteData(ctx, id)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal server Error"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Data Berasil Dihapus!", "status": "success"})
}
func (i *InvoicesHandler) InvoicesRoutes(rg *gin.RouterGroup) {
	rg.GET("/invoices", i.GetData)
	rg.GET("/get/:id", i.GetDataByID)
	rg.POST("/create", i.CreateData)
	rg.DELETE("/delete/:id", i.DeleteDataID)
}
