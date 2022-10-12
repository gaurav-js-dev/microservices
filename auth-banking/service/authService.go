package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gaurav-js-dev/microservices/auth-banking/domain"
	"github.com/gaurav-js-dev/microservices/auth-banking/dto"
	"github.com/gaurav-js-dev/microservices/banking/errs"
	"github.com/gaurav-js-dev/microservices/banking/logger"
)

type AuthService interface {
	Verify(urlParams map[string]string) *errs.AppError
	Refresh(request dto.RefreshTokenRequest) (*dto.LoginResponse, *errs.AppError)
}

type DefaultAuthService struct {
	rolePermissions domain.RolePermissions
}

func jwtTokenFromString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(domain.HMAC_SAMPLE_SECRET), nil
	})
	if err != nil {
		logger.Error("Error while parsing token: " + err.Error())
		return nil, err
	}
	return token, nil
}

func NewLoginService() DefaultAuthService {
	return DefaultAuthService{}
}
