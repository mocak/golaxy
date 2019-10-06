DROP TABLE IF EXISTS movies;
CREATE TABLE movies
(
    id         serial PRIMARY KEY,
    name       varchar(200) NOT NULL,
    year       int       DEFAULT 0 CHECK (year > 0),
    genre      text[]       NULL,
    rating     numeric(4, 2),
    director   varchar(50),
    movie_cast text[],
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (name, year)
);