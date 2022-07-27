package handlers

import (
	"net/http"
	"testService/internal/service/helpers"
	"testService/resources"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func HandleBlobGet(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	blob := resources.Blob{}

	blobData, err := DB(r).Get(id)

	if err != nil {
		Log(r).WithError(err).Info("ERROR GETTING FROM DB")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if err := helpers.JSONToFields(&blob, blobData); err != nil {
		Log(r).WithError(err).Info("ERROR PARSE JSON")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	Log(r).Info("Blob found")

	w.WriteHeader(http.StatusFound)
	ape.Render(w, &resources.BlobResponse{
		Data:     blob,
		Included: resources.Included{},
	})
}
