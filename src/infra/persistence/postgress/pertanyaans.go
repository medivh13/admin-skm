package postgres

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : admin-skm
 */

import (
	"context"
	"fmt"

	dto "admin-skm/src/app/dtos/pertanyaan"
	repositories "admin-skm/src/domain/repositories"
	models "admin-skm/src/infra/models"

	"gorm.io/gorm"
)

type pertanyaanRepository struct {
	connection *gorm.DB
}

func NewPertanyaanRepository(db *gorm.DB) repositories.PertanyaansRepository {
	return &pertanyaanRepository{
		connection: db,
	}
}

func (repo *pertanyaanRepository) GetPertanyaanByLayananAndOPD(ctx context.Context, dto *dto.GetPertanyaanReqDTO) ([]*models.Pertanyaan, error) {
	var data []*models.Pertanyaan
	q := repo.connection.WithContext(ctx)

	if err := q.Raw(`SELECT master.pertanyaans.id, soal, master.pertanyaans.created_at, master.pertanyaans.updated_at, pilihan_satu
	from master.pertanyaans
	JOIN master.jawabans ON master.pertanyaans.id = master.jawabans.pertanyaan_id
	WHERE opd_id = ? AND layanan_id = ?`, dto.OpdID, dto.LayananID).Scan(&data).Error; err != nil {
		return nil, err
	}
	fmt.Println(data[0], "here")
	return data, nil
}

func (repo *pertanyaanRepository) GetPertanyaanJawabanByPertanyaanId(ctx context.Context, id string) (*models.Jawabans, error) {
	var pJModel models.Jawabans
	q := repo.connection.WithContext(ctx)

	if err := q.Raw(`
	SELECT j.id, j.pertanyaan_id, j.pilihan_satu, j.pilihan_dua, j.pilihan_tiga, j.pilihan_empat, p.soal
	from master.jawabans j
	join master.pertanyaans p ON p.id = j.pertanyaan_id
	WHERE j.pertanyaan_id = ? AND j. deleted_at IS NULL`, int64(1)).Scan(&pJModel).Error; err != nil {
		return nil, err
	}

	fmt.Printf("%+v here data", pJModel)
	fmt.Println(pJModel, "here")
	return &pJModel, nil
}
