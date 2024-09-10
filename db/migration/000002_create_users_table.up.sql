CREATE TABLE IF NOT EXISTS users(
  user_id uuid PRIMARY KEY,
  company_id uuid NOT NULL,
  name text NOT NULL,
  email text NOT NULL,
  password text NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (company_id) REFERENCES companies(company_id) ON DELETE CASCADE
);