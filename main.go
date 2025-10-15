package main

import (
	"embed"
	"log"
	"net/http"
	"pastebin/db"
	"pastebin/handlers"

	"github.com/teris-io/shortid"
)

//go:embed static/*
var static embed.FS

//go:generate tailwindcss -i static/input.css -o static/styles.css -m

func main() {
	mux := http.NewServeMux()

	conn := db.GetDBConnection()
	idGenerator, err := shortid.New(1, shortid.DefaultABC, 6969)
	if err != nil {
		panic(err)
	}

	handler := handlers.NewHandler(conn, idGenerator)

	mux.HandleFunc("GET /", handler.RenderHomePage)
	mux.HandleFunc("GET /pastes", handler.RenderPastesPage)
	mux.HandleFunc("POST /api/paste", handler.CreateNewPaste)
	mux.HandleFunc("DELETE /api/paste", handler.DeletePaste)
	mux.HandleFunc("GET /static/", http.FileServer(http.FS(static)).ServeHTTP)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("Starting Server on 8080. Visit http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
