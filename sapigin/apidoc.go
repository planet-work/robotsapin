package main

import (
	"github.com/gin-gonic/gin"
)

func ApiDoc(c *gin.Context) {
	logD.Println(Engine.Routes())
}
