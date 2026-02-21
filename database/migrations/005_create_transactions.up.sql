CREATE TABLE transactions (
id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
user_id UUID REFERENCES users(id),
amount BIGINT NOT NULL,
type VARCHAR(20) NOT NULL,
status VARCHAR(20) DEFAULT 'pending',
created_at TIMESTAMP DEFAULT NOW()
);


