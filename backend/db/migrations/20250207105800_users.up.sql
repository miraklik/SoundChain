CREATE TABLE users (
    ID BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
);

CREATE INDEX users_username_idx ON users (username);
