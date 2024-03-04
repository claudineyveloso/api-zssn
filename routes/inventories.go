package routes

import (
	"net/http"
	"time"

	"github.com/claudineyveloso/api-zssn/internal/db"
	"github.com/google/uuid"
)

func CreateInventory(w http.ResponseWriter, r *http.Request, queries *db.Queries, userID uuid.UUID) error {
	inventory := db.Inventory{
		ID:        uuid.New(),
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := queries.CreateInventory(r.Context(), db.CreateInventoryParams{
		ID:        inventory.ID,
		UserID:    inventory.UserID,
		CreatedAt: inventory.CreatedAt,
		UpdatedAt: inventory.UpdatedAt,
	}); err != nil {
		http.Error(w, "Erro ao criar inventário", http.StatusInternalServerError)
		return err
	}
	return nil
}
