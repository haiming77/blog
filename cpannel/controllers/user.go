package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
	ApplicationController
}

func (uc UserController) Signup(c *gin.Context) {

	c.JSON(1, "")
}
