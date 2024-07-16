package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/types"
)

type UserHandler struct {
	store types.UserStore
}

func NewUserHandler(store types.UserStore) *UserHandler {
	return &UserHandler{store: store}
}

func (u *UserHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /register", u.RegisterUser)
}

func (u *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	user := new(types.User)

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(
			map[string]string{
				"error": err.Error(),
			},
		)
		return
	}

	fmt.Println(user)

	error := u.store.CreateUser(*user)

	if error != nil {
		log.Fatal(error)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
