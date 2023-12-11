package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nightlord189/example-hasher/internal/entity"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	Port    int
	Usecase IUsecase
}

func New(port int, uc IUsecase) *Handler {
	return &Handler{
		Port:    port,
		Usecase: uc,
	}
}

func (h *Handler) Run() error {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/api/hash", h.GetHashes)

	return http.ListenAndServe(fmt.Sprintf(":%d", h.Port), r)
}

func (h *Handler) GetHashes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req HashRequest
	if err := parseBodyJSON(r, &req); err != nil {
		log.Ctx(ctx).Error().Msgf("parse request error: %v", err.Error())
		responseString(ctx, w, http.StatusBadRequest, "parse request body error: "+err.Error())
		return
	}

	entities := make([]entity.HashRequestItem, 0, len(req.Items))

	for _, item := range req.Items {
		entityConverted, err := item.toEntity()
		if err != nil {
			responseString(ctx, w, http.StatusBadRequest, "convert item error: "+err.Error())
			return
		}
		entities = append(entities, entityConverted)
	}

	hashes, err := h.Usecase.GetHashes(ctx, entities)
	if err != nil {
		log.Ctx(ctx).Error().Msgf("parse request error: %v", err.Error())
		responseString(ctx, w, http.StatusBadRequest, "parse request body error: "+err.Error())
		return
	}

	convertedHashes := make([]HashResponseItem, 0, len(hashes))
	for _, hash := range hashes {
		converted := responseItemFromEntity(&hash)
		convertedHashes = append(convertedHashes, converted)
	}

	response := &HashResponse{
		Message: entity.SuccessResponseMessage,
		Items:   convertedHashes,
	}

	responseJSON(ctx, w, http.StatusOK, response)
}

func responseJSON(ctx context.Context, w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	marshaled, err := json.Marshal(data)
	if err != nil {
		log.Ctx(ctx).Error().Msgf("marshal response error: %v", err.Error())
	}
	if _, err = w.Write(marshaled); err != nil {
		log.Ctx(ctx).Error().Msgf("write response error: %v", err.Error())
	}
}

func parseBodyJSON(r *http.Request, out interface{}) error {
	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("read request body error: %w", err)
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Ctx(r.Context()).Error().Msgf("close request body error: %v", err.Error())
		}
	}()

	if err := json.Unmarshal(rawBody, out); err != nil {
		return fmt.Errorf("unmarshal json error: %w", err)
	}

	return nil
}

func responseString(ctx context.Context, w http.ResponseWriter, statusCode int, data string) {
	w.WriteHeader(statusCode)
	if _, err := w.Write([]byte(data)); err != nil {
		log.Ctx(ctx).Error().Msgf("write response error: %v", err.Error())
	}
}
