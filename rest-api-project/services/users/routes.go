package users

import "github.com/OlyMahmudMugdho/go-projects/rest-api-project/types"

type UserHandler struct {
	store types.UserStore
}

func NewUserHandler(store types.UserStore) *UserHandler {
	return &UserHandler{store: store}
}
