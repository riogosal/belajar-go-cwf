package handler

import (
	"chenwoo-simple-invoices/domain"
	"chenwoo-simple-invoices/models"

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

func (i *InvoicesHandler) DeleteDataContent(ctx *gin.Context) {
	invoiceUUID := ctx.Param("inv_uuid")
	contentUUID := ctx.Param("contents_uuid")

	// fmt.Println("Content UUID", contentUUID)
	// fmt.Println("Invoice UUID", invoiceUUID)

	err := i.Invoices.DeleteDataContent(ctx, invoiceUUID, contentUUID)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	ctx.JSON(200, gin.H{"message": "data berhasil dihapus!", "status": "success"})
}

func (i *InvoicesHandler) AddDataProduct(ctx *gin.Context) {
	invoiceUUID := ctx.Param("invoice_uuid")

	var newProduct models.Content

	if err := ctx.ShouldBindJSON(&newProduct); err != nil {
		ctx.JSON(400, gin.H{"message": "bad request"})
		return
	}

	err := i.Invoices.AddDataConten(ctx, invoiceUUID, &newProduct)
	if err != nil {
		ctx.JSON(404, gin.H{"message": "data tidak UUID ditemukan"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Produk berhasil di tambahkan", "status": "success"})
}

func (i *InvoicesHandler) GetData(ctx *gin.Context) {
	dataInvoices, err := i.Invoices.GetData(ctx)

	if err != nil {
		ctx.JSON(500, gin.H{"message": "internal server error"})
		return
	}

	ctx.JSON(200, gin.H{"data": dataInvoices})
}

func (i *InvoicesHandler) GetDataByID(ctx *gin.Context) {
	id := ctx.Param("id")

	idInvoices, err := i.Invoices.GetDataID(ctx, id)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "bad request"})
		return
	}
	ctx.JSON(200, gin.H{"data": idInvoices})
}

func (i *InvoicesHandler) CreateData(ctx *gin.Context) {
	var dataInvoices models.Invoices

	if err := ctx.ShouldBindJSON(&dataInvoices); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	err := i.Invoices.CreateData(ctx, &dataInvoices)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
	}

	ctx.JSON(200, gin.H{"data": dataInvoices, "status": "success"})
}

func (i *InvoicesHandler) DeleteDataID(ctx *gin.Context) {
	id := ctx.Param("id")

	err := i.Invoices.DeleteData(ctx, id)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "bad request"})
		return
	}
	ctx.JSON(200, gin.H{"message": "data berhasil dihapus!", "status": "success"})
}

func (i *InvoicesHandler) InvoicesRoutes(rg *gin.RouterGroup) {
	rg.GET("/", i.GetData)
	rg.GET("/get/:id", i.GetDataByID)
	rg.POST("/create", i.CreateData)
	rg.DELETE("/delete/:id", i.DeleteDataID)
	rg.POST("/:invoice_uuid/contents", i.AddDataProduct)
	rg.DELETE("/:inv_uuid/contents/:contents_uuid/invoices", i.DeleteDataContent)
}
