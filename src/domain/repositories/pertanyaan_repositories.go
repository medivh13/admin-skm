package repositories

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : admin-skm
 */

import (
	dto "admin-skm/src/app/dtos/pertanyaan"
	models "admin-skm/src/infra/models"
	"context"
)

type PertanyaansRepository interface {
	GetPertanyaanByLayananAndOPD(ctx context.Context, dto *dto.GetPertanyaanReqDTO) ([]*models.Pertanyaan, error)
	GetPertanyaanJawabanByPertanyaanId(ctx context.Context, id string) (*models.Jawabans, error)
}
