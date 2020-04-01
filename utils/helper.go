package utils

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseParamID(c *gin.Context) (uint, error) {
	id := c.Param("id")
	parseID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, errors.New("id must be an unsigned int")
	}
	return uint(parseID), nil
}
