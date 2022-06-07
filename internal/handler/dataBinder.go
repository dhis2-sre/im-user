package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func DataBinder(c *gin.Context, req interface{}) error {
	return c.MustBindWith(req, binding.Default(c.Request.Method, c.ContentType()))
}
