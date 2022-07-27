package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

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
