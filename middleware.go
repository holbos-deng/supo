package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func register(handlerFunc gin.HandlerFunc, g *gin.Engine) {
	g.Use(handlerFunc)
}

func middlewareInit(g *gin.Engine) {
	for _, f := range MiddleWare() {
		register(f, g)
	}
}

func mainMiddle(c *gin.Context) {

	defer catchException(func() {
		cr, ok := c.Get("dbActuator")
		var Cr *gorm.DB
		if ok {
			Cr = cr.(*gorm.DB)
			Cr.Commit()
		}
	}, func(e interface{}) {
		cr, ok := c.Get("dbActuator")
		var Cr *gorm.DB
		if ok {
			InternalError(c, fmt.Sprintf("%v", e))
			if ok {
				Cr = cr.(*gorm.DB)
				Cr.Rollback()
			}
		}
	})

	t := time.Now()
	fmt.Println("中间件开始执行了")
	c.Set("request", "中间件")
	c.Set("dbActuator", DB.Begin())
	c.Next()

	status := c.Writer.Status()
	fmt.Println("中间件执行完毕。结果状态:", status)
	t2 := time.Since(t)
	fmt.Println("time:", t2)
}

func MiddleWare() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		mainMiddle,
	}
}
