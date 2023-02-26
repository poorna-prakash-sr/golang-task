package movies

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"crud-movies-api/model"
)

type RequestBody struct {
	Title      string `json:"title" binding:"required"`
	Director   string `json:"director" binding:"required"`
	Hero       string `json:"hero" binding:"required"`
	RelaseDate string `json:"reasedate" binding:"required"`
}

func (h handler) addMovie(ctx *gin.Context) {
	body := RequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "please provide a valid inputs",
		})
		return
	}

	var movie model.Movie

	movie.Title = body.Title
	movie.Director = body.Director
	movie.Hero = body.Hero
	movie.RelaseDate = body.RelaseDate

	if result := h.DB.Create(&movie); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &movie)
}

func (h handler) getMovie(ctx *gin.Context) {
	var movies []model.Movie

	if result := h.DB.Find(&movies); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	if len(movies) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "No movies found",
		})
		return
	}
	ctx.JSON(http.StatusOK, &movies)
}

func (h handler) getMovieById(ctx *gin.Context) {
	id := ctx.Param("id")

	var movie []model.Movie

	if result := h.DB.First(&movie, id); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "invalid id",
		})
		return
	}

	ctx.JSON(http.StatusOK, &movie)
}

func (h handler) updateMovieById(ctx *gin.Context) {
	id := ctx.Param("id")
	body := RequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var movie model.Movie

	if result := h.DB.First(&movie, id); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "please provide a valid inputs",
		})
		return
	}

	movie.Title = body.Title
	movie.Director = body.Director
	movie.Hero = body.Hero
	movie.RelaseDate = body.RelaseDate

	h.DB.Save(&movie)

	ctx.JSON(http.StatusOK, &movie)
}

func (h handler) deleteMovie(ctx *gin.Context) {
	id := ctx.Param("id")

	var movie model.Movie

	if result := h.DB.First(&movie, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&movie)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Movie deleted successfully" + id,
	})
}
