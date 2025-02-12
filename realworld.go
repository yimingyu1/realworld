package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"realworld/cmd/api/internal/config"
	"realworld/cmd/api/internal/handler"
	"realworld/cmd/api/internal/svc"
	"realworld/common/xlog"
	"realworld/common/xvalidate"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "cmd/api/etc/realworld.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	fileWriter := logx.Reset()
	writer, err := xlog.NewMultiWriter(fileWriter)
	httpx.SetValidator(xvalidate.NewApiValidate())
	logx.Must(err)
	logx.SetWriter(writer)
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
