CREATE TABLE IF NOT EXISTS albums
(
    id     uuid  NOT NULL PRIMARY KEY,
    title  text  NOT NULL,
    artist text  NOT NULL,
    price  float NOT NULL
);
