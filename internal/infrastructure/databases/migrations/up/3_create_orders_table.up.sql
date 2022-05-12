CREATE TABLE IF NOT EXISTS orders
(
    id            BIGSERIAL NOT NULL UNIQUE,
    user_id       BIGINT NOT NULL,
    password      VARCHAR(255) NOT NULL,
    status        int constraint status_id_fk references order_status,
    PRIMARY KEY(id)
);