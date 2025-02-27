CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE subscription_plans (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    billing_cycle VARCHAR(10) NOT NULL CHECK (billing_cycle IN ('monthly', 'yearly')),
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_subscription_plans_billing_cycle ON subscription_plans (billing_cycle);

-- Таблица методов оплаты
CREATE TABLE payment_methods (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid_v4(),
    user_id UUID NOT NULL UNIQUE,  -- Каждый пользователь может иметь только один способ оплаты по умолчанию
    type VARCHAR(50) NOT NULL CHECK (type IN ('credit_card', 'paypal', 'bank_transfer')),
    provider_token VARCHAR(255) NOT NULL,  -- Токенизированные данные карты
    is_default BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_payment_methods_user FOREIGN KEY (user_id) REFERENCES subscriptions(user_id) ON DELETE CASCADE
);

CREATE INDEX idx_payment_methods_type ON payment_methods (type);

-- Основная таблица подписок
CREATE TABLE subscriptions (
    user_id UUID PRIMARY KEY,  -- Один пользователь = одна подписка
    plan_id UUID NOT NULL REFERENCES subscription_plans(id),
    payment_method_id UUID REFERENCES payment_methods(id),
    status VARCHAR(20) NOT NULL CHECK (status IN ('active', 'canceled', 'expired', 'pending')),
    start_date TIMESTAMP NOT NULL DEFAULT NOW(),
    end_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_subscriptions_status ON subscriptions (status);
CREATE INDEX idx_subscriptions_end_date ON subscriptions (end_date);

-- Таблица истории транзакций
CREATE TABLE subscription_transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid_v4(),
    user_id UUID NOT NULL REFERENCES subscriptions(user_id) ON DELETE CASCADE,
    payment_method_id UUID REFERENCES payment_methods(id),
    amount DECIMAL(10,2) NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    status VARCHAR(20) NOT NULL CHECK (status IN ('success', 'failed', 'pending', 'canceled')),
    payment_provider VARCHAR(50) NOT NULL CHECK (payment_provider IN ('Stripe', 'PayPal')),
    payment_id VARCHAR(100) UNIQUE,  -- ID транзакции у провайдера
    transaction_date TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_subscription_transactions_status ON subscription_transactions (status);
CREATE INDEX idx_subscription_transactions_date ON subscription_transactions (transaction_date);

-- Расписание списаний
CREATE TABLE billing_schedules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid_v4(),
    user_id UUID NOT NULL REFERENCES subscriptions(user_id) ON DELETE CASCADE,
    scheduled_date TIMESTAMP NOT NULL,  -- Запланированная дата списания
    attempted_date TIMESTAMP,           -- Фактическая дата списания
    status VARCHAR(20) NOT NULL CHECK (status IN ('scheduled', 'processed', 'failed')),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_billing_schedules_scheduled_date ON billing_schedules (scheduled_date);
CREATE INDEX idx_billing_schedules_status ON billing_schedules (status);

-- Таблица отмен подписок
CREATE TABLE subscription_cancellations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid_v4(),
    user_id UUID NOT NULL REFERENCES subscriptions(user_id) ON DELETE CASCADE,
    cancellation_date TIMESTAMP DEFAULT NOW(),
    reason TEXT
);

CREATE INDEX idx_subscription_cancellations_date ON subscription_cancellations (cancellation_date);
