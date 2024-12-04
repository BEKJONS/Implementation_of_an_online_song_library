package controller

import (
	"github.com/jmoiron/sqlx"
	"library_song/internal/usecase"
	"library_song/internal/usecase/repo"
	"log/slog"
)

type Controller struct {
	Song *usecase.SongService
}

func NewController(db *sqlx.DB, log *slog.Logger) *Controller {
	songRepo := repo.NewSongsRepo(db)

	ctr := &Controller{
		Song: usecase.NewSongService(songRepo, log),
	}

	return ctr
}
