package usecase

import (
	"errors"
	"library_song/internal/entity"
	"log/slog"
)

type SongService struct {
	repo SongsRepo
	log  *slog.Logger
}

// NewSongService creates a new instance of SongService.
func NewSongService(repo SongsRepo, log *slog.Logger) *SongService {
	return &SongService{repo: repo, log: log}
}

func (s *SongService) CreateSong(in entity.Song) (entity.Song, error) {
	s.log.Info("started creating song", "title", in.Song)

	defer s.log.Info("ended creating song", "title", in.Song)

	// Validation checks
	if in.Song == "" {
		s.log.Error("error creating song", "error", errors.New("title is required"))
		return entity.Song{}, errors.New("title is required")
	}

	req := entity.Song{
		Song:        in.Song,
		Groups:      in.Groups,
		ReleaseDate: in.ReleaseDate,
		Text:        in.Text,
		Link:        in.Link,
	}

	song, err := s.repo.CreateSong(req)
	if err != nil {
		s.log.Error("error creating song", "error", err)
		return entity.Song{}, err
	}

	return song, nil
}

func (s *SongService) GetSong(songID string) (entity.Song, error) {
	s.log.Info("started getting song", "id", songID)
	defer s.log.Info("ended getting song", "id", songID)

	song, err := s.repo.GetSongByID(songID)
	if err != nil {
		s.log.Error("error getting song", "error", err)
		return entity.Song{}, err
	}
	return song, nil
}

func (s *SongService) ListSongs(filter entity.SongFilter) ([]entity.Song, error) {
	s.log.Info("started listing songs")
	defer s.log.Info("ended listing songs")

	songs, err := s.repo.ListSongs(filter)
	if err != nil {
		s.log.Error("error listing songs", "error", err)
		return nil, err
	}

	return songs, nil
}

func (s *SongService) UpdateSong(song entity.UpdateSong) (entity.Song, error) {
	s.log.Info("started updating song", "id", song.ID)
	defer s.log.Info("ended updating song", "id", song.ID)

	updatedSong, err := s.repo.UpdateSong(song)
	if err != nil {
		s.log.Error("error updating song", "error", err)
		return entity.Song{}, err
	}

	return updatedSong, nil
}

func (s *SongService) DeleteSong(songID string) (entity.Message, error) {
	s.log.Info("started deleting song", "id", songID)
	defer s.log.Info("ended deleting song", "id", songID)

	err := s.repo.DeleteSong(songID)
	if err != nil {
		s.log.Error("error deleting song", "error", err)
		return entity.Message{}, err
	}

	msg := entity.Message{
		Message: "Song deleted successfully",
	}

	return msg, nil
}

func (s *SongService) PaginateText(fullText string, limit int, offset int) (string, error) {
	s.log.Info("Paginating text", "limit", limit, "offset", offset)
	defer s.log.Info("Paginated text successfully ended")
	text, err := s.repo.PaginateText(fullText, limit, offset)
	if err != nil {
		s.log.Error("Error paginating text", "error", err)
		return "", err
	}

	return text, nil
}
