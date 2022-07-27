package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type BlobListRequest struct {
	pgdb.OffsetPageParams
	FilterAuthor *string `filter:"author_id"`
}

func GetBlobListRequest(r *http.Request) (BlobListRequest, error) {
	request := BlobListRequest{}

	if err := urlval.Decode(r.URL.Query(), &request); err != nil {
		return request, err
	}

	return request, nil
}
