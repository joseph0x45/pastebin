package main

import (
	"embed"
	"flag"
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
	port := flag.String("port", "8080", "The port to start Pastebin on.")
	dbPath := flag.String("db", "pastebin.db", "Path to the SQLite database file used by PasteBin.")
	flag.Parse()
	mux := http.NewServeMux()

	conn := db.GetDBConnection(*dbPath)
	idGenerator, err := shortid.New(1, shortid.DefaultABC, 6969)
	if err != nil {
		panic(err)
	}

	handler := handlers.NewHandler(conn, idGenerator)

	mux.HandleFunc("GET /", handler.RenderHomePage)
	mux.HandleFunc("GET /pastes", handler.RenderPastesPage)
	mux.HandleFunc("POST /api/pastes", handler.CreateNewPaste)
	mux.HandleFunc("GET /api/pastes/{id}", handler.GetPasteByID)
	mux.HandleFunc("GET /api/pastes", handler.GetAllPastes)
	mux.HandleFunc("DELETE /api/pastes/{id}", handler.DeletePaste)
	mux.HandleFunc("GET /static/", http.FileServer(http.FS(static)).ServeHTTP)

	server := http.Server{
		Addr:    ":" + *port,
		Handler: mux,
	}
	log.Printf("Pastebin launched!\nVisit http://0.0.0.0:%s\nLearn more at https://github.com/joseph0x45/pastebin?tab=readme-ov-file#pastebin", *port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
