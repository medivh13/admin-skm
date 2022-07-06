package models

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : admin-skm
 */

type Jawabans struct {
	ID           int64  `gorm:"id"`
	PertanyaanID int64  `gorm:"pertanyaan_id"`
	Soal         string `gorm:"soal"`
	Pilihan1     string `gorm:"pilihan_satu"`
	Pilihan2     string `gorm:"pilihan_dua"`
	Pilihan3     string `gorm:"pilihan_tiga"`
	Pilihan4     string `gorm:"pilihan_empat"`
}
