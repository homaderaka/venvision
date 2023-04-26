package product

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"venvision/internal/apperror"
	"venvision/internal/product/service"
	"venvision/pkg/logging"
	"venvision/protofiles"
)

type Handler struct {
	Logger         logging.Logger
	ProductService service.Service
}

func (h *Handler) Register(group *gin.RouterGroup) {
	group.Handle(http.MethodGet, "/all", apperror.CustomMiddleware(h.GetAll))
}

func (h *Handler) GetAll(c *gin.Context) (err error) {
	p, err := h.ProductService.GetAll(c.Request.Context(), &protofiles.GetAllRequest{})
	if err != nil {
		return
	}
	err = json.NewDecoder(c.Request.Body).Decode(p.GetProducts())
	if err != nil {
		return
	}

	c.Status(http.StatusOK)
	return
}
