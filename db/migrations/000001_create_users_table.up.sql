CREATE TABLE IF NOT EXISTS users
(
    id            BIGSERIAL NOT NULL UNIQUE,
    username      VARCHAR(255) NOT NULL UNIQUE,
    password      VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
);
