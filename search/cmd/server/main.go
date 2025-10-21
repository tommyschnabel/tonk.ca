package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/blevesearch/bleve"
)

const (
	indexFile = "./index.bleve"
	port      = 8080
)

type SearchRequest struct {
	Query string `json:"query"`
}

func main() {
	// Open the Bleve index
	index, err := bleve.Open(indexFile)
	if err != nil {
		log.Fatal(err)
	}
	defer index.Close()

	// Create a handler for search requests
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var searchReq SearchRequest
		if err := json.NewDecoder(r.Body).Decode(&searchReq); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Create and execute the search
		query := bleve.NewQueryStringQuery(searchReq.Query)
		searchRequest := bleve.NewSearchRequest(query)
		result, err := index.Search(searchRequest)
		if err != nil {
			http.Error(w, "Search failed", http.StatusInternalServerError)
			return
		}

		// Return the results as JSON
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	})

	// Start the server
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
