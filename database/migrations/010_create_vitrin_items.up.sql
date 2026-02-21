CREATE TABLE vitrin_items (
id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
business_id UUID REFERENCES businesses(id),
media_url VARCHAR(255) NOT NULL,
size_type VARCHAR(20) NOT NULL,
views BIGINT DEFAULT 0,
likes BIGINT DEFAULT 0,
is_active BOOLEAN DEFAULT TRUE,
created_at TIMESTAMP DEFAULT NOW()
);


