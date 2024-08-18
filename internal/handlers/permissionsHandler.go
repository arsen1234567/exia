package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"tender/internal/models"
	"tender/internal/services"
)

type PermissionHandler struct {
	Service *services.PermissionService
}

// AddPermission adds a new permission for a user.
func (h *PermissionHandler) AddPermission(w http.ResponseWriter, r *http.Request) {
	var permission models.Permission
	err := json.NewDecoder(r.Body).Decode(&permission)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdPermission, err := h.Service.AddPermission(r.Context(), permission)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPermission)
}

// DeletePermission deletes a permission for a user.
func (h *PermissionHandler) DeletePermission(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get(":id")
	if idStr == "" {
		http.Error(w, "Missing permission ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid permission ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeletePermission(r.Context(), id)
	if err != nil {
		if errors.Is(err, models.ErrPermissionNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// UpdatePermission updates an existing permission.
func (h *PermissionHandler) UpdatePermission(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get(":id")
	if idStr == "" {
		http.Error(w, "Missing permission ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid permission ID", http.StatusBadRequest)
		return
	}

	var permission models.Permission
	err = json.NewDecoder(r.Body).Decode(&permission)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	permission.ID = id

	updatedPermission, err := h.Service.UpdatePermission(r.Context(), permission)
	if err != nil {
		if errors.Is(err, models.ErrPermissionNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedPermission)
}

// GetPermissionsByUserID getting permissions by user_id
func (h *PermissionHandler) GetPermissionsByUserID(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get(":user_id")
	if userIDStr == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	permissions, err := h.Service.GetPermissionsByUserID(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(permissions)
}
