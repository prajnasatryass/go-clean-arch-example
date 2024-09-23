package main

import (
	"github.com/prajnasatryass/tic-be/config"
	"github.com/prajnasatryass/tic-be/internal/app"
	"github.com/samber/lo"
	"time"
)

func main() {
	cfg := lo.Must(config.LoadConfig())
	lo.Must(time.LoadLocation(cfg.Server.TimeZone))
	app.NewApp(cfg).Run()
}
