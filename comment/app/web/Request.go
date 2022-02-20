package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// UnmarshalJSON parses data from request and return otherwise error return.
func UnmarshalJSON(c *gin.Context, out interface{}) error {
	// var body interface{}

	err := c.BindJSON(out)
	if err != nil {
		return err
	}

	fmt.Println(" ================ tst ->", out)

	// if out == nil {
	// 	return err
	// }

	// if len(out) == 0 {
	// 	return err
	// }

	// err = json.Unmarshal(body, out)
	// if err != nil {
	// 	return err
	// }
	return nil
}
