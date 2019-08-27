package net

import "github.com/gin-gonic/gin"

var g *gin.Engine

func init() {
	g = gin.Default()
}

func Router() {
	g.GET()
}