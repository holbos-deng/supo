package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"eason/supo"
)

func init() {
	supo.Models().Register("student", &Student{})
	Router("any", "/show_name", func(c *gin.Context) {
		//Student{Name: "小明"}.GetName()
		student := Models().Get("student") // 提供了根据model名获取model实例的可能性。根据具体场景决定是否使用
		student.Init(map[string]interface{}{
			"first_name": "小",
			"last_name":  "名",
		})
		fmt.Println(student, "==========================---------------------===========================")
		c.String(200, student.(Student).GetName())
	})
}

type Student struct {
	BaseModel
	FistName string
	LastName string
	Age      int
	Sex      int8
}

func (s Student) GetName() string {
	stu := Models().Get("student")
	Log.Println("get model:", stu)
	return "111"
}

func main() {
	Log.Infoln("服务启动......")
	main.Start("0.0.0.0", "9999")
}
