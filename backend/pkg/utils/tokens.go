package utils

import (
	"crypto/rsa"
	"errors"
	"loan-back-services/conf"
	"loan-back-services/pkg/dto"
	"loan-back-services/pkg/logger"
	"loan-back-services/pkg/model"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type adminAccessTokenCustomClaims struct {
	Admin *model.Admin `json:"admin"`
	jwt.StandardClaims
}

type adminRefreshTokenCustomClaims struct {
	AdminId uint64 `json:"admin_id"`
	jwt.StandardClaims
}

func GenerateAccessToken(admin *model.Admin, key *rsa.PrivateKey) (string, error) {
	unixTime := time.Now().Unix()
	tokenExp := unixTime + 60*10 // 10 minutes

	claims := adminAccessTokenCustomClaims{
		Admin: admin,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  unixTime,
			ExpiresAt: tokenExp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		logger.Sugar.Error(err)
		return "", err
	}

	return ss, nil
}

func GenerateRefreshToken(key string) (*dto.RefreshTokenData, error) {
	currentTime := time.Now()
	tokenExp := currentTime.Add(time.Duration(60*60*24*1) * time.Second) // 1 day

	tokenId, err := uuid.NewRandom()
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	claim, err := ValidateAccessToken(key, conf.Rsa().PublicKey)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	claims := adminRefreshTokenCustomClaims{
		AdminId: claim.Admin.Id,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  currentTime.Unix(),
			ExpiresAt: tokenExp.Unix(),
			Id:        tokenId.String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString([]byte(key))
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	return &dto.RefreshTokenData{
		SS:        ss,
		Id:        tokenId,
		ExpiresIn: tokenExp.Sub(currentTime),
	}, nil
}

func ValidateAccessToken(tokenStr string, key *rsa.PublicKey) (*adminAccessTokenCustomClaims, error) {
	claims := &adminAccessTokenCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("access token is invalid")
	}

	claims, ok := token.Claims.(*adminAccessTokenCustomClaims)
	if !ok {
		return nil, errors.New("access token valid but couldn't parse claims")
	}

	return claims, nil
}

func ValidateRefreshToken(tokenStr, key string) (*adminRefreshTokenCustomClaims, error) {
	claims := &adminRefreshTokenCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("refresh token is invalid")
	}

	claims, ok := token.Claims.(*adminRefreshTokenCustomClaims)
	if !ok {
		return nil, errors.New("refresh token valid but couldn't parse claims")
	}

	return claims, nil
}
