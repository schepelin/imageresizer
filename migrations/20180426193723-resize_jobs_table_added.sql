
-- +migrate Up
CREATE TABLE resize_jobs (
  id bigserial PRIMARY KEY,
  image_id varchar REFERENCES images ON DELETE CASCADE,
  width integer NOT NULL,
  height integer NOT NULL,
  status varchar NOT NULL,
  raw bytea,
  created_at timestamp WITHOUT TIME ZONE DEFAULT NOW(),
  UNIQUE (image_id, width, height)
);
-- +migrate Down
DROP TABLE resize_jobs;
