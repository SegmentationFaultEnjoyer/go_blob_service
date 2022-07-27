package handlers

import (
	"net/http"
	"testService/internal/data"
	"testService/internal/service/helpers"
	"testService/internal/service/requests"
	"testService/resources"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/kit/pgdb"
)

func HandleBlobGetAll(w http.ResponseWriter, r *http.Request) {
	request, err := requests.GetBlobListRequest(r)

	if err != nil {
		Log(r).WithError(err).Info("Error parsing request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	links := helpers.GetOffsetLinks(r, pgdb.OffsetPageParams{})
	blobs := make([]resources.Blob, 0)

	applyFilters(DB(r), request)
	result, err := DB(r).GetAll()

	if err != nil {
		Log(r).WithError(err).Info("Failed to get from data base")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	for _, blob := range result {
		newblob := resources.Blob{}

		if err := helpers.JSONToFields(&newblob, &blob); err != nil {
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

func applyFilters(q data.Blobs, request requests.BlobListRequest) {
	q.Page(request.OffsetPageParams)

	if request.FilterAuthor != nil {
		q.FilterByID(*request.FilterAuthor)
	}

}
