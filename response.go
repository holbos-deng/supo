package main

import (
	"github.com/gin-gonic/gin"
)

type callModelMethodResponseBody struct {
	Code    int         `json:"code"`
	Result  interface{} `json:"result"`
	Message string      `json:"message"`
	Model   string      `json:"model"`
	Method  string      `json:"method"`
}

func Fail(c *gin.Context, data interface{}) {
	code := 100
	response(code, data, c)
}

func Success(c *gin.Context, data interface{}) {
	code := 200
	response(code, data, c)
}

func BadRequest(c *gin.Context, data interface{}) {
	code := 400
	response(code, data, c)
}

func IdentityAuthenticationFaild(c *gin.Context, data interface{}) {
	code := 401
	response(code, data, c)
}

func Forbidden(c *gin.Context, data interface{}) {
	code := 403
	response(code, data, c)
}

func ResourceNotFound(c *gin.Context, data interface{}) {
	code := 404
	response(code, data, c)
}

func InternalError(c *gin.Context, data interface{}) {
	code := 500
	response(code, data, c)
}

func ParamterCanNotNull(c *gin.Context, data interface{}) {
	code := 600
	response(code, data, c)
}

func ParamterInvalid(c *gin.Context, data interface{}) {
	code := 601
	response(code, data, c)
}

func DataNul(c *gin.Context, data interface{}) {
	code := 602
	response(code, data, c)
}

func ReadExcelError(c *gin.Context, data interface{}) {
	code := 603
	response(code, data, c)
}

func DataExists(c *gin.Context, data interface{}) {
	code := 604
	response(code, data, c)
}

var errorStatusMessage = map[int][2]string{
	100: {"FAIL", "操作失败"},
	200: {"SUCCESS", "操作成功"},
	400: {"BAD REQUEST", "错误的请求"},
	401: {"IDENTITY AUTHENTICATION FAILED", "身份认证失败"},
	403: {"FORBIDDEN", "没有权限"},
	404: {"RESOURCE NOT FOUND", "资源不存在"},
	500: {"INTERNAL ERROR", "服务器处理失败"},
	600: {"PARAMETER CAN NOT NULL", "缺少必要参数"},
	601: {"PARAMETER INVALID", "非法参数"},
	602: {"DATA NULL", "数据为空"},
	603: {"READ EXCEL ERROR", "读取excel表格数据异常"},
	604: {"DB FIELD REPETITION", "数据已经存在"},
	20:  {"UNKNOWN ERROR", "未知错误"},
}

func response(code int, data interface{}, c *gin.Context) {
	var message string
	messages, ok := errorStatusMessage[code]
	if !ok {
		messages, ok = errorStatusMessage[code]
	}
	message = messages[0] + "(" + messages[1] + ")"
	mode, ok := c.Get("requestMode")
	var body interface{}
	if ok {
		switch mode {
		case "callModelMethod":
			body = &callModelMethodResponseBody{
				Code:    code,
				Result:  data,
				Message: message,
				Model:   c.Param("model"),
				Method:  c.Param("method"),
			}
		default:
			body = struct {
				Code    int         `json:"code"`
				Result  interface{} `json:"result"`
				Message string      `json:"message"`
			}{code, body, message}
		}
	}

	c.JSON(code, body)
}
