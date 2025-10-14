CREATE TABLE IF NOT EXISTS sessions (
  id_session UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  token TEXT NOT NULL,
  expired_at timestamp NOT NULL,
  created_at timestamp DEFAULT NOW(),
  updated_at timestamp DEFAULT NOW(),
  user_agent TEXT,
  ip_address TEXT,
  user_id UUID REFERENCES users(id) ON DELETE CASCADE
);