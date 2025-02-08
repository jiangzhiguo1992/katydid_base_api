package handler

import (
	"github.com/gin-gonic/gin"
	"katydid_base_api/internal/team/service"
	"net/http"
	"strconv"
)

func GetClient(c *gin.Context) {
	name := c.Param("id")
	u, err := strconv.ParseUint(name, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID")
		return
	}
	client, codeError := service.GetClient(u)
	if codeError != nil {
		c.JSON(http.StatusNotFound, client)
	}
	c.JSON(http.StatusOK, client)
}
