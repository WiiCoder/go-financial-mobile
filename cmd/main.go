package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wiiCoder/installment/internal/api"
	"github.com/wiiCoder/installment/pkg/common/config"
	"github.com/wiiCoder/installment/pkg/common/db/mysql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// 初始化配置
	config.InitConfig()

	// 获取数据库连接
	db, err := mysql.NewMysql()
	if err != nil {
		panic(err)
	}

	gin.DisableConsoleColor() // 禁用控制台颜色

	router := gin.Default()

	// 控制台日志格式
	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf(" [%s] - [IP:%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.TimeStamp.Format(time.RFC1123),
			params.ClientIP,
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authApi := api.NewAuthApi(db)
	router.POST("/login", authApi.Login)

	srv := &http.Server{
		Addr:    config.Config.App.Addr,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号关闭服务器（设置 15 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
