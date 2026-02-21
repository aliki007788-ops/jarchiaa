CREATE TABLE ads (
id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
business_id UUID REFERENCES businesses(id),
title VARCHAR(200) NOT NULL,
price BIGINT DEFAULT 0,
status VARCHAR(20) DEFAULT 'pending',
created_at TIMESTAMP DEFAULT NOW()
);


