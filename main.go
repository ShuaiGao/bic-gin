package main

import (
	"fmt"
	"gin-bic/config"
	"gin-bic/internal"
	"gin-bic/internal/schema"
	"gin-bic/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title         gin-bic api文档
// @version       1.0
// @contact.name  GaoZiJia
// @contact.email boringmanman@qq.com

// @host     api.gin-bic.cn
// @BasePath /api/gin-bic

var (
	gitTag      string
	gitHash     string
	buildTime   string
	goVersion   string
	nodeVersion string // 节点控制中使用的版本号，优先使用tag号
)

func MustInit() {
	config.MustSetupYaml()
	db.MustInitMysql(config.MysqlSetting)
	schema.AutoMigrate()
}

func main() {
	MustInit()
	gin.SetMode(config.ServerSetting.RunMode)
	ginInstance := internal.SetupRouter()
	readTimeout := time.Duration(config.ServerSetting.ReadTimeoutSecond) * time.Second
	writeTimeout := time.Duration(config.ServerSetting.WriteTimeoutSecond) * time.Second
	endPoint := fmt.Sprintf(":%d", config.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        ginInstance,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	hostname, _ := os.Hostname()
	log.Printf("[info] start http server listening, host:%s %s", hostname, endPoint)
	log.Printf("[info] start http server mongodb, %s %s", config.MongoDBSetting.Uri, config.MongoDBSetting.Database)
	log.Printf("[info] start http server sql, %s %s", config.MysqlSetting.Host, config.MysqlSetting.Database)
	log.Printf("[info] compile param %v %v %v, RunMode=%v", goVersion, gitHash, gitTag, config.ServerSetting.RunMode)

	// https://golang.org/pkg/os/signal/#Notify
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR2, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	log.Printf("pid-------> %d", os.Getpid())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("receive interrupt signal")
		if err := server.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()
	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpect")
		}
	}

	log.Println("Server exiting")
}
