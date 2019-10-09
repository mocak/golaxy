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

DROP TABLE IF EXISTS directors;
CREATE TABLE directors
(
    id         serial PRIMARY KEY,
    name       varchar(200) NOT NULL,
    surname       varchar(200) NOT NULL   ,
    birth_year       int       DEFAULT  0 CHECK ( birth_year >0 ),
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (name)
);