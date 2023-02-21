package router

import (
	v1 "my-admin/api/v1"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct{ gin.RouterGroup }

func (r *BaseRouter) InitRouter() {
	userGroup := r.Group("user")
	userGroup.Use()
	//TODO:test need delete
	userGroup.GET("/admin_register", v1.TestRespose)

	userGroup.POST("/register", v1.Register)
	userGroup.POST("/login", v1.Login)

}
