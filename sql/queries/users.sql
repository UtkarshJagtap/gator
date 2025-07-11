-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users 
WHERE name = $1;

-- name: DeletUsers :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT name from users;

-- name: GetFeedFollowsForUser :many

select 
users.name as current_username,
feeds.name as feed_name

from users

inner join feed_follows 
on feed_follows.user_id = users.id

inner join feeds
on feed_follows.feed_id = feeds.id

where users.name = $1;
