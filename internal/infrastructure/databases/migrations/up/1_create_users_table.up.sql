CREATE TABLE IF NOT EXISTS users
(
    id            BIGSERIAL NOT NULL UNIQUE,
    login         VARCHAR(255) NOT NULL UNIQUE,
    password      VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
);
