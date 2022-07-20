package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testService/internal/model"
)

func (c *Controller) HandleUserAdd() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			c.error(w, r, http.StatusBadRequest, err)
			return
		}

		user := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := c.store.User().Create(user); err != nil {
			//s.logger.Error("ERROR ADDING USER")
			fmt.Println("ERROR ADDING USER")
			c.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		c.respond(w, r, http.StatusCreated, user)
	}
}
