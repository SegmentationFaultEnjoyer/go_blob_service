package requests

import (
	"encoding/json"
	"net/http"
	"testService/internal/service/helpers"
	"testService/resources"
)

type BlobCreateRequest struct {
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

func ParseBlobRequest(w http.ResponseWriter, r *http.Request) (*resources.Blob, error) {
	req := &BlobCreateRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return nil, err
	}

	blob, err := helpers.GenerateBlob(
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
		return nil, err
	}

	return blob, nil
}
