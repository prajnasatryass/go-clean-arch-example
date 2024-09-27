package main

import (
	"github.com/prajnasatryass/go-clean-arch-example/config"
	"github.com/prajnasatryass/go-clean-arch-example/internal/app"
	"github.com/samber/lo"
	"time"
)

func main() {
	cfg := lo.Must(config.LoadConfig())
	lo.Must(time.LoadLocation(cfg.Server.TimeZone))
	app.NewApp(cfg).Run()
}
