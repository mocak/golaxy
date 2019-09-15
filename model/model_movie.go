package model

type Movie struct {
	Id        int      `db:id`
	Name      string   `db:name`
	Year      int      `db:year`
	Genre     []string `db:genre`
	Rating    float32  `db:rating`
	Director  string   `db:director`
	Cast      []string `db:movie_cast`
	CreatedAt string   `db:created_at`
}
