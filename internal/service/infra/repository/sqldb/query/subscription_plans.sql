-- name: CreateSubscriptionPlanVariant :one
INSERT INTO subscription_plans (
    type_id,
    price,
    billing_interval,
    billing_interval_unit,
    currency
) VALUES (
             $1, $2, $3, $4, $5
         )
    RETURNING *;

-- name: GetSubscriptionPlanVariantByID :one
SELECT * FROM subscription_plans
WHERE id = $1;

-- name: ListSubscriptionPlanVariantsByType :many
SELECT * FROM subscription_plans
WHERE type_id = $1
ORDER BY created_at DESC;

-- name: UpdateSubscriptionPlanVariant :one
UPDATE subscription_plans
SET
    price = $2,
    billing_interval = $3,
    billing_interval_unit = $4,
    currency = $5
WHERE id = $1
    RETURNING *;

-- name: DeleteSubscriptionPlanVariant :exec
DELETE FROM subscription_plans
WHERE id = $1;
