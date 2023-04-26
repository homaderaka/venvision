package apperror

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ginHandler func(c *gin.Context) error

// CustomMiddleware returns a middleware function that can be used with a Gin web framework.
// This middleware function is used to handle errors that might occur during the execution of a handler function.
func CustomMiddleware(h ginHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "application/json")
		var appErr *AppError
		err := h(c)
		if err != nil {
			if errors.As(err, &appErr) {
				if errors.Is(err, ErrNotFound) {
					c.JSON(http.StatusNotFound, ErrNotFound)
					return
				}
				err := err.(*AppError)
				c.JSON(http.StatusBadRequest, err)
				return
			}
			c.JSON(http.StatusBadRequest, systemError(err.Error()))
			return
		}
	}
}
