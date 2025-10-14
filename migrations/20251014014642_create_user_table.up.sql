CREATE extension IF NOT EXISTS "pgcrypto";

create table if not exists users (
  id UUID primary key default gen_random_uuid(),
  email text not null,
  password_hash text not null,
  created_at timestamp default now(),
  updated_at timestamp default now(),
);