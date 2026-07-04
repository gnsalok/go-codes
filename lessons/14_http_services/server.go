package httpservices

import (
	"encoding/json"
	"net/http"
	"strings"
)

type CreateGreetingRequest struct {
	Name string `json:"name"`
}

type GreetingResponse struct {
	Message string `json:"message"`
}

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("/greetings", createGreeting)
	return mux
}

func createGreeting(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var req CreateGreetingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
		return
	}
	name := strings.TrimSpace(req.Name)
	if name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "name is required"})
		return
	}
	writeJSON(w, http.StatusCreated, GreetingResponse{Message: "hello, " + name})
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
