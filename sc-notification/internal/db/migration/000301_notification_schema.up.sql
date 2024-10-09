SET TIMEZONE = 'UTC';

-- creating the notifications table
CREATE TABLE IF NOT EXISTS notifications (
   id UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
   user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
   content TEXT NOT NULL,
   created_at TIMESTAMP DEFAULT NOW()
);