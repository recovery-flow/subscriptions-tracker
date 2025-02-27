-- name: CreateSubscriptionCancellationsTable :exec
CREATE TABLE IF NOT EXISTS subscription_cancellations (
                                                          id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES subscriptions(user_id) ON DELETE CASCADE,
    cancellation_date TIMESTAMP DEFAULT NOW(),
    reason TEXT
    );

CREATE INDEX IF NOT EXISTS idx_subscription_cancellations_date ON subscription_cancellations (cancellation_date);

-------------------------------------------------
-- CRUD и полезные запросы
-------------------------------------------------

-- name: CreateSubscriptionCancellation :one
INSERT INTO subscription_cancellations (
    user_id,
    reason
) VALUES (
             $1, $2
         )
    RETURNING *;

-- name: GetSubscriptionCancellationByID :one
SELECT * FROM subscription_cancellations
WHERE id = $1;

-- name: ListCancellationsByUserID :many
SELECT * FROM subscription_cancellations
WHERE user_id = $1
ORDER BY cancellation_date DESC;

-- name: DeleteSubscriptionCancellation :exec
DELETE FROM subscription_cancellations
WHERE id = $1;
