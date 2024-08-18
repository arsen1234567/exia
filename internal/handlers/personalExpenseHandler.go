package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"tender/internal/models"
	"tender/internal/services"
)

// PersonalExpenseHandler handles HTTP requests for personal expenses.
type PersonalExpenseHandler struct {
	Service *services.PersonalExpenseService
}

// CreatePersonalExpense creates a new personal expense.
func (h *PersonalExpenseHandler) CreatePersonalExpense(w http.ResponseWriter, r *http.Request) {
	var expense models.PersonalExpense
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdPersonalExpanse, err := h.Service.CreatePersonalExpense(r.Context(), expense)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPersonalExpanse)
}

// GetPersonalExpenseByID retrieves a personal expense by ID.
func (h *PersonalExpenseHandler) GetPersonalExpenseByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get(":id")
	if idStr == "" {
		http.Error(w, "Missing expense ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid expense ID", http.StatusBadRequest)
		return
	}

	expense, err := h.Service.GetPersonalExpenseByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, models.ErrExpenseNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expense)
}

// GetAllPersonalExpenses retrieves all personal expenses.
func (h *PersonalExpenseHandler) GetAllPersonalExpenses(w http.ResponseWriter, r *http.Request) {
	expenses, err := h.Service.GetAllPersonalExpenses(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expenses)
}

// UpdatePersonalExpense updates an existing personal expense.
func (h *PersonalExpenseHandler) UpdatePersonalExpense(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get(":id")
	if idStr == "" {
		http.Error(w, "Missing expense ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid expense ID", http.StatusBadRequest)
		return
	}

	var expense models.PersonalExpense
	err = json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	expense.ID = id

	updatedExpense, err := h.Service.UpdatePersonalExpense(r.Context(), expense)
	if err != nil {
		if errors.Is(err, models.ErrExpenseNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedExpense)
}

// DeletePersonalExpense deletes a personal expense by ID.
func (h *PersonalExpenseHandler) DeletePersonalExpense(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get(":id")
	if idStr == "" {
		http.Error(w, "Missing expense ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid expense ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeletePersonalExpense(r.Context(), id)
	if err != nil {
		if errors.Is(err, models.ErrExpenseNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
