package handlers

import (
  "net/http"
  "strings"

  "epl-connect/internal/models"
  "epl-connect/internal/storage"
)

type PostsHandler struct {
  Store *storage.MemoryStore
}

// POST /posts
func (h *PostsHandler) Create(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    models.WriteJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{Error: "method not allowed"})
    return
  }

  var req struct {
    AuthorID string `json:"authorId"`
    Title    string `json:"title"`
    Content  string `json:"content"`
  }
  if err := models.ReadJSON(r, &req); err != nil {
    models.WriteJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: "invalid json"})
    return
  }
  if strings.TrimSpace(req.Title) == "" {
    models.WriteJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: "title required"})
    return
  }

  post := h.Store.CreatePost(models.Post{
    AuthorID: req.AuthorID,
    Title:    req.Title,
    Content:  req.Content,
  })
  models.WriteJSON(w, http.StatusCreated, post)
}

// GET /posts
func (h *PostsHandler) List(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodGet {
    models.WriteJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{Error: "method not allowed"})
    return
  }
  models.WriteJSON(w, http.StatusOK, h.Store.ListPosts())
}

// PATCH /posts/id/approve
func (h *PostsHandler) Approve(w http.ResponseWriter, r *http.Request, id string) {
  if r.Method != http.MethodPatch {
    models.WriteJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{Error: "method not allowed"})
    return
  }
  post, err := h.Store.ApprovePost(id)
  if err != nil {
    models.WriteJSON(w, http.StatusNotFound, models.ErrorResponse{Error: err.Error()})
    return
  }
  models.WriteJSON(w, http.StatusOK, post)
}

// DELETE /posts/id
func (h *PostsHandler) Remove(w http.ResponseWriter, r *http.Request, id string) {
  if r.Method != http.MethodDelete {
    models.WriteJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{Error: "method not allowed"})
    return
  }
  post, err := h.Store.RemovePost(id)
  if err != nil {
    models.WriteJSON(w, http.StatusNotFound, models.ErrorResponse{Error: err.Error()})
    return
  }
  models.WriteJSON(w, http.StatusOK, post)
}