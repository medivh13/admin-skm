package pertanyaan_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : admin-skm
 */

import (
	dto "admin-skm/src/app/dtos/pertanyaan"
	"admin-skm/src/domain/repositories"
	"context"
	"log"
)

type PertanyaanUsecaseInterface interface {
	GetPertanyaanByOpdAndLayanan(ctx context.Context, data *dto.GetPertanyaanReqDTO) ([]*dto.PertanyaanRespDTO, error)
}

type pertanyaanUseCase struct {
	PertanyaanRepo repositories.PertanyaansRepository
}

func NewPertanyaanUseCase(pertanyaanRepo repositories.PertanyaansRepository) *pertanyaanUseCase {
	return &pertanyaanUseCase{
		PertanyaanRepo: pertanyaanRepo,
	}
}

func (uc *pertanyaanUseCase) GetPertanyaanByOpdAndLayanan(ctx context.Context, dtoReq *dto.GetPertanyaanReqDTO) ([]*dto.PertanyaanRespDTO, error) {

	data, err := uc.PertanyaanRepo.GetPertanyaanByLayananAndOPD(ctx, dtoReq)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return dto.ToPertanyaans(data), nil
}
