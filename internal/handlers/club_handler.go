package handlers

import (
  "net/http"
  "strings"

  "epl-connect/internal/models"
  "epl-connect/internal/storage"
)

type ClubsHandler struct {
  Store *storage.MemoryStore
}

// POST /clubs
func (h *ClubsHandler) Create(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    models.WriteJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{Error: "method not allowed"})
    return
  }

  var req struct {
    OwnerUserID  string `json:"ownerUserId"`
    Name         string `json:"name"`
    Description  string `json:"description"`
  }
  if err := models.ReadJSON(r, &req); err != nil {
    models.WriteJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: "invalid json"})
    return
  }
  if strings.TrimSpace(req.Name) == "" {
    models.WriteJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: "club name required"})
    return
  }

  club := h.Store.CreateClub(models.Club{
    OwnerUserID: req.OwnerUserID,
    Name:        req.Name,
    Description: req.Description,
  })
  models.WriteJSON(w, http.StatusCreated, club)
}

// GET /clubs
func (h *ClubsHandler) List(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodGet {
    models.WriteJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{Error: "method not allowed"})
    return
  }
  models.WriteJSON(w, http.StatusOK, h.Store.ListClubs())
}