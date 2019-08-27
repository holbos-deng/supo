package model

import (
	"easonsu/supo/db"
	"time"
)

var models = make(map[string]interface{})

func GetModel(modelName string) interface{} {
	model, ok := models[modelName]
	if ok {
		return model
	}
	return BaseModel{}
}

func RegisterModel(modelName string, entity interface{}) {
	models[modelName] = entity
}

type BaseModel struct {
	CreateUid  int
	CreateDate time.Time
	WriteUid   int
	WriteDate  time.Time
}

func (base BaseModel) Create(value interface{}) {
	db.PG.Model(base).Create(value)
}

func (base BaseModel) Write(value interface{}) {
	switch value.(type) {
	case map[string]interface{}:
		db.PG.Model(base).Updates(value)
	default:
		db.PG.Model(base).Save(value)
	}
}

func (base BaseModel) Delete() {
	db.PG.Delete(base)
}

func (base BaseModel) Search(value interface{}) {
	db.PG.Model(base).Select(value)
}
