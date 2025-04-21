package docs

import (
	// "fmt"
	"net/http"
	"github.com/gorilla/mux"
	"search-app/database/models"
	"search-app/database"
	"encoding/json"
	"log"
	"io/ioutil"
)

// GetDocumentsForUserHandler will process the request to get all documents for a specific user
func DocHandler(w http.ResponseWriter, r *http.Request) {
	// Get the 'docID' parameter from the URL
	vars := mux.Vars(r) // Get the parameters from the URL
	docID := vars["id"] // Access the 'docID' from the map

	// Initialize a slice to store the documents for the user
	var documents []models.UserFile

	// Fetch all documents for the user from the database using GORM
	log.Println("docID:", docID)
	result := database.DB.Where("user_id = ?", docID).Find(&documents)
	if result.Error != nil {
		http.Error(w, "Error retrieving documents", http.StatusInternalServerError)
		return
	}

	// Set the response header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Marshal the documents slice to JSON and write the response
	err := json.NewEncoder(w).Encode(documents)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func DocByIdHandler(w http.ResponseWriter, r *http.Request) {
	// Get the 'docID' parameter from the URL
	vars := mux.Vars(r) // Get the parameters from the URL
	docID := vars["id"] // Access the 'docID' from the map

	// Initialize a slice to store the documents for the user
	var document models.UserFile

	// Fetch the document by ID from the database using GORM
	result := database.DB.First(&document, "id = ?", docID)

	if result.Error != nil {
		http.Error(w, "Document not found", http.StatusNotFound)
		return
	}

	content, err := ioutil.ReadFile(document.Filepath)
	if err != nil {
		http.Error(w, "Error reading file content", http.StatusInternalServerError)
		return
	}

	// Construct the response with document metadata and its content
	response := map[string]interface{}{
		"document_id": document.ID,
		"filepath": document.Filepath,
		"filename":    document.Filename,
		"uploaded_at": document.UploadedAt,
		"content":     string(content), // Convert the byte slice to string
	}

	// Set the content-type header to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Send the response as JSON
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
