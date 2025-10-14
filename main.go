package main

import (
	"context"
	"embed"
	"log"
	"net/http"
	"pastebin/components"
)

//go:embed static/*
var static embed.FS

//go:generate tailwindcss -i static/input.css -o static/styles.css -m

func main() {
	mux := http.NewServeMux()
	ctx := context.Background()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		components.Index().Render(ctx, w)
	})
	mux.Handle("/static/", http.FileServer(http.FS(static)))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("Starting Server on 8080. Visit http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
