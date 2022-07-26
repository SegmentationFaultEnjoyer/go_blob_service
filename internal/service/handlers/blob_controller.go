package handlers

import (
	"net/http"
	"testService/internal/data"
	"testService/internal/service/helpers"
	"testService/internal/service/requests"
	"testService/resources"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/kit/pgdb"
)

func HandleBlobAdd(w http.ResponseWriter, r *http.Request) {
	blob, err := requests.ParseBlobRequest(w, r)

	if err != nil {
		Log(r).WithError(err).Info("ERROR GENERATE BLOB")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	blobData, err := requests.FieldsToJSON(blob)

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

func HandleBlobGet(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	blob := resources.Blob{}

	blobData, err := DB(r).Get(id)

	if err != nil {
		Log(r).WithError(err).Info("ERROR GETTING FROM DB")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if err := requests.JSONToFields(&blob, blobData); err != nil {
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

func HandleBlobDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := DB(r).Delete(id); err != nil {
		Log(r).WithError(err).Info("ERROR DELETING BLOB")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	Log(r).Info("BLOB DELETED")

	w.WriteHeader(http.StatusNoContent)
}

func HandleBlobGetAll(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("author_id")

	links := helpers.GetOffsetLinks(r, pgdb.OffsetPageParams{PageNumber: 1})
	blobs := make([]resources.Blob, 0)

	var result []data.Blob
	var err error

	if id == "" {
		result, err = DB(r).GetAll()
	} else {
		result, err = DB(r).FilterByID(id).GetAll()

	}

	if err != nil {
		Log(r).WithError(err).Info("Failed to get from data base")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	for _, blob := range result {
		newblob := resources.Blob{}

		if err := requests.JSONToFields(&newblob, &blob); err != nil {
			Log(r).WithError(err).Info("ERROR PARSING JSON")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		blobs = append(blobs, newblob)
	}

	Log(r).Info("ALL BLOBS FOUND")

	w.WriteHeader(http.StatusFound)
	ape.Render(w, &resources.BlobListResponse{
		Data:     blobs,
		Included: resources.Included{},
		Links:    links,
	})
}
