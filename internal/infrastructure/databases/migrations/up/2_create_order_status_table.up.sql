CREATE TABLE IF NOT EXISTS order_status
(
    id            SERIAL NOT NULL UNIQUE,
    name          VARCHAR NOT NULL,
    created_at    TIMESTAMP NOT NULL,
    PRIMARY KEY(id)
);
