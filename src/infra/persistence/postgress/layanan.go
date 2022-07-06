package postgres

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : admin-skm
 */

import (
	"context"

	dto "admin-skm/src/app/dtos/layanan"
	repositories "admin-skm/src/domain/repositories"
	models "admin-skm/src/infra/models"

	"gorm.io/gorm"
)

type layananRepository struct {
	connection *gorm.DB
}

func NewLayananRepository(db *gorm.DB) repositories.LayananRepository {
	return &layananRepository{
		connection: db,
	}
}

func (repo *layananRepository) GetLayananByOpdId(ctx context.Context, dto *dto.GetLayananReqDTO) ([]*models.Layanan, error) {
	var data []*models.Layanan
	q := repo.connection.WithContext(ctx)

	if err := q.Raw(`SELECT id, display_name, unsur_pelayanan, created_at, updated_at
	from master.layanan_opds
	WHERE opd_id = ?`, dto.OpdID).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
