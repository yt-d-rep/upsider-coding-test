package auth

type (
	PasswordService interface {
		// NewHashedIfValid はすでにハッシュ化されている文字列を専用の型に変換して返します。
		NewHashedIfValid(password string) (HashedPassword, error)
		Hash(password string) (HashedPassword, error)
		Match(password RawPassword, hashed HashedPassword) bool
	}
	TokenService interface {
		Generate() (Token, error)
		Validate(token Token) (bool, error)
	}
)
