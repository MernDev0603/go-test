package file

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"log"
	"sync"
	"search-app/database/models"
	"search-app/database"
)

var KnowledgeBase []string
var mu sync.Mutex

// UploadFileHandler handles the file upload and stores the file content for later searching
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read the uploaded file
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error reading file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	originalFileName := r.FormValue("file_name")

	// Create the uploads directory if it doesn't exist
	userID := r.FormValue("user_id")
	uploadDir := fmt.Sprintf("%s%s%s", "uploads/", userID, "/")
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.MkdirAll(uploadDir, os.ModePerm)
		if err != nil {
			http.Error(w, "Error creating upload directory", http.StatusInternalServerError)
			return
		}
	}

	// Generate a unique filename using the current timestamp and the original file extension
	currentTime := time.Now()
	formattedTime := currentTime.Format("20060102_150405")
	ext := filepath.Ext(originalFileName)
	fileName := fmt.Sprintf("%s%s", formattedTime, ext)

	// Create the full path to save the file
	filePath := filepath.Join(uploadDir, fileName)

	// Create or open the file on the server
	outFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error creating file on server", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	// Copy the uploaded file's content to the new file
	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, "Error saving file content", http.StatusInternalServerError)
		return
	}

	// Optionally: Process file content for knowledge base or other logic
	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file content", http.StatusInternalServerError)
		return
	}

	// Split the file content into lines and store it in KnowledgeBase
	lines := strings.Split(string(content), "\n")
	mu.Lock()
	KnowledgeBase = lines
	mu.Unlock()

	// Log the file details (for debugging purposes)
	log.Println("File saved successfully at:", filePath)
	log.Println("Generated filename:", fileName)
	log.Println("Original file name:", originalFileName)

	userFile := models.UserFile{
		UserID:   userID,
		Filename: originalFileName,
		Filepath: filePath,
	}

	err = database.DB.Create(&userFile).Error
	if err != nil {
		http.Error(w, "Error saving file data to the database", http.StatusInternalServerError)
		return
	}

	// Return the new document ID in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Respond with the newly created document ID
	response := fmt.Sprintf(`{"document_id": %d}`, userFile.ID)
	w.Write([]byte(response))
}
