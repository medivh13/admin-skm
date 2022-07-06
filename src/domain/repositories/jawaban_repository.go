package repositories

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : admin-skm
 */

import (
	models "admin-skm/src/infra/models"
	"context"
)

type JawabansRepository interface {
	GetJawabanByPertanyaanId(ctx context.Context, id int64) ([]*models.Jawabans, error)
}
