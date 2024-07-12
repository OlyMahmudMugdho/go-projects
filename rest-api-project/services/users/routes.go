package users

import (
	"encoding/json"
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
	router.HandleFunc("/register", u.RegisterUser)
}

func (u *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	user := new(types.User)
	json.NewDecoder(r.Body).Decode(&user)

	error := u.store.CreateUser(*user)
	if error != nil {
		log.Fatal(error)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
