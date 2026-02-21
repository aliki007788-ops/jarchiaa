CREATE TABLE campaigns (
id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
title VARCHAR(200) NOT NULL,
budget BIGINT,
spent BIGINT DEFAULT 0,
clicks BIGINT DEFAULT 0,
views BIGINT DEFAULT 0,
start_date TIMESTAMP NOT NULL,
end_date TIMESTAMP NOT NULL,
status VARCHAR(20) DEFAULT 'draft',
created_at TIMESTAMP DEFAULT NOW()
);


