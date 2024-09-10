package middleware

import "upsider-base/domain/auth"

type (
	tokenService struct{}
)

// TODO: トークン生成・検証処理を実装

const dummyToken = "TOKEN"

func (s *tokenService) Generate() (auth.Token, error) {
	return auth.NewToken(dummyToken), nil
}

func (s *tokenService) Validate(token auth.Token) (bool, error) {
	return token == dummyToken, nil
}
