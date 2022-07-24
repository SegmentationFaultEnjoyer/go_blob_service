package errors

import (
	"net/http"
	"strconv"

	"github.com/google/jsonapi"
)

func UnprocessebleEntity(detail string) *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Title:  "Unprocesseble Entity",
		Status: strconv.Itoa(http.StatusUnprocessableEntity),
		Detail: detail,
	}
}

func InternalError(detail string) *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Title:  "Internal server error",
		Status: strconv.Itoa(http.StatusInternalServerError),
		Detail: detail,
	}
}

func BadRequest(detail string) *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Title:  "Bad Request",
		Status: strconv.Itoa(http.StatusBadRequest),
		Detail: detail,
	}
}
