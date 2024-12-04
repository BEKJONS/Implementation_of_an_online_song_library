package entity

// Song - основная сущность для работы с песнями.
type Song struct {
	ID          string `json:"id"`                             // Уникальный идентификатор
	Groups      string `json:"groups"`                         // Название группы
	Song        string `json:"song"`                           // Название песни
	ReleaseDate string `json:"release_date" db:"release_date"` // Дата выпуска
	Text        string `json:"text"`                           // Текст песни
	Link        string `json:"link"`                           // Ссылка на дополнительные данные
	CreatedAt   string `json:"created_at" db:"created_at"`     // Время создания записи
	UpdatedAt   string `json:"updated_at" db:"updated_at"`     // Время последнего обновления
}
type Song1 struct {
	Groups      string `json:"groups"`                         // Название группы
	Song        string `json:"song"`                           // Название песни
	ReleaseDate string `json:"release_date" db:"release_date"` // Дата выпуска
	Text        string `json:"text"`                           // Текст песни
	Link        string `json:"link"`                           // Ссылка на дополнительные данные
}

// SongDetails - детальная информация о песне из внешнего API.
type SongDetails struct {
	ReleaseDate string `json:"release_date" db:"release_date"` // Дата выпуска
	Text        string `json:"text"`                           // Текст песни
	Link        string `json:"link"`                           // Ссылка на данные
}

// SongFilter - фильтр для поиска и фильтрации песен.
type SongFilter struct {
	Groups      string `json:"groups"`                         // Фильтрация по группе
	Song        string `json:"song"`                           // Фильтрация по названию
	ReleaseDate string `json:"release_date" db:"release_date"` // Фильтрация по дате выпуска
	Text        string `json:"text"`
	Page        int    `json:"page"`  // Номер страницы для пагинации
	Limit       int    `json:"limit"` // Лимит записей на страницу
}

type Message struct {
	Message string `json:"message"`
}
type Error struct {
	Error string `json:"error"`
}
type UpdateSong struct {
	ID          string `json:"id"`
	Groups      string `json:"groups"`
	Song        string `json:"song"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
type UpdateSong1 struct {
	Groups      string `json:"groups"`
	Song        string `json:"song"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
