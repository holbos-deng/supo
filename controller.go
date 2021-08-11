package main

import (
	"reflect"
	common2 "supo/common"
)

type baseModelDefault struct {
}

func (baseModelDefault) create(iModel IModel) interface{} {
	http := iModel.env().Http
	method := iModel.env().Ctx.Get("method").(string)
	err := http.BindJSON(&iModel)
	if err != nil {
		ParamterInvalid(http, err)
		return nil
	}

	data := call(iModel, method, nil)
	return data
}

func (baseModelDefault) searchRead(iModel IModel) interface{} {
	http := iModel.env().Http
	method := iModel.env().Ctx.Get("method").(string)
	var body common2.Domain
	err := http.BindJSON(&body)
	if err != nil {
		ParamterInvalid(http, err)
		return nil
	}

	in := []reflect.Value{reflect.ValueOf(body)}
	data := call(iModel, method, in)
	return data
}


var baseModelDefaultCRUD = &baseModelDefault{}
