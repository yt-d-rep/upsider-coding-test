CREATE TABLE IF NOT EXISTS partner_bank_accounts(
  account_id uuid PRIMARY KEY,
  partner_id uuid NOT NULL,
  bank_name text NOT NULL,
  branch_name text NOT NULL,
  account_number text NOT NULL,
  account_name text NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (partner_id) REFERENCES partners(partner_id) ON DELETE CASCADE,
  UNIQUE (partner_id)
);