package usecase

import "library_song/internal/entity"

// SongsRepo - интерфейс для работы с данными песен.
type SongsRepo interface {
	CreateSong(song entity.Song) (entity.Song, error)                    // Добавить новую песню
	GetSongByID(songID string) (entity.Song, error)                      // Получить песню по ID
	ListSongs(filter entity.SongFilter) ([]entity.Song, error)           // Список песен с фильтрацией
	UpdateSong(song entity.UpdateSong) (entity.Song, error)              // Обновить данные песни
	DeleteSong(songID string) error                                      // Удалить песню
	PaginateText(fullText string, limit int, offset int) (string, error) // Пагинация текста по куплетам
}

// ExternalAPI - интерфейс для взаимодействия с внешним API.
type ExternalAPI interface {
	FetchSongDetails(group, song string) (entity.SongDetails, error) // Получить данные песни из внешнего API
}

// SongTextHandler - интерфейс для обработки текста песни.
type SongTextHandler interface {
}
