package http

import (
	"api/internal/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

func (s *Server) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.store.User().All(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

	render.JSON(w, r, users)
}

func (s *Server) getUserById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	user, err := s.store.User().ByID(r.Context(), id)
	if err != nil {
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

	render.JSON(w, r, user)
}

func (s *Server) updateUser(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	if err := s.store.User().Update(r.Context(), user); err != nil {
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}
}

func (s *Server) deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	if err := s.store.User().Delete(r.Context(), id); err != nil {
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}
}
