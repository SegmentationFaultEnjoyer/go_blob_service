CREATE TABLE IF NOT EXISTS blobs (
    id bigserial not null primary key,
    user_id bigserial not null references users(id),
    info json not null
);