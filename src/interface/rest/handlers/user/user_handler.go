package users_handlers

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : admin-skm
 */

import (
	"encoding/json"
	"fmt"
	"net/http"

	dtoLogin "admin-skm/src/app/dtos/login"
	dtoUser "admin-skm/src/app/dtos/users"
	loginUsecase "admin-skm/src/app/use_cases/login"
	usecases "admin-skm/src/app/use_cases/user"

	common_error "admin-skm/src/infra/errors"
	"admin-skm/src/interface/rest/response"
	_ "net/http/pprof"
)

type UserHandlerInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	response  response.IResponseClient
	usecase   usecases.UserUsecaseInterface
	lgUsecase loginUsecase.LoginUsecaseInterface
}

func NewUserHandler(r response.IResponseClient, u usecases.UserUsecaseInterface, l loginUsecase.LoginUsecaseInterface) UserHandlerInterface {
	return &userHandler{
		response:  r,
		usecase:   u,
		lgUsecase: l,
	}
}

func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) {

	postDTO := dtoUser.UserReqDTO{}
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	err = postDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.usecase.Register(r.Context(), &postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.UNKNOWN_ERROR, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Register",
		data,
		nil,
	)
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {

	postDTO := dtoLogin.LoginReqDTO{}
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	err = postDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.lgUsecase.Login(r.Context(), &postDTO)
	fmt.Printf("+%v", err)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Login",
		data,
		nil,
	)
}
