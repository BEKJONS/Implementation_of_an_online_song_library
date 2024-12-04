package main

import (
	"library_song/config"
	"library_song/internal/app"
)

func main() {
	cfg := config.NewConfig()

	app.Run(cfg)
}
