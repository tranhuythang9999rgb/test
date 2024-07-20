package usecase

import (
	"ap_sell_products/common/configs"
	"ap_sell_products/common/errors"
	"ap_sell_products/common/log"
	"ap_sell_products/core/domain"
	"ap_sell_products/core/entities"
	"ap_sell_products/mcache"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis/v9"
)

type JwtUseCase struct {
	config *configs.Configs
	user   domain.RepositoryUser
}

func NewJwtUseCasee(config *configs.Configs, user domain.RepositoryUser) *JwtUseCase {
	return &JwtUseCase{
		config: config,
		user:   user,
	}
}

func createToken(creatorID int64, username string) (string, error) {
	expireDuration, err := time.ParseDuration(configs.Get().ExpireAccess)
	if err != nil {
		return "", err
	}

	expireTime := time.Now().Add(expireDuration).Unix()

	claims := &entities.UserJwtClaim{
		CreatorID: creatorID,
		UserName:  username,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expireTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(configs.Get().SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (u *JwtUseCase) VerifyToken(ctx context.Context, tokenString string) (*entities.UserJwtClaim, error) {

	token, err := jwt.ParseWithClaims(tokenString,
		&entities.UserJwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(configs.Get().SecretKey), nil
		})
	if err != nil {
		log.Error("error", err)
		return nil, jwt.ErrTokenInvalidIssuer
	}

	claims, ok := token.Claims.(*entities.UserJwtClaim)
	if !ok {
		return nil, jwt.ErrHashUnavailable
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, jwt.ErrTokenExpired
	}
	return claims, nil
}

func (u *JwtUseCase) Login(ctx context.Context, req *entities.LoginRequest) (*entities.LoginResponse, errors.Error) {
	user, err := u.user.FindUserByUserName(ctx, req.UserName)
	if err != nil {
		return nil, errors.NewCustomHttpError(http.StatusInternalServerError, errors.SYSTEM_ERROR_CODE, err.Error())
	}
	if user == nil {
		return nil, errors.NewCustomHttpError(http.StatusConflict, errors.NOT_EXIST_CODE, errors.NOT_EXIST_MESS)
	}
	//check password
	token, err := createToken(user.ID, req.UserName)
	if err != nil {
		return nil, errors.NewCustomHttpError(http.StatusProcessing, errors.SYSTEM_ERROR_CODE, err.Error())
	}

	osType := runtime.GOOS
	hostname, _ := os.Hostname()
	deviceInfo := fmt.Sprintf("%s , %s", osType, hostname)

	newSession := &entities.Session{
		Token:  token,
		Device: deviceInfo,
	}

	var sessions []*entities.Session
	sessionsJSON, err := mcache.GetRDB().HGet(ctx, req.UserName, "sessions").Result()
	if err != nil && err != redis.Nil {
		return nil, errors.NewCustomHttpError(http.StatusInternalServerError, errors.SYSTEM_ERROR_CODE, "Failed to get existing sessions")
	}

	if sessionsJSON != "" {
		err = json.Unmarshal([]byte(sessionsJSON), &sessions)
		if err != nil {
			return nil, errors.NewCustomHttpError(http.StatusInternalServerError, errors.SYSTEM_ERROR_CODE, "Failed to unmarshal sessions")
		}
	}

	sessions = append(sessions, newSession)

	updatedSessionJSON, err := json.Marshal(sessions)
	if err != nil {
		return nil, errors.NewCustomHttpError(http.StatusInternalServerError, errors.SYSTEM_ERROR_CODE, "Failed to marshal sessions")
	}

	err = mcache.GetRDB().HSet(ctx, req.UserName, "sessions", updatedSessionJSON).Err()
	if err != nil {
		return nil, errors.NewCustomHttpError(http.StatusInternalServerError, errors.SYSTEM_ERROR_CODE, fmt.Sprintf("Failed to store sessions: %v", err))
	}

	return &entities.LoginResponse{
		Token: token,
	}, nil
}

func (u *JwtUseCase) Logout(ctx context.Context, token string) errors.Error {

	infor, err := u.VerifyToken(ctx, token)
	if err != nil {
		return errors.NewCustomHttpError(http.StatusInternalServerError, errors.SYSTEM_ERROR_CODE, "Failed to get existing sessions")
	}

	var sessions []*entities.Session
	sessionsJSON, err := mcache.GetRDB().HGet(ctx, infor.UserName, "sessions").Result()
	if err != nil && err != redis.Nil {
		return errors.NewCustomHttpError(http.StatusInternalServerError, errors.SYSTEM_ERROR_CODE, "Failed to get existing sessions")
	}

	if sessionsJSON != "" {
		err = json.Unmarshal([]byte(sessionsJSON), &sessions)
		if err != nil {
			return errors.NewCustomHttpError(http.StatusInternalServerError, errors.SYSTEM_ERROR_CODE, "Failed to unmarshal sessions")
		}
	}

	var updatedSessions []*entities.Session
	for _, session := range sessions {
		if session.Token != token {
			updatedSessions = append(updatedSessions, session)
		}
	}

	updatedSessionsJSON, err := json.Marshal(updatedSessions)
	if err != nil {
		return errors.NewCustomHttpError(http.StatusInternalServerError, errors.SYSTEM_ERROR_CODE, "Failed to marshal sessions")
	}

	err = mcache.GetRDB().HSet(ctx, infor.UserName, "sessions", updatedSessionsJSON).Err()
	if err != nil {
		return errors.NewCustomHttpError(http.StatusInternalServerError, errors.SYSTEM_ERROR_CODE, fmt.Sprintf("Failed to update sessions in Redis: %v", err))
	}

	return nil
}

func (u *JwtUseCase) ListSession(ctx context.Context, token string) ([]*entities.Session, errors.Error) {

	var sessions []*entities.Session

	infor, err := u.VerifyToken(ctx, token)
	if err != nil {
		return nil, errors.NewCustomHttpError(http.StatusInternalServerError, errors.SYSTEM_ERROR_CODE, "Failed to get existing sessions")
	}

	sessionsJSON, err := mcache.GetRDB().HGet(ctx, infor.UserName, "sessions").Result()
	if err != nil && err != redis.Nil {
		return nil, errors.NewCustomHttpError(http.StatusInternalServerError, errors.SYSTEM_ERROR_CODE, "Failed to get existing sessions")
	}

	if sessionsJSON != "" {
		err = json.Unmarshal([]byte(sessionsJSON), &sessions)
		if err != nil {
			return nil, errors.NewCustomHttpError(http.StatusInternalServerError, errors.SYSTEM_ERROR_CODE, "Failed to unmarshal sessions")
		}
	}

	return sessions, nil
}
