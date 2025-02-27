-- name: CreateSubscription :one
INSERT INTO subscriptions (
    user_id,
    plan_id,
    payment_method_id,
    status,
    start_date,
    end_date
) VALUES (
             $1, $2, $3, $4, $5, $6
         )
    RETURNING *;

-- name: GetSubscriptionByUserID :one
SELECT * FROM subscriptions
WHERE user_id = $1;

-- name: UpdateSubscription :one
UPDATE subscriptions
SET
    plan_id           = $2,
    payment_method_id = $3,
    status            = $4,
    start_date        = $5,
    end_date          = $6,
    updated_at        = NOW()
WHERE user_id = $1
    RETURNING *;

-- name: DeleteSubscription :exec
DELETE FROM subscriptions
WHERE user_id = $1;