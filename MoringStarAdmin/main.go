package main

import (
	"os"

	"github.com/urfave/cli/v2"

)

// Usage: go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "v1.0.0"

// @title moringstaradmin
// @version v1.0.0
// @description 晨星项目后台管理系统
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http https
// @basePath /
func main() {
	app := cli.NewApp()
	app.Name = "moringstaradmin"
	app.Version = VERSION
	app.Usage = "晨星项目后台管理系统"
	app.Commands = []*cli.Command{
		cmd.StartCmd(),
		cmd.StopCmd(),
		cmd.VersionCmd(VERSION),
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
