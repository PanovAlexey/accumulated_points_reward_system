CREATE TABLE IF NOT EXISTS orders
(
    id          BIGSERIAL NOT NULL UNIQUE,
    user_id     BIGINT CONSTRAINT user_id_fk references users,
    number      BIGINT NOT NULL UNIQUE,
    status      VARCHAR NOT NULL,
    uploaded_at VARCHAR NOT NULL,
    PRIMARY KEY(id)
);