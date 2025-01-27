package lib

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type APIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func MakeHTTPFunc(fn APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		if err := fn(ctx, w, r); err != nil {
			WriteJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
