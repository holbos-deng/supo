package main

import (
	"github.com/gin-gonic/gin"
	"reflect"
	str2 "supo/common/str"
)

type SearchRequest struct {
	Domain  []interface{}
	Context map[string]interface{}
	Fields  []string
	Page    Page
	Sort    string
}

type Page struct {
	Size   uint
	Number uint
}

func callModelMethod(c *gin.Context) {
	method := c.Param("method")
	c.Set("requestMode", "callModelMethod")
	decorator := _callMethod(c)
	var data interface{}
	switch method {
	case "create":
		data = decorator(baseModelDefaultCRUD.create)
	case "search_read":
		data = decorator(baseModelDefaultCRUD.searchRead)
	default:
		BadRequest(c, "Did not found method: "+method)
		return
	}

	var da interface{}
	if data != nil {
		da = data.([]interface{})[0]
	}

	Success(c, da)
}

func _callMethod(c *gin.Context) func(func(IModel) interface{}) interface{} {
	method := c.Param("method")
	modelName := c.Param("model")
	model := getModel(modelName).(IModel)
	model.initEnv(model, c)
	//TODO 绑定上下文
	ctx := &context{values: make(map[string]interface{})}
	ctx.Set("method", method)
	ctx.Set("model", modelName)
	model.env().Ctx = ctx

	var decorator = func(fn func(IModel) interface{}) interface{} {
		return fn(model)
	}
	return decorator
}

func getModel(name string) interface{} {
	v, ok := modelsStore[name]
	if ok {
		return instantiate(v.instance)
	}
	Log.Panicf("没有找到名字为[%s]的模型", name)
	return nil
}

func call(model interface{}, method string, in []reflect.Value) interface{} {
	method = str2.MethodNameHump(method)
	f := reflect.ValueOf(model).MethodByName(method)
	if f.IsValid() {
		returns := f.Call(in)
		returnsLen := len(returns)
		if returnsLen > 0 {
			var returnValue = make([]interface{}, returnsLen)
			for i := 0; i < returnsLen; i++ {
				returnValue[i] = returns[i].Interface()
			}
			return returnValue
		}
		return nil
	}
	Log.Panicf("方法[%s]不存在模型[%s]中", method, reflect.TypeOf(model).Kind().String())
	return nil
}
