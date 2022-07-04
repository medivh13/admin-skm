package pertanyaan_dto

import (
	models "admin-skm/src/infra/models"
)

func ToGetPertanyaan(d *models.Pertanyaan) *PertanyaanRespDTO {
	return &PertanyaanRespDTO{
		ID:   d.ID,
		Soal: d.Soal,
	}
}

func ToPertanyaans(d []*models.Pertanyaan) []*PertanyaanRespDTO {
	var data []*PertanyaanRespDTO
	for _, val := range d {
		data = append(data, ToGetPertanyaan(val))
	}
	return data
}
