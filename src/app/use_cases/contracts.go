package usecases

import (
	loginUc "admin-skm/src/app/use_cases/login"
	userUc "admin-skm/src/app/use_cases/user"
)

type AllUseCases struct {
	UserUseCase  userUc.UserUsecaseInterface
	LoginUseCase loginUc.LoginUsecaseInterface
}
