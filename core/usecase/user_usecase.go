package usecase

import (
	"ap_sell_products/common/errors"
	"ap_sell_products/common/log"
	"ap_sell_products/common/utils"
	"ap_sell_products/core/domain"
	"ap_sell_products/core/entities"
	"context"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	user domain.RepositoryUser
}

func NewUserUseCase(user domain.RepositoryUser) *UserUseCase {
	return &UserUseCase{
		user: user,
	}
}

func (u *UserUseCase) AddUser(ctx context.Context, req *entities.User) errors.Error {
	user, err := u.user.FindUserByUserName(ctx, req.UserName)
	if err != nil {
		return errors.NewCustomHttpError(http.StatusInternalServerError, errors.SYSTEM_ERROR_CODE, err.Error())
	}
	if user != nil {
		return errors.NewCustomHttpError(http.StatusConflict, 0, errors.EXIST_MESS)
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.NewCustomHttpError(http.StatusProcessing, errors.SYSTEM_ERROR_CODE, err.Error())
	}

	err = u.user.AddUser(ctx, &domain.User{
		ID:          utils.GenerateUniqueKey(),
		UserName:    req.UserName,
		DisplayName: req.DisplayName,
		Password:    string(hashPassword),
		Avatar:      req.Avatar,
		GoogleID:    "",
		CreateTime:  utils.GenTimeStemp(),
		UpdateTime:  utils.GenTimeStemp(),
	})
	if err != nil {
		log.Error("error", err)
		return errors.NewCustomHttpError(http.StatusInternalServerError, errors.SYSTEM_ERROR_CODE, err.Error())
	}
	return nil
}
