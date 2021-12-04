package http

import (
	"api/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (s *Server) createFeedback(w http.ResponseWriter, r *http.Request) {
	feedback := new(models.Feedback)
	if err := json.NewDecoder(r.Body).Decode(feedback); err != nil {
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	s.store.Feedback().Create(r.Context(), feedback)
	fmt.Println("Created", *feedback)
}

func (s *Server) getAllFeedbacks(w http.ResponseWriter, r *http.Request) {
	feedbacks, err := s.store.Feedback().All(r.Context())
	if err != nil {
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	render.JSON(w, r, feedbacks)
}

func (s *Server) getFeedbacksbyProfileId(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "userId")
	userId, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	feedback, err := s.store.Feedback().ByID(r.Context(), userId)
	if err != nil {
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	render.JSON(w, r, feedback)
}

func (s *Server) updateFeedbacks(w http.ResponseWriter, r *http.Request) {
	feedback := new(models.Feedback)
	if err := json.NewDecoder(r.Body).Decode(feedback); err != nil {
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	s.store.Feedback().Update(r.Context(), feedback)
}

func (s *Server) deleteFeedbackByid(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	s.store.Feedback().Delete(r.Context(), id)
}
