-- name: CreateSubscriptionPlan :one
INSERT INTO subscription_plans (
    name,
    description,
    price,
    billing_cycle,
    currency
) VALUES (
             $1, $2, $3, $4, $5
         )
    RETURNING *;

-- name: GetSubscriptionPlanByID :one
SELECT * FROM subscription_plans
WHERE id = $1;

-- name: ListSubscriptionPlans :many
SELECT * FROM subscription_plans
ORDER BY created_at DESC;

-- name: UpdateSubscriptionPlan :one
UPDATE subscription_plans
SET
    name = $2,
    description = $3,
    price = $4,
    billing_cycle = $5,
    currency = $6
WHERE id = $1
    RETURNING *;

-- name: DeleteSubscriptionPlan :exec
DELETE FROM subscription_plans
WHERE id = $1;