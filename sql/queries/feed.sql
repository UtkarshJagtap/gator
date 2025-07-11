-- name: CreateNewFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id) VALUES ( 
 $1,
 $2,
 $3,
 $4,
 $5,
 $6
)
Returning *;

-- name: GetFeeeds :many
SELECT feeds.name, feeds.url, users.name as user_name 
from feeds 
join users 
on feeds.user_id = users.id;

-- name: GetFeedIdByURL :one
SELECT feeds.id FROM feeds where feeds.url = $1;

