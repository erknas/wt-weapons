package api

import (
	"encoding/json"
	"net/http"

	"github.com/erknas/wt-weapons/lib"
	"github.com/erknas/wt-weapons/types"
	"github.com/go-chi/chi/v5"
)

func (s *Server) handleGetWeaponsByCategory(w http.ResponseWriter, r *http.Request) error {
	category := chi.URLParam(r, "category")

	weapons, err := s.storer.GetWeaponsByCategory(r.Context(), category)
	if err != nil {
		s.log.Error("failed to get weapons", "error", err, "path", r.URL.Path)
		return err
	}

	s.log.InfoContext(r.Context(), "got weapons", "category", category, "path", r.URL.Path)

	return lib.WriteJSON(w, http.StatusOK, weapons)
}

func (s *Server) handleAddWeapon(w http.ResponseWriter, r *http.Request) error {
	weapon := new(types.Weapon)

	if err := json.NewDecoder(r.Body).Decode(weapon); err != nil {
		s.log.Error("failed to decode request body", "error", err, "path", r.URL.Path)
		return err
	}

	id, err := s.storer.AddWeapon(r.Context(), weapon)
	if err != nil {
		s.log.Error("failed to add weapon", "error", err, "path", r.URL.Path)
		return err
	}

	return lib.WriteJSON(w, http.StatusOK, id)
}

func (s *Server) handleUpdateWeapon(w http.ResponseWriter, r *http.Request) error {
	var (
		name   = chi.URLParam(r, "name")
		weapon = new(types.Weapon)
	)

	if err := json.NewDecoder(r.Body).Decode(weapon); err != nil {
		s.log.Error("failed to decode request body", "error", err, "path", r.URL.Path)
		return err
	}

	if err := s.storer.UpdateWeapon(r.Context(), name, weapon); err != nil {
		s.log.Error("failed to update weapon", "error", err, "path", r.URL.Path)
		return err
	}

	return lib.WriteJSON(w, http.StatusOK, "ok")
}
