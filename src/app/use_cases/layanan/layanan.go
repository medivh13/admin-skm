package layanan_usecase

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : admin-skm
 */

import (
	dto "admin-skm/src/app/dtos/layanan"
	"admin-skm/src/domain/repositories"
	"context"
	"log"
)

type LayananUsecaseInterface interface {
	GetLayananByOpdId(ctx context.Context, dtoReq *dto.GetLayananReqDTO) ([]*dto.LayananRespDTO, error)
}

type layananUseCase struct {
	LayananRepo repositories.LayananRepository
}

func NewLayananUseCase(layananRepo repositories.LayananRepository) *layananUseCase {
	return &layananUseCase{
		LayananRepo: layananRepo,
	}
}

func (uc *layananUseCase) GetLayananByOpdId(ctx context.Context, dtoReq *dto.GetLayananReqDTO) ([]*dto.LayananRespDTO, error) {

	data, err := uc.LayananRepo.GetLayananByOpdId(ctx, dtoReq)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return dto.ToLayanans(data), nil
}
