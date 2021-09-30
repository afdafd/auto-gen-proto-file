package common

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetId(ctx *gin.Context, field string, isPost bool) int32 {
	if isPost {
		baseSetId, _ := strconv.Atoi(ctx.PostForm(field))
		return int32(baseSetId)
	}

	baseSetId, _ := strconv.Atoi(ctx.Query(field))
	return int32(baseSetId)
}
