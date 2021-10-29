package main

import (
    "github.com/gin-gonic/gin"
    _ "test1/docs"
    "net/http"
	 ginSwagger "github.com/swaggo/gin-swagger"
	  "github.com/swaggo/gin-swagger/swaggerFiles"
)

var swagHandler gin.HandlerFunc


// @title xiayuedu backend api
// @version 1.0
// @description this is xiayuedu backend server
// @termsOfService https://xiayuedu.com
//
// @contact.name xiayuedu.com
// @contact.url https://xiayuedu.com
// @contact.name www.xiayuedu.com
//
// @contact.email
// @license.name Apache 2.0
// @license.url https://xiayuedu.com
// @host 127.0.0.1:8089
// @BasePath /api/v1
func main() {
    engine := gin.Default()
    
        engine.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
   

    v1 := engine.Group("/api/v1")
    {
        v1.GET("/hello", HelloHandler)
    }
    engine.Run(":8089")

}

// @Summary hellohandler
// @Description hellohandler
// @Tags 测试
// @Accept json
// @Produce json
// @Param name query string true "名字"
// @Param age query string true "年龄"
// @Success 200 {string} string "{"msg":""hello razeen"}"
// @Failure 400 {string} string "{"msg":""who are you"}"
// @Router /hello [get]
func HelloHandler(ctx *gin.Context){
    name := ctx.Query("name")
    age := ctx.Query("age")
    if name == "" {
        ctx.JSON(http.StatusBadRequest,gin.H{"message":"who are you"})
        return
    }
    ctx.JSON(http.StatusOK,gin.H{"message":"hello " + name + " " + age})
}