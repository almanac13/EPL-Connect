package handlers

import (
  "net/http"
  "strings"

  "epl-connect/internal/models"
  "epl-connect/internal/storage"
)

type AuthHandler struct {
  Store *storage.MemoryStore
}

// POST /users  (create user)
func (h *AuthHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    models.WriteJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{Error: "method not allowed"})
    return
  }

  var req struct {
    Name  string      `json:"name"`
    Email string      `json:"email"`
    Role  models.Role `json:"role"`
  }
  if err := models.ReadJSON(r, &req); err != nil {
    models.WriteJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: "invalid json"})
    return
  }
  if strings.TrimSpace(req.Email) == "" {
    models.WriteJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: "email required"})
    return
  }
  if req.Role == "" {
    req.Role = models.RoleFan
  }

  u := h.Store.CreateUser(models.User{
    Name:  req.Name,
    Email: req.Email,
    Role:  req.Role,
  })
  models.WriteJSON(w, http.StatusCreated, u)
}

// GET /users
func (h *AuthHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodGet {
    models.WriteJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{Error: "method not allowed"})
    return
  }
  models.WriteJSON(w, http.StatusOK, h.Store.ListUsers())
}