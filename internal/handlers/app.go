package handlers

import (
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/joseph0x45/pastebin/internal/models"
	"github.com/teris-io/shortid"
)

func (h *Handler) renderApp(w http.ResponseWriter, r *http.Request) {
	pastes, err := h.conn.GetAllPastes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.render(w, "app", map[string]any{
		"Pastes": pastes,
	})
}

func (h *Handler) canDelete(r *http.Request) bool {
	var deleteToken string
	if h.version == "debug" {
		deleteToken = "DELETE_TOKEN"
	} else {
		deleteToken = os.Getenv("DELETE_TOKEN")
	}
	cookie, err := r.Cookie("delete_token")
	if err != nil {
		return false
	}
	return cookie.Value == deleteToken
}

func (h *Handler) deletePaste(w http.ResponseWriter, r *http.Request) {
	if !h.canDelete(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	id := chi.URLParam(r, "id")
	if err := h.conn.DeletePaste(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func generatePreview(content string, maxLen int) string {
	content = strings.Join(strings.Fields(content), " ")
	if len(content) <= maxLen {
		return content
	}
	return content[:maxLen] + "…"
}

func (h *Handler) createPaste(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	title := r.FormValue("title")
	content := r.FormValue("content")
	newPaste := &models.Paste{
		ID:      shortid.MustGenerate(),
		Title:   title,
		Preview: generatePreview(content, 80),
		Content: content,
	}
	if err := h.conn.InsertPaste(newPaste); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
