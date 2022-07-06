package models

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : admin-skm
 */

import (
	"time"
)

type Layanan struct {
	ID             int64     `gorm:"id"`
	DisplayName    string    `gorm:"display_name"`
	UnsurPelayanan string    `gorm:"unsur_pelayanan"`
	CreatedAt      time.Time `gorm:"created_at"`
	UpdatedAt      time.Time `gorm:"updated_at"`
}
