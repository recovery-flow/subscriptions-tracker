-- name: CreateTransaction :one
INSERT INTO subscription_transactions (
    user_id,
    payment_method_id,
    amount,
    currency,
    status,
    payment_provider,
    payment_id
) VALUES (
             $1, $2, $3, $4, $5, $6, $7
         )
    RETURNING *;

-- name: GetTransactionByID :one
SELECT * FROM subscription_transactions
WHERE id = $1;

-- name: ListTransactionsByUserID :many
SELECT * FROM subscription_transactions
WHERE user_id = $1
ORDER BY transaction_date DESC;

-- name: UpdateTransactionStatus :one
UPDATE subscription_transactions
SET
    status = $2
WHERE id = $1
    RETURNING *;

-- name: DeleteTransaction :exec
DELETE FROM subscription_transactions
WHERE id = $1;