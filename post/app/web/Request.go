package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/shaileshhb/microservices/post/errors"
)

// UnmarshalJSON parses data from request and return otherwise error return.
func UnmarshalJSON(request *http.Request, out interface{}) error {
	if request.Body == nil {
		return errors.NewHTTPError(errors.ErrorCodeEmptyRequestBody, http.StatusBadRequest)
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return errors.NewHTTPError(errors.ErrorCodeReadWriteFailure, http.StatusBadRequest)
	}

	if len(body) == 0 {
		return errors.NewHTTPError(errors.ErrorCodeEmptyRequestBody, http.StatusBadRequest)
	}

	err = json.Unmarshal(body, out)
	if err != nil {
		return errors.NewHTTPError(errors.ErrorCodeInvalidJSON, http.StatusBadRequest)
	}
	return nil
}
