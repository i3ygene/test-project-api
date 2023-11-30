package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Test() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		msg := "test controller"
		fmt.Println(msg)
		status := 200
		ctx.Set("ResponseBody", msg)
		ctx.JSON(status, msg)
	}
}
