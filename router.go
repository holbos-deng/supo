package main

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"runtime"
	"strings"
)

type router struct {
	method string
	url    string
	fn     func(*gin.Context)
}

var tempRouterTable = make(map[string]router)

func Router(method string, url string, fn func(*gin.Context)) {
	if method == "" {
		return
	}
	method = strings.ToUpper(method)
	key := method + ":" + url
	if _, ok := tempRouterTable[key]; ok {
		_, file, line, _ := runtime.Caller(1)
		Log.Errorf("路由 %s 已经存在。%s %d行 注册 %v 失败",
			url,
			file, line,
			runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name())
		return
	}
	tempRouterTable[key] = router{method: strings.ToUpper(method), url: url, fn: fn}
}

func routerBinding() {
	for _, r := range tempRouterTable {
		if !_routerBinding(r) {
			r.method = r.method[0:1] + strings.ToLower(r.method[1:])
			_routerBinding(r)
		}
	}
	tempRouterTable = nil
}

func _routerBinding(r router) bool {
	handlerFn := reflect.ValueOf(http).MethodByName(r.method)
	if handlerFn.IsValid() {
		in := []reflect.Value{reflect.ValueOf(r.url), reflect.ValueOf(r.fn)}
		handlerFn.Call(in)
		return true
	}
	return false
}
