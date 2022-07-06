package repositories

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : admin-skm
 */

import (
	dto "admin-skm/src/app/dtos/layanan"
	models "admin-skm/src/infra/models"
	"context"
)

type LayananRepository interface {
	GetLayananByOpdId(ctx context.Context, dto *dto.GetLayananReqDTO) ([]*models.Layanan, error)
}
