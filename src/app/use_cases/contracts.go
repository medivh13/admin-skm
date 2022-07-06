package usecases

import (
	jawabanUc "admin-skm/src/app/use_cases/jawaban"
	layananUc "admin-skm/src/app/use_cases/layanan"
	loginUc "admin-skm/src/app/use_cases/login"
	pertanyaanUc "admin-skm/src/app/use_cases/pertanyaan"
	userUc "admin-skm/src/app/use_cases/user"
)

type AllUseCases struct {
	UserUseCase       userUc.UserUsecaseInterface
	LoginUseCase      loginUc.LoginUsecaseInterface
	PertanyaanUseCase pertanyaanUc.PertanyaanUsecaseInterface
	LayananUseCase    layananUc.LayananUsecaseInterface
	JawabanUseCase    jawabanUc.JawabanUsecaseInterface
}
