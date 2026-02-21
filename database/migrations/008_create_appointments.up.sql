CREATE TABLE appointments (
id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
business_id UUID REFERENCES businesses(id),
user_id UUID REFERENCES users(id),
appointment_date DATE NOT NULL,
appointment_time TIME NOT NULL,
status VARCHAR(20) DEFAULT 'pending',
created_at TIMESTAMP DEFAULT NOW()
);


