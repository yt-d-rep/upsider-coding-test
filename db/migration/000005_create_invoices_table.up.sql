CREATE TABLE IF NOT EXISTS invoices(
  invoice_id uuid PRIMARY KEY,
  company_id uuid NOT NULL,
  partner_id uuid NOT NULL,
  issued_at timestamp NOT NULL,
  payment_amount decimal(10, 2) NOT NULL,
  fee decimal(10, 2) NOT NULL,
  fee_rate decimal(10, 2) NOT NULL,
  consumption_tax decimal(10, 2) NOT NULL,
  consumption_tax_rate decimal(10, 2) NOT NULL,
  invoice_amount decimal(10, 2) NOT NULL,
  payment_due_at timestamp NOT NULL,
  status smallint NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (company_id) REFERENCES companies(company_id) ON DELETE CASCADE,
  FOREIGN KEY (partner_id) REFERENCES partners(partner_id) ON DELETE CASCADE
);