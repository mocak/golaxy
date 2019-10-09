package storage

// Movie is the DTO for table movies
type Director struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Surname        string   `json:"surname"`
	Year      int      `json:"birth_year"`
	CreatedAt   string   `json:"created_at"`
}
