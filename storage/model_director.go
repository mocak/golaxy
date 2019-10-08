package storage

// Movie is the DTO for table movies
type Director struct {
	ID          int      `json:"id"`
	Name        string   `json:"name_surname"`
	MovieCount  int      `json:"movie_count"`
	Genre       []string `json:"genre"`
	CreatedAt   string   `json:"created_at"`
}
