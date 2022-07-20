-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id bigserial not null primary key,
    email varchar not null unique,
    encrypted_password varchar not null
);

CREATE TABLE IF NOT EXISTS blobs (
    id bigserial not null primary key,
    user_id bigserial not null references users(id),
    type varchar not null,
    attributes json not null,
    relationships json not null
);
-- +migrate Down
DROP TABLE IF EXISTS  blobs;

DROP TABLE IF EXISTS  users;

