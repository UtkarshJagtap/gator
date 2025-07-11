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

-- name: MarKFeedFetched :one
UPDATE feeds
SET
last_fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
select * from feeds
order by last_fetched_at asc nulls first
limit 1;

