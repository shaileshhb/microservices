package web

import (
	"github.com/gin-gonic/gin"
)

// UnmarshalJSON parses data from request and return otherwise error return.
func UnmarshalJSON(c *gin.Context, out interface{}) error {
	// var body interface{}

	err := c.ShouldBindJSON(out)
	if err != nil {
		return err
	}

	// fmt.Println(" ================ tst ->", out)

	// if body == nil {
	// 	return err
	// }

	// if len(body) == 0 {
	// 	return err
	// }

	// err = json.Unmarshal(body, out)
	// if err != nil {
	// 	return err
	// }
	return nil
}
