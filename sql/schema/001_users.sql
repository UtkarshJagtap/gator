-- +goose Up
CREATE TABLE users(
id UUID PRIMARY KEY,
created_at TIMESTAMP not null,
updated_at TIMESTAMP not null,
name text unique not null
);

CREATE TABLE feeds(
id UUID PRIMARY KEY,
created_at TIMESTAMP not null,
updated_at TIMESTAMP not null,
name text not null,
url text unique not null,
user_id uuid not null, 
FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE feed_follows(
id UUID PRIMARY KEY,
created_at TIMESTAMP not null,
updated_at TIMESTAMP not null,
user_id uuid not null,
feed_id uuid not null,
UNIQUE(user_id, feed_id),
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feed_follows;
DROP TABLE feeds;
DROP TABLE users;
