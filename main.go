package main

import "github.com/gin-gonic/gin"

func main() {
	//初始化一个http的对象
	r := gin.Default()

	//设置一个GET请求路由，url为ping,处理函数是一个闭包函数。
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	gin.SetMode(gin.DebugMode) //设置debug模式
	//gin.SetMode(gin.ReleaseMode) //设置 release模式
	r.Run()

}
