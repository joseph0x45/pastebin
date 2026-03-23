package handlers

import "github.com/go-chi/chi/v5"

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/", h.renderApp)
	r.Post("/pastes", h.createPaste)
	r.Post("/pastes/{id}/delete", h.deletePaste)
}
