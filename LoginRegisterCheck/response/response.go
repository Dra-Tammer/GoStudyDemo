package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//封装我们的http返回，一般我们开发的时候都希望我们返回形式统一例如
/*
{
	code:200,
	data:xxx,
	msg:xxx,
}
*/

// 一个是http的标准状态码，一个是业务状态码，一个给前端看，一个给自己看
func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, msg string, data gin.H) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
