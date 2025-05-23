package controller

import (
	"ddb-go-demo/internal/dto"
	"ddb-go-demo/internal/repository"
	"fmt"
	"net/http"
)

func (c *BaseController) SaveUser(w http.ResponseWriter, r *http.Request) {
	var args dto.User
	if err := c.BindJsonBody(r, &args); err != nil {
		c.JsonReponse(w, http.StatusInternalServerError, c.CreateMessage("Failed to parse request body"))
		return
	}

	err := repository.CreateUser(c.ctx.Db, &args)
	if err != nil {
		c.JsonReponse(w, http.StatusInternalServerError, c.CreateMessage("Failed to get user"))
		return
	}

	c.JsonReponse(w, http.StatusOK, c.CreateMessage("USer was created successfully"))
}

func (c *BaseController) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		c.JsonReponse(w, http.StatusBadRequest, c.CreateMessage("Email parameter is required"))
		return
	}

	user, err := repository.GetUserByEmail(c.ctx.Db, email)
	if err != nil {
		if err.Error() == fmt.Sprintf("no user found with email: %s", email) {
			c.JsonReponse(w, http.StatusNotFound, c.CreateMessage("User not found"))
		} else {
			c.JsonReponse(w, http.StatusInternalServerError, c.CreateMessage("Failed to get user : "+err.Error()))
		}
		return
	}

	c.JsonReponse(w, http.StatusOK, user)
}
