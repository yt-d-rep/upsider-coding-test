CREATE TABLE IF NOT EXISTS partners(
  partner_id uuid PRIMARY KEY,
  company_id uuid NOT NULL,
  name text NOT NULL,
  representative_name text NOT NULL,
  phone_number text NOT NULL,
  postal_code text NOT NULL,
  address text NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (company_id) REFERENCES companies(company_id) ON DELETE CASCADE
);