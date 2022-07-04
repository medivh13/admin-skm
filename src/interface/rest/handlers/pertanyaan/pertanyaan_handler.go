package pertanyaan_handlers

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : admin-skm
 */

import (
	"log"
	"net/http"
	"strconv"

	dto "admin-skm/src/app/dtos/pertanyaan"

	usecases "admin-skm/src/app/use_cases/pertanyaan"

	common_error "admin-skm/src/infra/errors"
	"admin-skm/src/interface/rest/response"
	_ "net/http/pprof"
)

type PertanyaanHandlerInterface interface {
	GetPertanyaanByOpdAndLayanan(w http.ResponseWriter, r *http.Request)
}

type pertanyaanHandler struct {
	response         response.IResponseClient
	pertanyanUseCase usecases.PertanyaanUsecaseInterface
}

func NewPertanyaanHandler(r response.IResponseClient, u usecases.PertanyaanUsecaseInterface) PertanyaanHandlerInterface {
	return &pertanyaanHandler{
		response:         r,
		pertanyanUseCase: u,
	}
}

func (h *pertanyaanHandler) GetPertanyaanByOpdAndLayanan(w http.ResponseWriter, r *http.Request) {

	getDTO := dto.GetPertanyaanReqDTO{}

	dataOpdId, err := strconv.ParseInt(r.URL.Query().Get("opd_id"), 10, 64)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	getDTO.OpdID = dataOpdId

	dataLayananId, err := strconv.ParseInt(r.URL.Query().Get("opd_id"), 10, 64)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	getDTO.LayananID = dataLayananId

	err = getDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.pertanyanUseCase.GetPertanyaanByOpdAndLayanan(r.Context(), &getDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.UNKNOWN_ERROR, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Get Data",
		data,
		nil,
	)
}
