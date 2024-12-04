package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "library_song/docs"
	"library_song/internal/controller"

	"log/slog"
)

// @title Music Library API
// @version 1.0
// @description API for managing an online music library
func NewRouter(engine *gin.Engine, log *slog.Logger, ctr *controller.Controller) {

	engine.Use(CORSMiddleware())

	// Serve Swagger UI
	engine.GET("/swagger/*eny", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Define API routes
	songs := engine.Group("/songs")

	// Register routes
	newSongRoutes(songs, ctr.Song, log)
}
