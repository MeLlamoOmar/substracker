CREATE TABLE IF NOT EXISTS subscriptions (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  service_name TEXT NOT null,
  price INTEGER NOT null,
  category TEXT,
  payment_method TEXT NOT null,
  billing_cycle TEXT NOT null,
  is_pinned BOOLEAN NOT null,
  notes TEXT,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now(),
  user_id UUID REFERENCES users(id) ON DELETE CASCADE
);