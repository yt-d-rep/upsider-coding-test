package persistent

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"upsider-coding-test/domain/auth"
	"upsider-coding-test/domain/invoice"
	"upsider-coding-test/domain/user"

	"github.com/google/wire"
	_ "github.com/lib/pq"
)

var (
	db     *sql.DB
	dbOnce sync.Once

	uRepo     *userRepository
	uRepoOnce sync.Once

	ivcRepo     *invoiceRepository
	ivcRepoOnce sync.Once

	PersistentProviderSet wire.ProviderSet = wire.NewSet(
		ProvideDB,
		ProvideUserRepository,
		ProvideInvoiceRepository,
		wire.Bind(new(user.UserRepository), new(*userRepository)),
		wire.Bind(new(invoice.InvoiceRepository), new(*invoiceRepository)),
	)
)

func ProvideDB() *sql.DB {
	dbOnce.Do(func() {
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")
		dataSource := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
		var err error
		db, err = sql.Open("postgres", dataSource)
		if err != nil {
			panic(err)
		}
	})
	return db
}

func ProvideUserRepository(db *sql.DB, pSvc auth.PasswordService) *userRepository {
	uRepoOnce.Do(func() {
		uRepo = &userRepository{
			db:   db,
			pSvc: pSvc,
		}
	})
	return uRepo
}

func ProvideInvoiceRepository(db *sql.DB) *invoiceRepository {
	ivcRepoOnce.Do(func() {
		ivcRepo = &invoiceRepository{
			db: db,
		}
	})
	return ivcRepo
}
