package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type BaseHandler struct{}

func (bh *BaseHandler) getIdFromRequest(c *gin.Context) (int64, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, errors.New("id cant be 0")
	}

	return id, nil
}
