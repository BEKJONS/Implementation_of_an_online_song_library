package http

import (
	"github.com/gin-gonic/gin"
	"library_song/internal/entity"
	"library_song/internal/usecase"
	"log/slog"
	"net/http"
	"strconv"
)

type songRoutes struct {
	ss  *usecase.SongService
	log *slog.Logger
}

func newSongRoutes(router *gin.RouterGroup, ss *usecase.SongService, log *slog.Logger) {
	song := songRoutes{ss, log}
	router.POST("/", song.createSong)
	router.GET("/", song.listSongs)
	router.GET("/:id", song.getSong)
	router.PUT("/:id", song.updateSong)
	router.DELETE("/:id", song.deleteSong)
	router.GET("/:id/paginate", song.paginateText)
}

// CreateSong godoc
// @Summary Create a new song
// @Description Create a new song with details like title, group, release date, etc.
// @Tags Song
// @Accept json
// @Produce json
// @Param CreateSong body entity.Song1 true "Create song"
// @Success 201 {object} entity.Song
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /songs [post]
func (s *songRoutes) createSong(c *gin.Context) {
	var req entity.Song
	if err := c.ShouldBindJSON(&req); err != nil {
		s.log.Error("Error in getting from body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service to create song
	song, err := s.ss.CreateSong(req)
	if err != nil {
		s.log.Error("Error in creating song", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, song)
}

// ListSongs godoc
// @Summary List all songs
// @Description Get all songs with optional filters
// @Tags Song
// @Accept json
// @Produce json
// @Param groups query string false "Group name filter"
// @Param title query string false "Song title filter"
// @Param release_date query string false "Release date filter"
// @Param text query string false "Text filter"
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success 200 {array} entity.Song
// @Failure 500 {object} entity.Error
// @Router /songs [get]
func (s *songRoutes) listSongs(c *gin.Context) {
	group := c.DefaultQuery("groups", "")
	title := c.DefaultQuery("title", "")
	releaseDate := c.DefaultQuery("release_date", "")
	text := c.DefaultQuery("text", "")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		s.log.Error("Error in converting page to int", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		s.log.Error("Error in converting limit to int", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	songs, err := s.ss.ListSongs(entity.SongFilter{Groups: group, Song: title, Page: pageInt, Limit: limitInt, ReleaseDate: releaseDate, Text: text})
	if err != nil {
		s.log.Error("Error in listing songs", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, songs)
}

// GetSong godoc
// @Summary Get song details by ID
// @Description Fetch song details by its ID
// @Tags Song
// @Accept json
// @Produce json
// @Param id path string true "Song ID"
// @Success 200 {object} entity.Song
// @Failure 500 {object} entity.Error
// @Router /songs/{id} [get]
func (s *songRoutes) getSong(c *gin.Context) {
	songID := c.Param("id")
	song, err := s.ss.GetSong(songID)
	if err != nil {
		s.log.Error("Error in getting song", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, song)
}

// UpdateSong godoc
// @Summary Update song data
// @Description Update details of an existing song
// @Tags Song
// @Accept json
// @Produce json
// @Param id path string true "Song ID"
// @Param UpdateSong body entity.UpdateSong1 true "Updated song data"
// @Success 200 {object} entity.Song
// @Failure 500 {object} entity.Error
// @Router /songs/{id} [put]
func (s *songRoutes) updateSong(c *gin.Context) {
	var req entity.UpdateSong
	id := c.Param("id")
	if err := c.ShouldBindJSON(&req); err != nil {
		s.log.Error("Error in getting from body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ID = id
	updatedSong, err := s.ss.UpdateSong(req)
	if err != nil {
		s.log.Error("Error in updating song", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedSong)
}

// DeleteSong godoc
// @Summary Delete a song
// @Description Delete a song by ID
// @Tags Song
// @Accept json
// @Produce json
// @Param id path string true "Song ID"
// @Success 200 {object} entity.Message
// @Failure 500 {object} entity.Error
// @Router /songs/{id} [delete]
func (s *songRoutes) deleteSong(c *gin.Context) {
	songID := c.Param("id")
	msg, err := s.ss.DeleteSong(songID)
	if err != nil {
		s.log.Error("Error in deleting song", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, msg)
}

// PaginateText godoc
// @Summary Paginate text
// @Description Paginate the lyrics text for a specific song by verse.
// @Tags Song
// @Accept json
// @Produce json
// @Param id path string true "Song ID"
// @Param limit query int true "Number of verses per page"
// @Param offset query int true "Offset for pagination"
// @Success 200 {string} string "Paginated lyrics text"
// @Failure 400 {object} entity.Error "Invalid request parameters"
// @Failure 404 {object} entity.Error "Song not found"
// @Failure 500 {object} entity.Error "Internal server error"
// @Router /songs/{id}/paginate [get]
func (s *songRoutes) paginateText(c *gin.Context) {
	songID := c.Param("id")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	song, err := s.ss.GetSong(songID)
	if err != nil {
		s.log.Error("Error fetching song", "songID", songID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch song"})
		return
	}

	text, err := s.ss.PaginateText(song.Text, limit, offset)
	if err != nil {
		s.log.Error("Error paginating song text", "songID", songID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to paginate song text"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"text": text})
}
