package del

import (
	"errors"
	"net/http"

	resp "RestAPIURLShortener/internal/lib/api/response"
	"RestAPIURLShortener/internal/lib/logger/sl"
	"RestAPIURLShortener/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
)

type URLDelete interface {
	Delete(alias string) error
}

func New(log *slog.Logger, urlDelete URLDelete) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		const op = "handlers.url.delete.new"

		log := log.With(slog.String("op", op), slog.String("request_id", middleware.GetReqID(r.Context())))

		alias := chi.URLParam(r, "alias")

		if alias == "" {
			log.Info("alias is empty")
			render.JSON(w, r, resp.Error("invalid request"))
			return
		}

		err := urlDelete.Delete(alias)

		if errors.Is(err, storage.ErrURLNotFound) {
			log.Info("url not found", "alias", alias)
			render.JSON(w, r, resp.Error("not found"))
			return
		}
		if err != nil {
			log.Info("failed to get url", sl.Err(err))
			render.JSON(w, r, resp.Error("internal error"))
			return
		}

		log.Info("deleted url successfully", slog.String("alias", alias), slog.String("status", "ok"))

		render.JSON(w, r, resp.Ok())
	}
}
