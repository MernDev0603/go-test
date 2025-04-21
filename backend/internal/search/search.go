package search

import (
	// "fmt"
	"net/http"
	"github.com/schollz/closestmatch"
	"log"
	"io/ioutil"
	"strings"
	"os"
	"encoding/json"
)

var KnowledgeBase []string

// Perform the fuzzy search on the uploaded file content
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("search start:")

	var request struct {
		ID          string `json:"id"`
		Filepath    string `json:"filepath"`
		SearchQuery string `json:"searchquery"`
	}
	// Parse the JSON body
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := request.SearchQuery
	filepath := request.Filepath

	if query == "" {
		http.Error(w, "Query parameter missing", http.StatusBadRequest)
		return
	}

	file, err := os.Open(filepath)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file content", http.StatusInternalServerError)
		return
	}

	// Split the file content into lines and store it in KnowledgeBase
	lines := strings.Split(string(content), "\n")
	KnowledgeBase = lines

	// Here we use the file content stored in the global variable
	cm := closestmatch.New(KnowledgeBase, []int{2})
	results := cm.ClosestN(query, 5) // Get top 5 closest matches

	response := map[string]interface{}{
		"results": results,
	}

	// Set content-type header to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode the response to JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
