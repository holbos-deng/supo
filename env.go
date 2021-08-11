package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type environment struct {
	Cr   *gorm.DB
	Ctx  *context
	Http *gin.Context
	Uid  user
	self interface{}
	su   bool
	//Models iMO
}

type user struct {
	ID   uint
	Name string
}
