package main

import (
	"github.com/samber/lo"
	"tic-be/config"
	"tic-be/internal/app"
	"time"
)

func main() {
	cfg := lo.Must(config.LoadConfig())
	lo.Must(time.LoadLocation(cfg.Server.TimeZone))
	app.NewApp(cfg).Run()
}
