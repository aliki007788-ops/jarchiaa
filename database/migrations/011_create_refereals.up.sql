CREATE TABLE referrals (
id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
referrer_id UUID REFERENCES users(id),
referred_id UUID REFERENCES users(id),
code VARCHAR(20) NOT NULL,
status VARCHAR(20) DEFAULT 'pending',
reward_amount BIGINT DEFAULT 0,
created_at TIMESTAMP DEFAULT NOW()
);


