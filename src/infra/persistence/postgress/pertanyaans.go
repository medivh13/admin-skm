package postgres

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : admin-skm
 */

import (
	"context"

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

// func (repo *usersRepository) Register(ctx context.Context, data *dto.UserReqDTO) (*models.Users, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(data.PassWord), 14)
// 	if err != nil {
// 		return nil, err
// 	}
// 	password := string(bytes)
// 	userModel := models.Users{
// 		Name:  data.Name,
// 		Email: data.Email,
// 	}

// 	q := repo.connection.WithContext(ctx)
// 	tx := q.Begin()
// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	if err := tx.Error; err != nil {
// 		return nil, err
// 	}

// 	if err := tx.Raw("INSERT INTO master.users (display_name, email_address, password) VALUES (?,?,?)", data.Name, data.Email, password).Scan(&userModel).Error; err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}

// 	return &userModel, tx.Commit().Error
// }

func (repo *pertanyaanRepository) GetPertanyaanByLayananAndOPD(ctx context.Context, dto *dto.GetPertanyaanReqDTO) ([]*models.Pertanyaan, error) {
	var data []*models.Pertanyaan
	q := repo.connection.WithContext(ctx)

	if err := q.Raw(`SELECT id, soal, created_at, updated_at 
	from master.pertanyaans 
	WHERE opd_id = ? AND layanan_id = ?`, dto.OpdID, dto.LayananID).Scan(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

// func (repo *usersRepository) UpdateToken(ctx context.Context, id int64, token string) error {
// 	q := repo.connection.WithContext(ctx)

// 	if err := q.Raw("UPDATE master.users SET token = ? WHERE id = ? RETURNING token", token, id).Scan(&token).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }
