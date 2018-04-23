
-- +migrate Up
CREATE TABLE images (
    hash varchar(32) PRIMARY KEY,
    data bytea,
    created_at timestamp DEFAULT NOW()
);

-- +migrate Down
DROP TABLE images;
