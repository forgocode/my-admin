package router

import (
	//"my-admin/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router interface {
	InitRouter()
}

type AllRouter struct {
	BaseRouter
	SystemRouter
}

func StartServer() error {
	// engine := gin.New()

	// engine.Use(middleware.Logger(), middleware.Recovery())
	engine := gin.Default()
	//TODO:(forgocode) need to update path and dir
	engine.StaticFS("/file", http.Dir(""))
	// r := new(AllRouter)
	// r.BaseRouter.RouterGroup = engine.RouterGroup
	// // r.SystemRouter = engine
	// r.BaseRouter.InitRouter()
	// r.SystemRouter.InitRouter()

	engine.POST("login")

	// InitRouter()
	// InitRouter()
	// baseRouter := engine.Group("base")

	// systemGroup := engine.Group("system")

	//TODO: api list
	//登录
	//验证码
	//健康检查，探针
	//检查数据库是否需要初始化
	//初始化数据库
	//注册api，增删改查
	//jwt相关路由，加入黑名单

	//管理员注册
	//修改密码
	//设置用户权限
	//删除用户
	//修改用户信息
	//修改自身用户信息
	//设置用户权限组
	//忘记密码，重置
	//分页获取用户列表
	//获取用户信息

	//新增菜单
	//新增用户和菜单的关联关系
	//删除菜单
	//更新菜单
	//分页获取菜单树列表
	//获取指定角色的菜单
	//根据id获取菜单

	//获取配置文件内容
	//修改配置文件内容
	//获取服务器信息
	//重启服务

	//创建自动化代码

	//系统路由
	// example路由
	//加载swagger

	//加载基础路由

	//健康检查
	//TODO: 页面上查看日志
	//系统路由+ 基础功能相关的路由，不需要鉴权

	return engine.Run(":8080")

}
