package storage

// Movie is the DTO for table movies
type Movie struct {
	ID        int      `json:"id"`
	Name      string   `json:"title"`
	Year      int      `json:"year_released"`
	Genre     []string `json:"genre"`
	Rating    float32  `json:"rating"`
	Director  Director   `json:"director"`
	Cast      []string `json:"cast"`
	CreatedAt string   `json:"created_at"`
}
