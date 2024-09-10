CREATE TABLE IF NOT EXISTS companies(
  company_id uuid PRIMARY KEY,
  name text NOT NULL,
  representative_name text NOT NULL,
  phone_number text NOT NULL,
  postal_code text NOT NULL,
  address text NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);