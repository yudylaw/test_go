package main

import (
	"flag"
	"fmt"
	log "git.inke.cn/BackendPlatform/golang/logging"
	"git.inke.cn/inkelogic/rpc-go"
	"git.inke.cn/video/starstar/server/starstar.link.service/src/handler"
	"git.inke.cn/video/starstar/server/starstar.link.service/src/client"
	"git.inke.cn/video/starstar/server/starstar.link.service/src/define"
	"os"
)

func main() {
	conf := flag.String("config", "./conf/config.toml", "rpc config file")
	flag.Parse()
	config, err := rpc.NewConfigToml(*conf, &define.Config)
	if err != nil {
		log.Info("parse config toml fail,err:", err)
		fmt.Print("parse config toml fail,err:", err)
		os.Exit(0)
	}

	//初始化action_client
	err = client.InitThriftDeal()
	if err != nil {
		log.Errorf("InitThriftDeal error,%+v", err)
		fmt.Println("InitThriftDeal error,%+v", err)
		os.Exit(0)
	}

	server := rpc.NewHTTPServerWithConfig(config)
	err = server.Register(
		&handler.LinkSlotAddrHandler{})
	if err != nil {
		log.Info("register handler failed,err:", err)
		fmt.Print("register handler failed,err:", err)
		os.Exit(0)
	}

	log.Info("start server ["+config.GetServiceName()+"],port:", config.Port())

	err = server.Serve(config.Port())

	log.Info("shutdown server ["+config.GetServiceName()+"],port:", config.Port(), "err:", err)
}
