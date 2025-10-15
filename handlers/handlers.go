package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"pastebin/components"
	"pastebin/db"
	"pastebin/models"

	"github.com/jmoiron/sqlx"
	"github.com/teris-io/shortid"
)

type Handler struct {
	conn        *sqlx.DB
	idGenerator *shortid.Shortid
}

func NewHandler(conn *sqlx.DB, idGenerator *shortid.Shortid) *Handler {
	return &Handler{
		conn:        conn,
		idGenerator: idGenerator,
	}
}

func (h *Handler) RenderHomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	components.Index().Render(context.Background(), w)
}

func (h *Handler) RenderPastesPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	pastes, err := db.GetAllPastes(h.conn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	components.Pastes(pastes).Render(context.Background(), w)
}

func (h *Handler) CreateNewPaste(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[ERROR]: Error while creating paste: Failed to decode request body: ", err.Error())
		return
	}
	id, err := h.idGenerator.Generate()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[ERROR]: Error while creating paste: Failed to generate shortid: ", err.Error())
		return
	}
	newPaste := models.Paste{
		ID:      id,
		Title:   payload.Title,
		Content: payload.Content,
	}
	err = db.InsertPaste(h.conn, newPaste)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[ERROR]: Error while creating paste: ", err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) DeletePaste(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := db.DeletePaste(h.conn, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[ERROR]: Error while deleting paste: ", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}
