package layanan_dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type LayananDTOInterface interface {
	Validate() error
}

type GetLayananReqDTO struct {
	OpdID int64 `json:"opd_id"`
}

func (dto *GetLayananReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.OpdID, validation.Required, validation.Min(1)),
	); err != nil {
		return err
	}
	return nil
}

type LayananRespDTO struct {
	ID             int64  `json:"id"`
	DisplayName    string `json:"display_name"`
	UnsurPelayanan string `json:"unsur_pelayanan"`
}
