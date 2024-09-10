package persistent

import (
	"database/sql"
	"fmt"
	"upsider-coding-test/domain/auth"
	"upsider-coding-test/domain/company"
	"upsider-coding-test/domain/user"
)

type (
	userRepository struct {
		db   *sql.DB
		pSvc auth.PasswordService
	}
)

func (r *userRepository) FindByEmail(email user.Email) (*user.User, error) {
	query := `
		SELECT
			u.user_id,
			u.company_id,
			u.name,
			u.email,
			u.password
		FROM users u
		WHERE u.email = $1
	;`
	type userRow struct {
		UserID         string
		CompanyID      string
		Name           string
		Email          string
		HashedPassword string
	}
	var ur userRow
	err := r.db.QueryRow(query, email).Scan(&ur.UserID, &ur.CompanyID, &ur.Name, &ur.Email, &ur.HashedPassword)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to find user by email in FindByEmail: %w", err)
	}
	companyID, err := company.ParseCompanyID(ur.CompanyID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse company id in FindByEmail: %w", err)
	}
	hashed, err := r.pSvc.NewHashedIfValid(ur.HashedPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to create hashed password in FindByEmail: %w", err)
	}
	user, err := user.NewUser(ur.Name, ur.Email, hashed, companyID)
	if err != nil {
		return nil, fmt.Errorf("failed to create user in FindByEmail: %w", err)
	}
	return user, nil
}

func (r *userRepository) Save(user *user.User) error {
	query := `
		INSERT INTO users (user_id, company_id, name, email, password)
		VALUES ($1, $2, $3, $4, $5)
	;`
	_, err := r.db.Exec(query,
		user.ID().String(),
		user.CompanyID().String(),
		user.Username().String(),
		user.Email().String(),
		user.HashedPassword().String(),
	)
	if err != nil {
		return fmt.Errorf("failed to save user in Save: %w", err)
	}
	return nil
}
