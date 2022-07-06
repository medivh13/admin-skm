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

	dto "admin-skm/src/app/dtos/jawaban"

	usecases "admin-skm/src/app/use_cases/jawaban"

	common_error "admin-skm/src/infra/errors"
	"admin-skm/src/interface/rest/response"
	_ "net/http/pprof"
)

type JawabanHandlerInterface interface {
	GetJawabanByPertanyaan(w http.ResponseWriter, r *http.Request)
}

type jawabanHandler struct {
	response       response.IResponseClient
	jawabanUsecase usecases.JawabanUsecaseInterface
}

func NewJawabanHandler(r response.IResponseClient, u usecases.JawabanUsecaseInterface) JawabanHandlerInterface {
	return &jawabanHandler{
		response:       r,
		jawabanUsecase: u,
	}
}

func (h *jawabanHandler) GetJawabanByPertanyaan(w http.ResponseWriter, r *http.Request) {

	getDTO := dto.GetPertanyaanJawabanReqDTO{}

	layanan_id, err := strconv.ParseInt(r.URL.Query().Get("layanan_id"), 10, 64)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	getDTO.LayananID = layanan_id

	err = getDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.jawabanUsecase.GetPertanyaanJawabanByPertanyaanId(r.Context(), &getDTO)
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
