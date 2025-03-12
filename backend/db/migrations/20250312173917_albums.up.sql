CREATE TABLE albums (
    ID BIGSERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Artist VARCHAR(255) NOT NULL,
    ReleaseDate DATE NOT NULL,
    Image VARCHAR(255) NOT NULL
);
CREATE INDEX name_albums_idx ON albums (Name);