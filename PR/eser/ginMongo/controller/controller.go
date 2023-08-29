package controller

import (
	"fmt"
	"go-mongo/model"
	"go-mongo/repository"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type FilmController struct {
	Repository repository.FilmRepository
}

func NewFilmController(repository repository.FilmRepository) *FilmController {
	return &FilmController{Repository: repository}
}

func (c *FilmController) GetFilmByID(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Parameter 'id' harus berupa angka"})
		return
	}

	film, err := c.Repository.FindByID(ctx, id)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if film != nil {
		ctx.JSON(200, gin.H{"status": "sukses", "data": film})
	} else {
		ctx.JSON(404, gin.H{"error": fmt.Sprintf("Film dengan ID %d tidak ditemukan", id)})
	}
}

func (c *FilmController) GetDataFilm(ctx *gin.Context) {
	films, err := c.Repository.FindAll(ctx)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if len(films) > 0 {
		ctx.JSON(200, gin.H{"status": "sukses", "data": films})
	} else {
		ctx.JSON(404, gin.H{"error": "Tidak ada data film yang tersedia"})
	}
}

func (c *FilmController) AddFilm(ctx *gin.Context) {
	var film model.Film

	if err := ctx.BindJSON(&film); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := c.Repository.AddFilm(ctx, &film)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"message": "Film berhasil ditambahkan", "data": result.InsertedID})
}

func (c *FilmController) Delete(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Parameter 'id' harus berupa angka"})
		return
	}

	err = c.Repository.Delete(ctx, id)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(404, gin.H{"message": "Data tidak ditemukan"})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Data berhasil dihapus"})
}
