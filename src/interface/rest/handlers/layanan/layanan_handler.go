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

	dto "admin-skm/src/app/dtos/layanan"

	usecases "admin-skm/src/app/use_cases/layanan"

	common_error "admin-skm/src/infra/errors"
	"admin-skm/src/interface/rest/response"
	_ "net/http/pprof"
)

type LayananHandlerInterface interface {
	GetLayananByOpdId(w http.ResponseWriter, r *http.Request)
}

type pertanyaanHandler struct {
	response       response.IResponseClient
	layananUseCase usecases.LayananUsecaseInterface
}

func NewLayananHandler(r response.IResponseClient, u usecases.LayananUsecaseInterface) LayananHandlerInterface {
	return &pertanyaanHandler{
		response:       r,
		layananUseCase: u,
	}
}

func (h *pertanyaanHandler) GetLayananByOpdId(w http.ResponseWriter, r *http.Request) {

	getDTO := dto.GetLayananReqDTO{}
	dataOpdId, err := strconv.ParseInt(r.URL.Query().Get("opd_id"), 10, 64)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	getDTO.OpdID = dataOpdId

	err = getDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.layananUseCase.GetLayananByOpdId(r.Context(), &getDTO)
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
