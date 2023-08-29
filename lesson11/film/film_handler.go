package film

import (
	"interfaces/domain"

	"github.com/gin-gonic/gin"
)

type GinFilmHandler struct {
	Repo *domain.FilmRepository
}

func (h *GinFilmHandler) GetAllFilms(c *gin.Context) {

}
