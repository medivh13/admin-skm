package models

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : admin-skm
 */

import (
	"time"
)

type Pertanyaan struct {
	ID        int64     `gorm:"id"`
	Soal      string    `gorm:"soal"`
	Password  string    `gorm:"password"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}
