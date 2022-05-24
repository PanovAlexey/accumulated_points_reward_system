INSERT INTO order_status (id, created_at, name)
VALUES (1, CURRENT_TIMESTAMP, 'NEW'),
       (2, CURRENT_TIMESTAMP, 'INVALID'),
       (3, CURRENT_TIMESTAMP, 'PROCESSING'),
       (4, CURRENT_TIMESTAMP, 'PROCESSED')
ON CONFLICT DO NOTHING;