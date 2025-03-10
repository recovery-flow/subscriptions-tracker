DROP TRIGGER IF EXISTS payment_methods_limit_trigger ON payment_methods;
DROP FUNCTION IF EXISTS check_payment_methods_limit();

DROP TABLE IF EXISTS billing_schedule;
DROP TABLE IF EXISTS subscriptions;
DROP TABLE IF EXISTS subscription_transactions;
DROP TABLE IF EXISTS payment_methods;
DROP TABLE IF EXISTS subscription_plans;
DROP TABLE IF EXISTS subscription_types;

DROP EXTENSION IF EXISTS "uuid-ossp";