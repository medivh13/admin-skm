package jawaban_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : admin-skm
 */

import (
	dto "admin-skm/src/app/dtos/jawaban"
	"admin-skm/src/domain/repositories"
	"context"
	"log"
)

type JawabanUsecaseInterface interface {
	GetPertanyaanJawabanByPertanyaanId(ctx context.Context, dtoReq *dto.GetPertanyaanJawabanReqDTO) ([]*dto.PertanyaanJawabanRespDTO, error)
}

type jawabanUseCase struct {
	JawabanRepo repositories.JawabansRepository
}

func NewJawabanUseCase(jawabanRepo repositories.JawabansRepository) *jawabanUseCase {
	return &jawabanUseCase{
		JawabanRepo: jawabanRepo,
	}
}

func (uc *jawabanUseCase) GetPertanyaanJawabanByPertanyaanId(ctx context.Context, dtoReq *dto.GetPertanyaanJawabanReqDTO) ([]*dto.PertanyaanJawabanRespDTO, error) {

	data, err := uc.JawabanRepo.GetJawabanByPertanyaanId(ctx, dtoReq.LayananID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return dto.ToPertanyaanJawabans(data), nil
}
