-- name: CreateSubscriptionType :one
INSERT INTO subscription_types (
    name,
    description
) VALUES (
             $1, $2
         )
    RETURNING *;

-- name: GetSubscriptionTypeByID :one
SELECT * FROM subscription_types
WHERE id = $1;

-- name: ListSubscriptionTypes :many
SELECT * FROM subscription_types
ORDER BY created_at DESC;

-- name: UpdateSubscriptionType :one
UPDATE subscription_types
SET
    name = $2,
    description = $3
WHERE id = $1
    RETURNING *;

-- name: DeleteSubscriptionType :exec
DELETE FROM subscription_types
WHERE id = $1;