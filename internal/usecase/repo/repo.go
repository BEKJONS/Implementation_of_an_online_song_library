package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"library_song/internal/entity"
	"library_song/internal/usecase"
	"strings"
)

type SongsRepo struct {
	db *sqlx.DB
}

// NewSongsRepo создает новый экземпляр SongsRepo.
func NewSongsRepo(db *sqlx.DB) usecase.SongsRepo {
	return &SongsRepo{db: db}
}

// CreateSong добавляет новую песню в базу данных.
func (r *SongsRepo) CreateSong(song entity.Song) (entity.Song, error) {
	query := `
		INSERT INTO songs (groups, song, release_date, text, link)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, groups, song, release_date, text, link, created_at, updated_at
	`
	var newSong entity.Song
	err := r.db.QueryRowx(query, song.Groups, song.Song, song.ReleaseDate, song.Text, song.Link).
		Scan(&newSong.ID, &newSong.Groups, &newSong.Song, &newSong.ReleaseDate, &newSong.Text, &newSong.Link, &newSong.CreatedAt, &newSong.UpdatedAt)
	if err != nil {
		return entity.Song{}, fmt.Errorf("failed to create song: %w", err)
	}
	return newSong, nil
}

// GetSongByID получает песню по ID.
func (r *SongsRepo) GetSongByID(songID string) (entity.Song, error) {
	var song entity.Song
	query := `
		SELECT id, groups, song, release_date, text, link, created_at, updated_at
		FROM songs
		WHERE id = $1
	`
	err := r.db.Get(&song, query, songID)
	if err != nil {
		return entity.Song{}, fmt.Errorf("failed to get song: %w", err)
	}
	return song, nil
}

// ListSongs получает список песен с фильтрацией и пагинацией.
func (r *SongsRepo) ListSongs(filter entity.SongFilter) ([]entity.Song, error) {
	var songs []entity.Song
	var conditions []string
	var args []interface{}
	argIndex := 1

	// Add filters dynamically
	if filter.Groups != "" {
		conditions = append(conditions, fmt.Sprintf("groups ILIKE $%d", argIndex))
		args = append(args, "%"+filter.Groups+"%") // Partial match
		argIndex++
	}
	if filter.ReleaseDate != "" {
		conditions = append(conditions, fmt.Sprintf("release_date = $%d", argIndex))
		args = append(args, filter.ReleaseDate)
		argIndex++
	}
	if filter.Text != "" {
		conditions = append(conditions, fmt.Sprintf("text ILIKE $%d", argIndex)) // Case-insensitive
		args = append(args, "%"+filter.Text+"%")
		argIndex++
	}
	if filter.Song != "" {
		conditions = append(conditions, fmt.Sprintf("song ILIKE $%d", argIndex)) // Case-insensitive
		args = append(args, "%"+filter.Song+"%")
		argIndex++
	}

	// Base query
	query := `
		SELECT id, groups, song, release_date, text, link, created_at, updated_at
		FROM songs
	`

	// Add conditions if any
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	// Add ordering, limit, and offset
	query += fmt.Sprintf(" ORDER BY release_date DESC LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, filter.Limit, (filter.Page-1)*filter.Limit)

	// Execute query
	err := r.db.Select(&songs, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list songs: %w", err)
	}
	return songs, nil
}

// UpdateSong обновляет данные песни.
func (r *SongsRepo) UpdateSong(song entity.UpdateSong) (entity.Song, error) {
	query := `
		UPDATE songs
		SET groups = $1, song = $2, release_date = $3, text = $4, link = $5, updated_at = NOW()
		WHERE id = $6
		RETURNING id, groups, song, release_date, text, link, created_at, updated_at
	`
	var updatedSong entity.Song
	err := r.db.QueryRowx(query, song.Groups, song.Song, song.ReleaseDate, song.Text, song.Link, song.ID).
		Scan(&updatedSong.ID, &updatedSong.Groups, &updatedSong.Song, &updatedSong.ReleaseDate, &updatedSong.Text, &updatedSong.Link, &updatedSong.CreatedAt, &updatedSong.UpdatedAt)
	if err != nil {
		return entity.Song{}, fmt.Errorf("failed to update song: %w", err)
	}
	return updatedSong, nil
}

// DeleteSong удаляет песню по ID.
func (r *SongsRepo) DeleteSong(songID string) error {
	query := `
		DELETE FROM songs
		WHERE id = $1
	`
	_, err := r.db.Exec(query, songID)
	if err != nil {
		return fmt.Errorf("failed to delete song: %w", err)
	}
	return nil
}
func (r *SongsRepo) PaginateText(fullText string, limit int, offset int) (string, error) {
	sections := strings.Split(fullText, "\n\n")
	if offset < 0 || offset >= len(sections) {
		return "", fmt.Errorf("offset out of range")
	}

	end := offset + limit
	if end > len(sections) {
		end = len(sections)
	}

	return strings.Join(sections[offset:end], "\n\n"), nil
}
