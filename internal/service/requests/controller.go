package requests

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testService/internal/service/helpers"
	"testService/internal/store"

	"github.com/sirupsen/logrus"
)

type Controller struct {
	store  store.Store
	logger *logrus.Logger
}

func NewController(store store.Store) *Controller {
	logger, _ := helpers.ConfigureLogger("debug")

	return &Controller{
		store:  store,
		logger: logger,
	}
}

func (c *Controller) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	c.respond(w, r, code, map[string]string{"errors": err.Error(), "status": strconv.Itoa(code)})
}

func (c *Controller) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
