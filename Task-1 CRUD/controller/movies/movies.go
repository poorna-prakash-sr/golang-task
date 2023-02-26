package movies

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/movies")
	routes.POST("/", h.addMovie)
	routes.GET("/", h.getMovie)
	routes.GET("/:id", h.getMovieById)
	routes.PUT("/:id", h.updateMovieById)
	routes.DELETE("/:id", h.deleteMovie)
}
