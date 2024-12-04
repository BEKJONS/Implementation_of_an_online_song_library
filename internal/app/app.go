package app

import (
	"github.com/gin-gonic/gin"
	"library_song/config"
	"library_song/internal/controller"
	"library_song/internal/controller/http"

	"library_song/pkg/logger"
	"library_song/pkg/postgres"
	"log"
)

func Run(cfg *config.Config) {

	logger1 := logger.NewLogger()

	// Connect to PostgreSQL database
	db, err := postgres.Connection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the controller
	controller1 := controller.NewController(db, logger1)

	// Initialize rate limiting (using default settings or another mechanism if needed)
	// rateLimit := rate_limiting.NewRateLimiter(rdb, 5, time.Minute)  // Remove Redis-related rate limiting

	// Set up Gin engine
	engine := gin.Default()

	// Initialize routes
	http.NewRouter(engine, logger1, controller1)

	// Run the application on the specified port
	log.Fatal(engine.Run(cfg.RUN_PORT))
}
