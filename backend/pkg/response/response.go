package response

import "net/http"

// Respond sends a JSON response with a success message
func Respond(w http.ResponseWriter, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"message": "%s", "data": %v}`, message, data)))
}
