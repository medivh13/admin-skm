package pertanyaan_dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type PertanyaanDTOInterface interface {
	Validate() error
}

type GetPertanyaanReqDTO struct {
	OpdID     int64 `json:"opd_id"`
	LayananID int64 `json:"layanan_id"`
}

func (dto *GetPertanyaanReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.OpdID, validation.Required, validation.Min(1)),
		validation.Field(&dto.LayananID, validation.Required, validation.Min(1)),
	); err != nil {
		return err
	}
	return nil
}

type PertanyaanRespDTO struct {
	ID   int64  `json:"id"`
	Soal string `json:"soal"`
}

// ------
type GetPertanyaanJawabanReqDTO struct {
	PertanyaanID int64 `json:"id"`
}
