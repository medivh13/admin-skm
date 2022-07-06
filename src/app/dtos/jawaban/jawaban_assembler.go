package jawaban_dto

import (
	models "admin-skm/src/infra/models"
)

func ToGetPertanyaanJawaban(d *models.Jawabans) *PertanyaanJawabanRespDTO {
	return &PertanyaanJawabanRespDTO{
		PertanyaanID: d.PertanyaanID,
		Soal:         d.Soal,
		JawabanID:    d.ID,
		Pilihan1:     d.Pilihan1,
		Pilihan2:     d.Pilihan2,
		Pilihan3:     d.Pilihan3,
		Pilihan4:     d.Pilihan4,
	}
}

func ToPertanyaanJawabans(d []*models.Jawabans) []*PertanyaanJawabanRespDTO {
	var data []*PertanyaanJawabanRespDTO
	for _, val := range d {
		data = append(data, ToGetPertanyaanJawaban(val))
	}
	return data
}
