package handlers

import (
  "net/http"
  "strings"

  "epl-connect/internal/models"
  "epl-connect/internal/storage"
)

type TeamsHandler struct {
  Store *storage.MemoryStore
}

// POST /teams
func (h *TeamsHandler) Create(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    models.WriteJSON(w, http.StatusMethodNotAllowed,
      models.ErrorResponse{Error: "method not allowed"})
    return
  }

  var req struct {
    Name    string `json:"name"`
    Coach   string `json:"coach"`
    Stadium string `json:"stadium"`
    Info    string `json:"info"`
  }

  if err := models.ReadJSON(r, &req); err != nil {
    models.WriteJSON(w, http.StatusBadRequest,
      models.ErrorResponse{Error: "invalid json"})
    return
  }

  if strings.TrimSpace(req.Name) == "" {
    models.WriteJSON(w, http.StatusBadRequest,
      models.ErrorResponse{Error: "team name required"})
    return
  }

  team := h.Store.CreateTeam(models.Team{
    Name:    req.Name,
    Coach:   req.Coach,
    Stadium: req.Stadium,
    Info:    req.Info,
  })

  models.WriteJSON(w, http.StatusCreated, team)
}

// GET /teams
func (h *TeamsHandler) List(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodGet {
    models.WriteJSON(w, http.StatusMethodNotAllowed,
      models.ErrorResponse{Error: "method not allowed"})
    return
  }

  models.WriteJSON(w, http.StatusOK, h.Store.ListTeams())
}