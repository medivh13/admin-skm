package layanan_dto

import (
	models "admin-skm/src/infra/models"
)

func ToGetLayanan(d *models.Layanan) *LayananRespDTO {
	return &LayananRespDTO{
		ID:   d.ID,
		DisplayName: d.DisplayName,
		UnsurPelayanan: d.UnsurPelayanan,
	}
}

func ToLayanans(d []*models.Layanan) []*LayananRespDTO {
	var data []*LayananRespDTO
	for _, val := range d {
		data = append(data, ToGetLayanan(val))
	}
	return data
}
