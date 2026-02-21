CREATE TABLE points (
id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
user_id UUID REFERENCES users(id),
balance INTEGER DEFAULT 0,
lifetime_earned INTEGER DEFAULT 0,
tier VARCHAR(20) DEFAULT 'bronze',
created_at TIMESTAMP DEFAULT NOW()
);


