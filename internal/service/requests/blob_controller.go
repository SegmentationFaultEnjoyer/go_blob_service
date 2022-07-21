package requests

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func (c *Controller) HandleBlobAdd() http.HandlerFunc {
	type request struct {
		Type       string `json:"type"`
		Title      string `json:"title"`
		LinkFirst  string `json:"link_first"`
		LinkNext   string `json:"link_next"`
		LinkPrev   string `json:"link_prev"`
		LinkSelf   string `json:"link_self"`
		LinkLast   string `json:"link_last"`
		AuthorType string `json:"author_type"`
		AuthorID   string `json:"author_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			c.error(w, r, http.StatusBadRequest, err)
			return
		}

		blob, err := c.store.Blob().GenerateBlob(
			req.Type,
			req.Title,
			req.LinkFirst,
			req.LinkLast,
			req.LinkPrev,
			req.LinkSelf,
			req.LinkNext,
			req.AuthorType,
			req.AuthorID,
		)
		if err != nil {
			c.error(w, r, http.StatusInternalServerError, err)
			c.logger.Error("ERROR GENERATING BLOB")
			return
		}

		if err := c.store.Blob().Create(blob); err != nil {
			c.logger.Error("ERROR ADDING BLOB")
			c.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		c.respond(w, r, http.StatusCreated, blob)
	}
}

func (c *Controller) HandleBlobGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		blob, err := c.store.Blob().GetBlobByID(id)
		if err != nil {
			c.logger.Error("ERROR GETTING BLOB")
			c.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		c.respond(w, r, http.StatusFound, blob)
	}
}

func (c *Controller) HandleBlobDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if err := c.store.Blob().DeleteBlob(id); err != nil {
			c.logger.Error("ERROR DELETING BLOB")
			c.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		c.respond(w, r, http.StatusAccepted, nil)
	}
}

func (c *Controller) HandleBlobGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "user_id")

		blobs, err := c.store.Blob().GetAllBlobs(id)
		if err != nil {
			c.logger.Error("ERROR GETTING ALL BLOBS")
			c.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		c.respond(w, r, http.StatusAccepted, blobs)
	}
}
