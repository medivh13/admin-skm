package users_dto

import (
	models "admin-skm/src/infra/models"
)

func ToReturnSaveUser(d *models.Users) *UserRespDTO {
	return &UserRespDTO{
		Name:  d.Name,
		Email: d.Email,
	}
}
