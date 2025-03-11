CREATE TABLE users (
    ID BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL, 
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE INDEX users_username_idx ON users (username);

CREATE TABLE songs (
    ID BIGSERIAL PRIMARY KEY,
    NameSong VARCHAR(255) NOT NULL,
    Artist VARCHAR(255) NOT NULL,
    Image VARCHAR(255) NOT NULL
);

CREATE INDEX name_songs_idx ON songs (NameSong);