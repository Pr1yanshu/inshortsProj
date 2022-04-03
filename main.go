package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopkg.in/natefinch/lumberjack.v2"
	"inshortsProj/constant"
	docs "inshortsProj/docs"
	covidRoutes "inshortsProj/router"
	"io"
	"os"
)

// @title Covid API
// @version 1.0
// @BasePath /api/covid

func main() {
	initializeServer()
}

func initializeServer() {
	f := &lumberjack.Logger{
		Filename:   constant.LogAccessFilePath,
		MaxSize:    1000, // megabytes
		MaxBackups: 3,
		Compress:   true,
	}
	env := constant.Development
	if env == constant.Production {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.MultiWriter(f)
	}
	gin.DefaultErrorWriter = io.MultiWriter(os.Stdout)

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s HTTP/1.1\" %d %d %s %s - %s\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 MST"),
			param.Method,
			param.Path,
			param.StatusCode,
			param.BodySize,
			param.Latency,
			param.Latency,
			param.Request.Header.Get("X-Forwarded-Host"),
		)
	}))
	router.Use(gin.Recovery())
	docs.SwaggerInfo.BasePath = "/api/covid"
	covidRoutesGroup := router.Group("/api/covid")
	covidRoutes.AddCovidRoutes(covidRoutesGroup)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	port := os.Getenv("PORT")
	if port == "" {
		port = constant.PORT
	}
	err := router.Run(":" + port)
	if err != nil {
		fmt.Println(err)
	}
}
