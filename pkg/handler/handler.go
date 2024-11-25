package handler

import (
	"github.com/EvgeniyGnetnev/GenreID/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("templates/*")

	router.Static("/assets", "./assets")
	router.Static("/wwwroot", "./wwwroot")

	api := router.Group("/api", h.Index) // сделать для главной страницы, без /api
	{
		artists := api.Group("/artists", h.ArtistsMainPage)
		{
			artists.POST("/", h.CreateArtist)
			artists.GET("/", h.Index) // для теста добавил обработчик главной страницы
			artists.PUT("/:id", h.UpdateArtist)
			artists.DELETE("/:id", h.DeleteArtist)
		}
		songs := api.Group("/songs", h.SongsMainPage)
		{
			songs.POST("/", h.CreateSong)
			songs.GET("/:song_id", h.SongPage)
			songs.PUT("/:song_id", h.UpdateSong)
			songs.DELETE("/:song_id", h.DeleteSong)
		}
		genres := api.Group("/genres", h.GenresMainPage)
		{
			genres.POST("/", h.CreateGenre)
			genres.GET("/:genre_id", h.GenrePage)
			genres.PUT("/:genre_id", h.UpdateGenre)
			genres.DELETE("/:genre_id", h.DeleteGenre)
		}
	}

	return router
}
