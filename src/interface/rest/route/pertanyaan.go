package route

import (
	"net/http"

	pertanyaanHandler "admin-skm/src/interface/rest/handlers/pertanyaan"

	"github.com/go-chi/chi/v5"
)

func PertanyaanRouter(h pertanyaanHandler.PertanyaanHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.GetPertanyaanByOpdAndLayanan)
	return r
}
