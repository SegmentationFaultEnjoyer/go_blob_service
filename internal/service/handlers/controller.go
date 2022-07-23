package handlers

import (
	"testService/internal/data"
	"testService/internal/data/postgres"
	"testService/internal/service/helpers"

	"github.com/sirupsen/logrus"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Controller struct {
	db     data.Blobs
	logger *logrus.Logger
}

func NewController(db *pgdb.DB) *Controller {
	logger, _ := helpers.ConfigureLogger("debug")

	return &Controller{
		db:     postgres.New(db),
		logger: logger,
	}
}
