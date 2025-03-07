DROP INDEX IF EXISTS idx_subscription_cancellations_date;
DROP INDEX IF EXISTS idx_billing_schedule_status;
DROP INDEX IF EXISTS idx_billing_schedule_scheduled_date;
DROP INDEX IF EXISTS idx_subscription_transactions_date;
DROP INDEX IF EXISTS idx_subscription_transactions_status;
DROP INDEX IF EXISTS idx_subscriptions_end_date;
DROP INDEX IF EXISTS idx_subscriptions_status;
DROP INDEX IF EXISTS idx_payment_methods_type;
DROP INDEX IF EXISTS idx_subscription_plans_billing_interval_unit;

-- УДАЛЯЕМ ТАБЛИЦЫ (если существуют)
DROP TABLE IF EXISTS subscription_cancellations;
DROP TABLE IF EXISTS billing_schedule;
DROP TABLE IF EXISTS subscription_transactions;
DROP TABLE IF EXISTS subscriptions;
DROP TABLE IF EXISTS payment_methods;
DROP TABLE IF EXISTS subscription_plans;
DROP TABLE IF EXISTS subscription_types;

-- УДАЛЯЕМ И/ИЛИ СОЗДАЕМ EXTENSION
DROP EXTENSION IF EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";