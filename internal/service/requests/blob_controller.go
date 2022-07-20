package requests

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (c *Controller) HandleBlobAdd() http.HandlerFunc {
	type request struct {
		Type        string `json:"type"`
		Title       string `json:"title"`
		LinkSelf    string `json:"link_self"`
		LinkRelated string `json:"link_related"`
		AuthorType  string `json:"author_type"`
		AuthorID    string `json:"author_id"`
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
			req.LinkSelf,
			req.LinkRelated,
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
		string_id := chi.URLParam(r, "id")
		id, err := strconv.Atoi(string_id)
		if err != nil {
			c.logger.Error("ERROR ATOI")
			c.error(w, r, http.StatusInternalServerError, err)
			return
		}

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
		string_id := chi.URLParam(r, "id")

		id, err := strconv.Atoi(string_id)
		if err != nil {
			c.logger.Error("ERROR ATOI")
			c.error(w, r, http.StatusInternalServerError, err)
			return
		}

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
		string_id := chi.URLParam(r, "user_id")

		id, err := strconv.Atoi(string_id)
		if err != nil {
			c.logger.Error("ERROR ATOI")
			c.error(w, r, http.StatusInternalServerError, err)
			return
		}

		blobs, err := c.store.Blob().GetAllBlobs(id)
		if err != nil {
			c.logger.Error("ERROR GETTING ALL BLOBS")
			c.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		c.respond(w, r, http.StatusAccepted, blobs)
	}
}
