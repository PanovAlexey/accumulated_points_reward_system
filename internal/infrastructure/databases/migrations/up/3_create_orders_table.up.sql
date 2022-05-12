CREATE TABLE IF NOT EXISTS orders
(
    id        BIGSERIAL NOT NULL UNIQUE,
    user_id   BIGINT CONSTRAINT user_id_fk references users,
    password  VARCHAR(255) NOT NULL,
    number    INT NOT NULL,
    status    INT CONSTRAINT status_id_fk references order_status,
    PRIMARY KEY(id)
);