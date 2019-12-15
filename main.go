package main

import (
	"github.com/WangPengJu/leafserver/conf"
	"github.com/WangPengJu/leafserver/game"
	"github.com/WangPengJu/leafserver/gate"
	"github.com/WangPengJu/leafserver/login"
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
