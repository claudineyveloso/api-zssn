package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/claudineyveloso/api-zssn/internal/db"
	"github.com/google/uuid"
)

func CreateUser(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	var user db.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Erro ao decodificar corpo da solicitação", http.StatusBadRequest)
		return
	}

	if user.Name == "" ||
		user.Age == 0 ||
		user.Gender == "" ||
		user.Latitude == "" ||
		user.Longitude == "" {
		http.Error(w, "Campos obrigatórios, não podem ficar em branco", http.StatusBadRequest)
		return
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := queries.CreateUser(r.Context(), db.CreateUserParams{
		ID:                        user.ID,
		Name:                      user.Name,
		Age:                       user.Age,
		Gender:                    user.Gender,
		Latitude:                  user.Latitude,
		Longitude:                 user.Longitude,
		Infected:                  user.Infected,
		ContaminationNotification: user.ContaminationNotification,
		CreatedAt:                 user.CreatedAt,
		UpdatedAt:                 user.UpdatedAt,
	}); err != nil {
		http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
		return
	}

	if err := CreateInventory(w, r, queries, user.ID); err != nil {
		http.Error(w, "Erro ao criar inventário para o usuário", http.StatusInternalServerError)
		return
	}

	// Encode user data in JSON format and send it as a response
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	userID := r.URL.Query().Get("id")
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, "ID do usuário não pode ser vazio", http.StatusBadRequest)
		return
	}
	user, err := queries.GetUser(r.Context(), parsedUserID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter usuário: %v", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	users, err := queries.GetUsers(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter usuário: %v", err), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		http.Error(w, "Erro ao codificar em JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Erro ao escrever resposta", http.StatusInternalServerError)
		return
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	userID := r.URL.Query().Get("id")
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, "ID do usuário não pode ser vazio", http.StatusBadRequest)
		return
	}

	err = queries.DeleteUser(r.Context(), parsedUserID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao excluir o usuário: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Usuário excluído com sucesso!")

}

func UpdateLocation(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	var user db.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao ler dados do usuário: %v", err), http.StatusBadRequest)
		return
	}
	user.UpdatedAt = time.Now()
	err = queries.UpdateLocation(r.Context(), db.UpdateLocationParams{
		Latitude:  user.Latitude,
		Longitude: user.Longitude,
		UpdatedAt: user.UpdatedAt,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao atualizar usuário: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	var user db.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao ler dados do usuário: %v", err), http.StatusBadRequest)
		return
	}
	user.UpdatedAt = time.Now()
	err = queries.UpdateUser(r.Context(), db.UpdateUserParams{
		ID:        user.ID,
		Name:      user.Name,
		Age:       user.Age,
		Gender:    user.Gender,
		Latitude:  user.Latitude,
		Longitude: user.Longitude,
		UpdatedAt: user.UpdatedAt,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao atualizar usuário: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
