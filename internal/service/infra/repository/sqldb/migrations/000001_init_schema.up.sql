CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE subscription_types (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE subscription_plans (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    type_id UUID NOT NULL REFERENCES subscription_types(id) ON DELETE CASCADE,
    price DECIMAL(10,2) NOT NULL,
    billing_interval INTEGER NOT NULL,
    billing_interval_unit VARCHAR(10) NOT NULL CHECK (billing_interval_unit IN ('once','day','week','month','year')),
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_subscription_plans_billing_interval_unit ON subscription_plans (billing_interval_unit);

CREATE TABLE payment_methods (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    type VARCHAR(50) NOT NULL CHECK (type IN ('google_pay', 'apple_pay', 'samsung_pay', 'paypal')),
    provider_token VARCHAR(255) NOT NULL,
    is_default BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_payment_methods_type ON payment_methods (type);

CREATE TABLE subscriptions (
    user_id UUID PRIMARY KEY,
    plan_id UUID NOT NULL REFERENCES subscription_plans(id),
    payment_method_id UUID NOT NULL REFERENCES payment_methods(id),
    status VARCHAR(20) NOT NULL CHECK (status IN ('active', 'inactive', 'expired')),
    start_date TIMESTAMP NOT NULL DEFAULT NOW(),
    end_date TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW() NOT NULL
);

CREATE INDEX idx_subscriptions_status ON subscriptions (status);
CREATE INDEX idx_subscriptions_end_date ON subscriptions (end_date);

CREATE TABLE subscription_transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    payment_method_id UUID REFERENCES payment_methods(id),
    amount DECIMAL(10,2) NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    status VARCHAR(20) NOT NULL CHECK (status IN ('success', 'failed')),
    payment_provider VARCHAR(50) NOT NULL CHECK (payment_provider IN ('Stripe', 'PayPal')),
    payment_id VARCHAR(100) UNIQUE,
    transaction_date TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_subscription_transactions_status ON subscription_transactions (status);
CREATE INDEX idx_subscription_transactions_date ON subscription_transactions (transaction_date);

CREATE TABLE billing_schedules (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    scheduled_date TIMESTAMP NOT NULL,
    attempted_date TIMESTAMP,
    status VARCHAR(20) NOT NULL CHECK (status IN ('planned', 'success', 'failed')),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_billing_schedules_scheduled_date ON billing_schedules (scheduled_date);
CREATE INDEX idx_billing_schedules_status ON billing_schedules (status);

ALTER TABLE subscription_plans
    ADD CONSTRAINT fk_subscription_plans_type FOREIGN KEY (type_id) REFERENCES subscription_types(id) ON DELETE CASCADE;

ALTER TABLE subscriptions
    ADD CONSTRAINT fk_subscriptions_plan FOREIGN KEY (plan_id) REFERENCES subscription_plans(id);

ALTER TABLE subscriptions
    ADD CONSTRAINT fk_subscriptions_payment_method FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id);

ALTER TABLE subscription_transactions
    ADD CONSTRAINT fk_subscription_transactions_payment_method FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id);

ALTER TABLE billing_schedules
    ADD CONSTRAINT fk_billing_schedules_subscription FOREIGN KEY (user_id) REFERENCES subscriptions(user_id) ON DELETE CASCADE;
