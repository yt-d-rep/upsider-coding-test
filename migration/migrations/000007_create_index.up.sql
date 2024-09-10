CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_invoices_issued_at ON invoices(issued_at);
-- 実際に支払い処理などを行う際に支払い期限が重要になり検索キーとなることが想像つくため、payment_due_at にインデックスを追加しておく
CREATE INDEX idx_invoices_payment_due_at ON invoices(payment_due_at);
