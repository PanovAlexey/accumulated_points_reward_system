CREATE TABLE IF NOT EXISTS payments
(
    id           BIGSERIAL NOT NULL UNIQUE,
    user_id      BIGINT CONSTRAINT user_id_payment_fk references users,
    order_id     BIGINT NOT NULL CONSTRAINT order_id_payment_fk references orders,
    processed_at VARCHAR NOT NULL,
    sum          DECIMAL(12, 2),
    PRIMARY KEY(id)
);
