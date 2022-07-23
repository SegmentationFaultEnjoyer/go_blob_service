package handlers

import (
	"net/http"
	"testService/internal/service/requests"
	"testService/resources"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

const url string = "http://localhost:8081/integrations/testService/blob/"

func (c *Controller) HandleBlobAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		blob, err := requests.ParseBlobRequest(w, r)

		if err != nil {
			c.logger.Error("ERROR GENERATE BLOB")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		attributes, relationships, key, err := requests.FieldsToJSON(blob)

		if err != nil {
			c.logger.Error("ERROR JSON STRINGIFY")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		id, err := c.db.Create(blob.Relationships.Author.Data.ID, attributes, relationships, key)

		if err != nil {
			c.logger.Error("ERROR ADDING BLOB")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		blob.ID = id

		w.WriteHeader(http.StatusCreated)
		ape.Render(w, blob)
	}
}

func (c *Controller) HandleBlobGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		blob := resources.Blob{}

		key, attributes, relationships, err := c.db.Get(id)

		if err != nil {
			c.logger.Error("ERROR GETTING FROM DB")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		if err := requests.JSONToFields(&blob, attributes, relationships, key); err != nil {
			c.logger.Error("ERROR PARSE JSON")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		blob.Key.ID = id

		w.WriteHeader(http.StatusFound)
		ape.Render(w, &resources.BlobResponse{
			Data:     blob,
			Included: resources.Included{},
		})
	}
}

func (c *Controller) HandleBlobDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if err := c.db.Delete(id); err != nil {
			c.logger.Error("ERROR DELETING BLOB")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func (c *Controller) HandleBlobGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		blobs := make([]resources.Blob, 0)
		links := &resources.Links{}

		rows, err := c.db.GetAll(id)

		if err != nil {
			ape.RenderErr(w, problems.InternalError())
			return
		}

		defer rows.Close()

		counter := 0

		for rows.Next() {
			var key []byte
			var attributes []byte
			var relationships []byte
			var id string

			blob := resources.Blob{}

			if err := rows.Scan(&id, &key, &attributes, &relationships); err != nil {
				c.logger.Error("ERROR ROWS PROCESSING")
				ape.RenderErr(w, problems.InternalError())
				return
			}

			if err := requests.JSONToFields(&blob, attributes, relationships, key); err != nil {
				c.logger.Error("ERROR PARSING JSON")
				ape.RenderErr(w, problems.InternalError())
				return
			}

			if counter == 0 {
				links.Self = url + id
				links.First = links.Self
				links.Prev = links.Self
			}

			if counter == 1 {
				links.Next = url + id
			}

			blob.Key.ID = id
			blobs = append(blobs, blob)
			counter++
		}

		if len(blobs) < 1 {
			ape.RenderErr(w, problems.InternalError())
			return
		}

		links.Last = url + blobs[len(blobs)-1].Key.ID

		c.logger.Info("ALL BLOBS FOUND")

		w.WriteHeader(http.StatusFound)
		ape.Render(w, &resources.BlobListResponse{
			Data:     blobs,
			Included: resources.Included{},
			Links:    links,
		})
	}
}
