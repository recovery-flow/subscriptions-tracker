-- Удаляем индексы перед удалением таблиц
DROP INDEX IF EXISTS idx_subscription_cancellations_date;
DROP INDEX IF EXISTS idx_billing_schedules_status;
DROP INDEX IF EXISTS idx_billing_schedules_scheduled_date;
DROP INDEX IF EXISTS idx_subscription_transactions_date;
DROP INDEX IF EXISTS idx_subscription_transactions_status;
DROP INDEX IF EXISTS idx_subscriptions_end_date;
DROP INDEX IF EXISTS idx_subscriptions_status;
DROP INDEX IF EXISTS idx_payment_methods_type;
DROP INDEX IF EXISTS idx_subscription_plans_billing_cycle;

-- Удаляем таблицы в обратном порядке, чтобы избежать проблем с зависимостями
DROP TABLE IF EXISTS subscription_cancellations CASCADE;
DROP TABLE IF EXISTS billing_schedules CASCADE;
DROP TABLE IF EXISTS subscription_transactions CASCADE;
DROP TABLE IF EXISTS subscriptions CASCADE;
DROP TABLE IF EXISTS payment_methods CASCADE;
DROP TABLE IF EXISTS subscription_plans CASCADE;

-- Отключаем расширение uuid-ossp (если оно не используется в других местах)
DROP EXTENSION IF EXISTS "uuid-ossp";
