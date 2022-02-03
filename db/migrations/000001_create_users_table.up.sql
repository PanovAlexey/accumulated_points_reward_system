CREATE TABLE IF NOT EXISTS users
(
    id            BIGSERIAL NOT NULL UNIQUE,
    username      VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
);
