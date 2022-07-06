package route

import (
	"net/http"

	jawabanHandler "admin-skm/src/interface/rest/handlers/jawaban"

	"github.com/go-chi/chi/v5"
)

func JawabanRouter(h jawabanHandler.JawabanHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.GetJawabanByPertanyaan)
	return r
}
