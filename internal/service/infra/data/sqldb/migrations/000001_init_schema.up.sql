CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE subscription_types (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    status VARCHAR(20) NOT NULL CHECK (status IN ('active', 'inactive')),
    updated_at TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE subscription_plans (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    type_id UUID NOT NULL REFERENCES subscription_types(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    billing_interval INTEGER NOT NULL,
    billing_cycle VARCHAR(10) NOT NULL CHECK (billing_cycle IN ('once','day','week','month','year')),
    status VARCHAR(20) NOT NULL CHECK (status IN ('active', 'inactive')),
    updated_at TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE payment_methods (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    type VARCHAR(50) NOT NULL CHECK (type IN ('google_pay', 'apple_pay', 'samsung_pay', 'paypal')),
    provider_token VARCHAR(255) NOT NULL,
    is_default BOOLEAN DEFAULT TRUE,
    updated_at TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE subscription_transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    payment_method_id UUID REFERENCES payment_methods(id),
    amount DECIMAL(10,2) NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    status VARCHAR(20) NOT NULL CHECK (status IN ('success', 'failed')),
    payment_provider VARCHAR(50) NOT NULL CHECK (payment_provider IN ('stripe', 'paypal')),
    payment_id VARCHAR(100) UNIQUE,
    transaction_date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE subscriptions (
   user_id UUID PRIMARY KEY,
   plan_id UUID NOT NULL REFERENCES subscription_plans (id),
   payment_method_id UUID NOT NULL DEFAULT uuid_nil() REFERENCES payment_methods(id) ON DELETE SET DEFAULT,
   status VARCHAR(20) NOT NULL CHECK (status IN ('active', 'inactive', 'canceled', 'expired')),
   availability VARCHAR(20) NOT NULL CHECK (availability IN ('available', 'unavailable', 'removed')),
   start_date TIMESTAMP NOT NULL DEFAULT NOW(),
   end_date TIMESTAMP NOT NULL,
   updated_at TIMESTAMP DEFAULT NOW(),
   created_at TIMESTAMP DEFAULT NOW() NOT NULL
);

CREATE TABLE billing_schedule (
    user_id UUID PRIMARY KEY NOT NULL REFERENCES subscriptions(user_id) ON DELETE CASCADE,
    scheduled_date TIMESTAMP NOT NULL,
    attempted_date TIMESTAMP,
    status VARCHAR(20) NOT NULL CHECK (status IN ('planned', 'processing', 'success', 'failed')),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION check_payment_methods_limit()
RETURNS trigger AS $$
BEGIN
    IF (SELECT COUNT(*) FROM payment_methods WHERE user_id = NEW.user_id) >= 5 THEN
        RAISE EXCEPTION 'User % already has 5 payment methods', NEW.user_id;
END IF;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER payment_methods_limit_trigger
    BEFORE INSERT ON payment_methods
    FOR EACH ROW
    EXECUTE FUNCTION check_payment_methods_limit();

CREATE UNIQUE INDEX uniq_user_default_payment_method
    ON payment_methods (user_id)
    WHERE is_default = true;
