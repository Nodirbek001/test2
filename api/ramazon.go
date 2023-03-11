package api

import (
	"net/http"
	"test/models"
)

func (api *api) SaveUser(W http.ResponseWriter, r *http.Request) {
	var body models.User
	if err := BodyParser(r, &body); err != nil {
		HandleErrorResponse(W, 500, "body parse error")
	}
	err:=api.userservice
}
