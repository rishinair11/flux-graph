package serve

import (
	"embed"
	"log"
	"net/http"
	"path/filepath"
)

//go:embed static/*.html
var staticFS embed.FS

// ServeAssets starts a web server to serve the generated graph.html containing the SVG
func ServeAssets(filePath string, port string) {
	http.HandleFunc("/", handleGraphSVG)

	http.HandleFunc("/svg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	})

	log.Println("Serving your graph at http://localhost:" + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// handleGraphSVG serves graph.html, which contains the graph SVG, for "/" path.
func handleGraphSVG(w http.ResponseWriter, r *http.Request) {
	data, err := staticFS.ReadFile(filepath.Join("static", "graph.html"))
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(data) //nolint:errcheck
}
