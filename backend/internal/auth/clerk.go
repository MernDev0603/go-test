package auth

import (
	"fmt"
	"net/http"
	"context"
	"github.com/clerk/clerk-sdk-go/v2"
    "github.com/clerk/clerk-sdk-go/v2/user"
)

// Get the user profile from Clerk API
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Get the Clerk user ID from JWT (you can extract this from token)
	// We assume the token contains the userID as a claim, this can be customized based on your Clerk setup
	userID := r.Header.Get("X-Clerk-User-Id")

	ctx := context.Background()

	// Set the API key with your Clerk Secret Key
	clerk.SetKey("sk_test_nOe5b9nEmwgyjV03JOVJ0JQB7lDkAZCraRmzHRWHNq")
	user, err := user.Get(ctx, userID)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching user profile: %v", err), http.StatusInternalServerError)
		return
	}

	// Return user profile as response
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"user": %v}`, user)))
}
