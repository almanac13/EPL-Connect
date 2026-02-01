package handlers

import (
  "net/http"
  "time"

  "epl-connect/internal/models"
  "epl-connect/internal/storage"
)

type MatchesHandler struct {
  Store *storage.MemoryStore
}

// POST /matches
func (h *MatchesHandler) Create(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    models.WriteJSON(w, http.StatusMethodNotAllowed,
      models.ErrorResponse{Error: "method not allowed"})
    return
  }

  var req struct {
    HomeTeamID string `json:"homeTeamId"`
    AwayTeamID string `json:"awayTeamId"`
    MatchDate  string `json:"matchDate"` // RFC3339
    Score      string `json:"score"`
  }

  if err := models.ReadJSON(r, &req); err != nil {
    models.WriteJSON(w, http.StatusBadRequest,
      models.ErrorResponse{Error: "invalid json"})
    return
  }

  date, err := time.Parse(time.RFC3339, req.MatchDate)
  if err != nil {
    models.WriteJSON(w, http.StatusBadRequest,
      models.ErrorResponse{Error: "invalid date format"})
    return
  }

  match := h.Store.CreateMatch(models.Match{
    HomeTeamID: req.HomeTeamID,
    AwayTeamID: req.AwayTeamID,
    MatchDate:  date,
    Score:      req.Score,
  })

  models.WriteJSON(w, http.StatusCreated, match)
}

// GET /matches
func (h *MatchesHandler) List(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodGet {
    models.WriteJSON(w, http.StatusMethodNotAllowed,
      models.ErrorResponse{Error: "method not allowed"})
    return
  }

  models.WriteJSON(w, http.StatusOK, h.Store.ListMatches())
}