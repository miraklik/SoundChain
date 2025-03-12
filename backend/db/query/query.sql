-- database: :memory:
CREATE TABLE users (
    ID BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE, 
    username VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL
);

CREATE INDEX users_username_idx ON users (username);

CREATE TABLE songs (
    ID BIGSERIAL PRIMARY KEY,
    NameSong VARCHAR(255) NOT NULL,
    Artist VARCHAR(255) NOT NULL,
    AlbumID BIGINT REFERENCES albums(ID) ON DELETE CASCADE,
    Image VARCHAR(255) NOT NULL
);

CREATE INDEX name_songs_idx ON songs (NameSong);

CREATE TABLE albums (
    ID BIGSERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Artist VARCHAR(255) NOT NULL,
    ReleaseDate DATE NOT NULL,
    Image VARCHAR(255) NOT NULL
);
CREATE INDEX name_albums_idx ON albums (Name);