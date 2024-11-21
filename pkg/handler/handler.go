package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api") // сделать для главной страницы, без /api
	{
		artists := api.Group("/artists", h.ArtistsMainPage)
		{
			artists.POST("/", h.CreateArtist)
			artists.GET("/:id", h.ArtistPage)
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
