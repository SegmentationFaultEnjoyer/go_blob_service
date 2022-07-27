package handlers

import (
	"net/http"
	"testService/internal/service/helpers"
	"testService/internal/service/requests"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func HandleBlobAdd(w http.ResponseWriter, r *http.Request) {
	blob, err := requests.ParseBlobRequest(w, r)

	if err != nil {
		Log(r).WithError(err).Info("ERROR GENERATE BLOB")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	blobData, err := helpers.FieldsToJSON(blob)

	if err != nil {
		Log(r).WithError(err).Info("ERROR JSON STRINGIFY")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	id, err := DB(r).Create(blob.Relationships.Author.Data.ID, blobData)

	if err != nil {
		Log(r).WithError(err).Info("ERROR ADDING BLOB")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	blob.ID = id

	Log(r).Info("BLOB CREATED")

	w.WriteHeader(http.StatusCreated)
	ape.Render(w, blob)
}
