package handler

import (
	// "html/template"
	// "log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Index(c *gin.Context) {
	c.HTML(200, "index", nil)
}
