package jawaban_dto

import (
	"log"

	validation "github.com/go-ozzo/ozzo-validation"
)

type PertanyaanJawabanDTOInterface interface {
	Validate() error
}

type GetPertanyaanJawabanReqDTO struct {
	LayananID int64 `json:"id"`
}

func (dto *GetPertanyaanJawabanReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.LayananID, validation.Required, validation.Min(1)),
	); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

type PertanyaanJawabanRespDTO struct {
	PertanyaanID int64  `json:"pertanyaan_id"`
	Soal         string `json:"soal"`
	JawabanID    int64  `json:"jawaban_id"`
	Pilihan1     string `json:"pilihan_1"`
	Pilihan2     string `json:"pilihan_2"`
	Pilihan3     string `json:"pilihan_3"`
	Pilihan4     string `json:"pilihan_4"`
}
