package main

import (
	"embed"
	"log"
	"net/http"
	"os"
)

//go:embed static
var static embed.FS

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, static, "static/index.html")
	})
	mux.Handle("GET /static/", http.FileServerFS(static))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
