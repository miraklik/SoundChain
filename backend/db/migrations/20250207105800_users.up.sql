CREATE TABLE users (
    ID BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE, 
    username VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL
);

CREATE INDEX users_username_idx ON users (username);
