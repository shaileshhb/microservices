package web

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/shaileshhb/microservices/post/app/errors"
)

// UnmarshalJSON parses data from request and return otherwise error return.
func UnmarshalJSON(c *fiber.Ctx, out interface{}) error {
	body := c.Body()

	if body == nil {
		return errors.NewHTTPError(errors.ErrorCodeEmptyRequestBody, http.StatusBadRequest)
	}

	if len(body) == 0 {
		return errors.NewHTTPError(errors.ErrorCodeEmptyRequestBody, http.StatusBadRequest)
	}

	err := json.Unmarshal(body, out)
	if err != nil {
		return errors.NewHTTPError(errors.ErrorCodeInvalidJSON, http.StatusBadRequest)
	}
	return nil
}
