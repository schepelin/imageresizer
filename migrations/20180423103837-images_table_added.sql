
-- +migrate Up
CREATE TABLE images (
    id varchar PRIMARY KEY,
    raw bytea NOT NULL,
    created_at timestamp DEFAULT NOW()
);

-- +migrate Down
DROP TABLE images;
