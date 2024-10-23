-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT id, created_at, updated_at, name, url, user_id FROM feeds;

-- name: GetFeed :one
SELECT id, created_at, updated_at, name, url, user_id
FROM feeds
WHERE url = $1;

-- name: CreateFeedFollow :one
WITH inserted_feed_follows AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5) RETURNING * )
SELECT inserted_feed_follows.*, feeds.name as feed_name, users.name as user_name
FROM inserted_feed_follows
INNER JOIN users
on inserted_feed_follows.user_id = users.id
INNER JOIN feeds
on inserted_feed_follows.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT feeds.name as feed_name, users.name as user_name
FROM feed_follows
INNER JOIN feeds
ON feed_follows.feed_id = feeds.id
INNER JOIN users
ON feed_follows.user_id = users.id
WHERE feed_follows.user_id = $1;

-- name: Unfollow :exec
DELETE FROM feed_follows
USING feeds
WHERE feed_follows.feed_id = feeds.id
AND feeds.url = $2
AND feed_follows.user_id = $1;
