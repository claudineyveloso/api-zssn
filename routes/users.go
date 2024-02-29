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
	w.WriteHeader(http.StatusCreated)
	// Codifique os dados do usuário em formato JSON e envie-os como resposta
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
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
