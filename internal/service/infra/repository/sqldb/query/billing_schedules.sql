-- name: CreateBillingSchedulesTable :exec
CREATE TABLE IF NOT EXISTS billing_schedules (
                                                 id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES subscriptions(user_id) ON DELETE CASCADE,
    scheduled_date TIMESTAMP NOT NULL,  -- Запланированная дата списания
    attempted_date TIMESTAMP,           -- Фактическая дата списания
    status VARCHAR(20) NOT NULL CHECK (status IN ('scheduled', 'processed', 'failed')),
    created_at TIMESTAMP DEFAULT NOW()
    );

CREATE INDEX IF NOT EXISTS idx_billing_schedules_scheduled_date ON billing_schedules (scheduled_date);
CREATE INDEX IF NOT EXISTS idx_billing_schedules_status ON billing_schedules (status);

-------------------------------------------------
-- CRUD и полезные запросы
-------------------------------------------------

-- name: CreateBillingSchedule :one
INSERT INTO billing_schedules (
    user_id,
    scheduled_date,
    attempted_date,
    status
) VALUES (
             $1, $2, $3, $4
         )
    RETURNING *;

-- name: GetBillingScheduleByID :one
SELECT * FROM billing_schedules
WHERE id = $1;

-- name: ListSchedulesByUserID :many
SELECT * FROM billing_schedules
WHERE user_id = $1
ORDER BY scheduled_date DESC;

-- name: UpdateBillingSchedule :one
UPDATE billing_schedules
SET
    attempted_date = $2,
    status         = $3
WHERE id = $1
    RETURNING *;

-- name: DeleteBillingSchedule :exec
DELETE FROM billing_schedules
WHERE id = $1;
