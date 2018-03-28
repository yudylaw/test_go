package main

import (
	"feed_api/common"
	"feed_api/handler"
	"feed_api/service"
	"fmt"
	"git.inke.cn/BackendPlatform/golang/logging"
	"git.inke.cn/BackendPlatform/golang/tomlconfig"
	"git.inke.cn/BackendPlatform/golang/utils"
	rpc "git.inke.cn/inkelogic/rpc-go"
	api "git.inke.cn/video/panshi/panshi_api"
	"git.inke.cn/video/panshi/panshi_api/auth"
	"github.com/gin-gonic/gin"
	"io"
)

var (
	log     *logging.Logger
	AppPath string
	AppConf common.Config
)

type Config struct {
}

func init() {
	pwd := api.GetCurrentDir()
	AppPath = api.GetParentDir(pwd)
	configFile := AppPath + "/conf/main.toml"
	err := tomlconfig.ParseTomlConfig(configFile, &AppConf)
	if err != nil {
		fmt.Println("parse main config error.", err)
		return
	}

	clientFile := AppPath + "/conf/client.toml"
	var cc Config
	_, err = rpc.NewConfigToml(clientFile, &cc)
	if err != nil {
		fmt.Println("parse client config error.", err)
		return
	}

	accessLog := AppPath + "/logs/" + AppConf.Log.AccessLog
	businessLog := AppPath + "/logs/" + AppConf.Log.BusinessLog
	errorLog := AppPath + "/logs/" + AppConf.Log.ErrorLog
	AppConf.Log.StatLog = AppPath + "/logs/" + AppConf.Log.StatLog

	log = logging.NewLogger(&logging.Options{
		Level:   AppConf.Log.Level,
		Rolling: logging.DAILY,
	}, accessLog, businessLog, errorLog)
}

func main() {

	feedService := service.NewFeedService()
	authService := auth.NewAuthService("auth_service")
	feedHandler := handler.NewFeedHandle(feedService)

	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(log.GetOutput())

	api.BusinessLog().Infof("stat log path=%s", AppConf.Log.StatLog)

	utils.SetStat(AppConf.Log.StatLog, "feed_api")

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(api.Recovery())

	v1 := router.Group("/api/v1")
	router.NoRoute(api.NotFoundHandle)
	//鉴权走以下路径
	authV1 := v1.Group("/")
	// per group middleware! in this case we use the custom created
	// CheckSessionMiddleWare() middleware just in the "authorized" group.
	authV1.Use(auth.AuthMiddleWare(authService))
	{
		authV1.PUT("/feeds", feedHandler.CreateOrderHandler)
		//		authV1.DELETE("/feeds/:feed_id", paymentHandler.ApplePayVerifyHandler)
		//		authV1.GET("/feeds/latest_feeds", paymentHandler.GetWealthHandler)
		//		authV1.GET("/feeds", paymentHandler.GetIncomeHandler)
		//		authV1.POST("/feeds/:feed_id/like", paymentHandler.ApplyHandler)
		//		authV1.POST("/feeds/:feed_id/unlike", paymentHandler.GetApplyHandler)
	}

	router.Run(AppConf.Server.Bind)
}
