package postgres

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : admin-skm
 */

import (
	"context"
	"fmt"

	repositories "admin-skm/src/domain/repositories"
	models "admin-skm/src/infra/models"

	"gorm.io/gorm"
)

type jawabanRepository struct {
	connection *gorm.DB
}

func NewJawabanRepository(db *gorm.DB) repositories.JawabansRepository {
	return &jawabanRepository{
		connection: db,
	}
}
func (repo *jawabanRepository) GetJawabanByPertanyaanId(ctx context.Context, id int64) ([]*models.Jawabans, error) {
	var pJModel []*models.Jawabans
	q := repo.connection.WithContext(ctx)

	if err := q.Raw(`
	SELECT j.id, j.pertanyaan_id, j.pilihan_satu, j.pilihan_dua, j.pilihan_tiga, j.pilihan_empat, p.soal
	from master.jawabans j
	join master.pertanyaans p ON p.id = j.pertanyaan_id
	WHERE p.layanan_id = ? AND j.deleted_at IS NULL`, int64(1)).Find(&pJModel).Error; err != nil {
		return nil, err
	}

	fmt.Printf("%+v here data", pJModel)
	fmt.Println(pJModel, "here")
	return pJModel, nil
}
