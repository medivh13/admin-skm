package route

import (
	"net/http"

	layananHandler "admin-skm/src/interface/rest/handlers/layanan"

	"github.com/go-chi/chi/v5"
)

func LayananRouter(h layananHandler.LayananHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.GetLayananByOpdId)
	return r
}
