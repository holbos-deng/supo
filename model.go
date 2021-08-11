package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goinggo/mapstructure"
	"gorm.io/gorm"
	"reflect"
	common2 "supo/common"
)

type IModel interface {
	Create() uint
	Update(map[string]interface{}) bool
	Delete(int) bool
	Search(common2.Domain) []uint
	SearchRead(domain common2.Domain) interface{}
	Read(uint) interface{}
	initEnv(interface{}, *gin.Context)
	Init(values map[string]interface{})
	env() *environment
}

type BaseModel struct {
	gorm.Model
	Env        *environment `json:"-" gorm:"-"`
	CreatedUid uint
	UpdatedUid uint
	DeletedUid uint
}

var models IModels = &model{}

func (base BaseModel) Create() uint {
	base.Env.Cr.Model(base.Env.self).Create(base.Env.self)
	return base.getId()
}

func (base BaseModel) Update(map[string]interface{}) bool {
	return false
}

func (base BaseModel) Delete(int) bool {
	return false
}

func (base BaseModel) Search(domain common2.Domain) []uint {
	var ids []uint
	base.Env.Cr.Model(base.Env.self).Where("").Select("id").Find(&ids)
	return ids
}

func (base BaseModel) SearchRead(domain common2.Domain) interface{} {
	Log.Println(domain, "domain")
	expr := expression(domain)
	Log.Println(expr, "exxxxxxxxxxxxxxxxxxxxxxpre")

	var out = reflect.New(reflect.SliceOf(reflect.TypeOf(base.Env.self).Elem())).Interface()
	base.Env.Cr.Model(base.Env.self).Where(expr).Find(out)

	Log.Println(out, "=================out===============", reflect.TypeOf(&out).Elem(), reflect.TypeOf(out))
	return out
}

func (base BaseModel) Read(uint) interface{} {
	return nil
}

func (base BaseModel) CheckAccessRights(operation string) bool {
	if base.Env.su {
		return true
	}
	m := getModel("supo.model.access")
	r:=reflect.ValueOf(m).MethodByName("check").Call([]reflect.Value{reflect.ValueOf(operation)})
	if len(r) > 0 {
		return r[0].Interface().(bool)
	}
	return false
}

func (base BaseModel) Sudo(flag bool)  {
	base.Env.su = flag
}

func (base BaseModel) Init(values map[string]interface{}) {
	err := mapstructure.Decode(values, &base)
	if err != nil {
		Log.Println(err)
	}
}

func (base BaseModel) initEnv(model interface{}, c *gin.Context) {
	cr, ok := c.Get("dbActuator")
	if !ok {
		Log.Panic("The database operation object is not initialized")
	}
	base.Env = &environment{
		Cr:   cr.(*gorm.DB),
		self: model,
		Uid:  user{ID: 0, Name: "admin"},
		Http: c,
		//Models: model{},
	}

	Log.Println(base.Env, model, "initEnv")
}

func (base BaseModel) env() *environment {
	return base.Env
}

func (base *BaseModel) getId() uint {
	return uint(reflect.ValueOf(base.Env.self).Elem().FieldByName("ID").Uint())
}

var modelsStore = make(map[string]*model)

func instantiate(model interface{}) IModel {
	t := reflect.ValueOf(model).Elem().Type()
	v := reflect.New(t).Interface()
	return v.(IModel)
}

type IModels interface {
	Get(string) IModel
	Register(string, IModel)
}

type model struct {
	instance        IModel
	fieldsMany2one  []string
	fieldsMany2Many []string
	fieldsProperty  []string
}

func (model) Register(modelName string, iModel IModel) {
	autoMigrate(iModel)
	ms := model{instance: iModel}
	modelsStore[modelName] = &ms
}

func (model) Get(name string) IModel {
	model, ok := modelsStore[name]
	if ok {
		return instantiate(model.instance)
	}
	return nil
}

func Models() IModels {
	return models
}
