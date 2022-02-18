package main

import (
	"casbin-learning/lib"
	"github.com/gin-gonic/gin"
)

// 1. 项目启动后删除casbin_rule表中数据，权限系统依然能够运行，由此判断casbin_rule的数据会被载入内存中，并没实时查表
// 2. 配置文件的内容也不会热更新
// 3. cas_bin先加载配置文件和数据源，之后可以通过代码动态添加
func main() {
	lib.InitAccess()
	r := gin.New()
	r.Use(lib.Middlewares()...)
	r.GET("/depts", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"result": "部门列表"})
	})
	r.POST("/depts", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"result": "批量修改部门列表"})
	})
	r.PUT("/depts/:id", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"result": "新增部门"})
	})

	r.GET("/info", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"result": "信息列表"})
	})
	r.POST("/info", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"result": "批量修改信息列表"})
	})

	r.POST("/admin", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"result": "超级管理员..."})
	})

	_ = r.Run(":8080")
}
