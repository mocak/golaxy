package storage

type Movie struct {
	Id        int
	Name      string
	Year      int
	Genre     []string
	Rating    float32
	Director  string
	Cast      []string
	CreatedAt string
}
