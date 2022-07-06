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
	Jawaban1  string    `gorm:"pilihan_satu"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}
