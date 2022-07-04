package repositories

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : admin-skm
 */

import (
	dtoLogin "admin-skm/src/app/dtos/login"
	dto "admin-skm/src/app/dtos/users"
	models "admin-skm/src/infra/models"
	"context"
)

type UsersRepository interface {
	Register(ctx context.Context, data *dto.UserReqDTO) (*models.Users, error)
	GetUser(ctx context.Context, data *dtoLogin.LoginReqDTO) (*models.Users, error)
	UpdateToken(ctx context.Context, id int64, token string) error
}
