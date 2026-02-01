_package handlers

import (
	"net/http"
	"strings"

	"epl-connect/internal/models"
	"epl-connect/internal/storage"
)

type ChatHandler struct {
	Store *storage.MemoryStore
}

// POST /chatrooms
func (h *ChatHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		models.WriteJSON(w, http.StatusMethodNotAllowed,
			models.ErrorResponse{Error: "method not allowed"})
		return
	}

	var req struct {
		Name      string `json:"name"`
		CreatedBy string `json:"createdById"`
	}

	if err := models.ReadJSON(r, &req); err != nil {
		models.WriteJSON(w, http.StatusBadRequest,
			models.ErrorResponse{Error: "invalid json"})
		return
	}

	if strings.TrimSpace(req.Name) == "" {
		models.WriteJSON(w, http.StatusBadRequest,
			models.ErrorResponse{Error: "name required"})
		return
	}

	room := h.Store.CreateRoom(models.ChatRoom{
		Name:      req.Name,
		CreatedBy: req.CreatedBy,
	})

	models.WriteJSON(w, http.StatusCreated, room)
}

// POST /chatrooms/{id}/messages
func (h *ChatHandler) SendMessage(w http.ResponseWriter, r *http.Request, roomID string) {
	if r.Method != http.MethodPost {
		models.WriteJSON(w, http.StatusMethodNotAllowed,
			models.ErrorResponse{Error: "method not allowed"})
		return
	}

	var req struct {
		SenderID string `json:"senderId"`
		Text     string `json:"text"`
	}

	if err := models.ReadJSON(r, &req); err != nil {
		models.WriteJSON(w, http.StatusBadRequest,
			models.ErrorResponse{Error: "invalid json"})
		return
	}

	if strings.TrimSpace(req.Text) == "" {
		models.WriteJSON(w, http.StatusBadRequest,
			models.ErrorResponse{Error: "text required"})
		return
	}

	msg, err := h.Store.AddMessage(models.Message{
		RoomID:   roomID,
		SenderID: req.SenderID,
		Text:     req.Text,
	})

	if err != nil {
		models.WriteJSON(w, http.StatusNotFound,
			models.ErrorResponse{Error: err.Error()})
		return
	}

	models.WriteJSON(w, http.StatusCreated, msg)
}

// GET /chatrooms/{id}/messages
func (h *ChatHandler) ListMessages(w http.ResponseWriter, r *http.Request, roomID string) {
	if r.Method != http.MethodGet {
		models.WriteJSON(w, http.StatusMethodNotAllowed,
			models.ErrorResponse{Error: "method not allowed"})
		return
	}

	msgs, err := h.Store.ListMessages(roomID)
	if err != nil {
		models.WriteJSON(w, http.StatusNotFound,
			models.ErrorResponse{Error: err.Error()})
		return
	}

	models.WriteJSON(w, http.StatusOK, msgs)
}
