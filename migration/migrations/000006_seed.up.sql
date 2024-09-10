DO $$
DECLARE
  company_id UUID := 'b8e7fce5-77a5-4e64-9e3c-90e0c5b4c17d';
  partner_id UUID := 'a5b6f8d4-9b44-4f9e-919f-d5cb2d7b8e9f';
  account_id UUID := 'c3d4e7f6-6b8e-4d4e-a55f-b1d5e2b6c8a1';
  BEGIN
    -- companies
    INSERT INTO companies (company_id, name, representative_name, phone_number, postal_code, address, created_at, updated_at)
    VALUES
      (company_id, '株式会社テスト', '鈴木一郎', '03-2345-6789', '150-0002', '東京都渋谷区2-2-2', NOW(), NOW());
    -- partners
    INSERT INTO partners (partner_id, company_id, name, representative_name, phone_number, postal_code, address, created_at, updated_at)
    VALUES
      (partner_id, company_id, '株式会社パートナー', '佐藤次郎', '03-4567-8901', '170-0004', '東京都杉並区4-4-4', NOW(), NOW());
    -- partner_bank_accounts
    INSERT INTO partner_bank_accounts (account_id, partner_id, bank_name, branch_name, account_number, account_name, created_at, updated_at)
    VALUES
      (account_id, partner_id, 'おかね銀行', '渋谷支店', '7654321', '株式会社パートナー', NOW(), NOW());
END $$;