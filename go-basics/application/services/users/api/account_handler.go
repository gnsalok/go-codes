package api

import (
	"net/http"

	"github.com/gnsalok/ps-go/application/services/users/types"
)

func HandleGetAccount(w http.ResponseWriter, r *http.Request) {
	var user types.Account
	_ = user
}
