CREATE TABLE businesses (
id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
user_id UUID REFERENCES users(id),
name VARCHAR(200) NOT NULL,
verified BOOLEAN DEFAULT FALSE,
created_at TIMESTAMP DEFAULT NOW()
);


