package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
)

var http = gin.New()

func run(host string, port string) {
	middlewareInit(http)
	registerDefaultRouter()
	routerBinding()
	err := http.Run(host + ":" + port)
	if err != nil {
		Log.Panicf("Server start error %s", err.Error())
	}
}

func setMode(mode string) {
	gin.SetMode(mode)
}

func registerDefaultRouter() {
	Router("POST", "call/:model/:method", callModelMethod)
}

func catchException(ok func(), error func(e interface{})) {
	if err := recover(); err != nil {
		e := printStackTrace(err)
		Log.Error(e)
		fmt.Println(e)
		error(err)
	} else {
		ok()
	}
}

func printStackTrace(err interface{}) string {
	buf := new(bytes.Buffer)
	_, _ = fmt.Fprintf(buf, "%v\n", err)
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		_, _ = fmt.Fprintf(buf, "%d：%s:%d (0x%x)\r\n", i, file, line, pc)

	}
	return buf.String()
}
