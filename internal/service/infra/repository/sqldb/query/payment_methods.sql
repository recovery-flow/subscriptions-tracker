-- name: CreatePaymentMethod :one
INSERT INTO payment_methods (
    user_id,
    type,
    provider_token,
    is_default
) VALUES (
             $1, $2, $3, $4
         )
    RETURNING *;

-- name: GetPaymentMethodByUserID :one
SELECT * FROM payment_methods
WHERE user_id = $1;

-- name: UpdatePaymentMethod :one
UPDATE payment_methods
SET
    type           = $2,
    provider_token = $3,
    is_default     = $4
WHERE user_id = $1
    RETURNING *;

-- name: DeletePaymentMethod :exec
DELETE FROM payment_methods
WHERE user_id = $1;