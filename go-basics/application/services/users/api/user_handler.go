package api

import (
	"net/http"

	"github.com/gnsalok/ps-go/application/services/users/types"
)

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	var user types.User
	_ = user
}
