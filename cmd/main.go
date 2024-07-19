// cmd/main.go

package main

import (
	"ap_sell_products/api/routers"
	"ap_sell_products/common/configs"
	"ap_sell_products/common/log"
	fxloader "ap_sell_products/loader"
	"ap_sell_products/mcache"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"go.uber.org/fx"
)

func init() {
	log.NewLogger()
	var pathConfig string
	flag.StringVar(&pathConfig, "configs", "configs/configs.json", "path config")
	flag.Parse()
	configs.LoadConfig(pathConfig) // Load configuration
	mcache.Init(configs.Get())
}
func serverLifecycle(lc fx.Lifecycle, cfg *configs.Configs, apiRouter *routers.ApiRouter) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				port := cfg.Port
				if port == "" {
					port = "3000" // Default port if not specified in the config
				}
				address := fmt.Sprintf(":%s", port)
				if err := apiRouter.Router.Listen(address); err != nil {
					log.Fatal("Error starting server", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return apiRouter.Router.Shutdown()
		},
	})
}
func main() {
	app := fx.New(
		fx.Provide(configs.Get),
		fx.Options(fxloader.Load()...),
		fx.Invoke(serverLifecycle),
		fx.Options(), // No need for conditional logic with nopLogger
	)

	// Run the application
	if err := app.Start(context.Background()); err != nil {
		log.Fatal("Error starting application", err)
	}

	// Wait for an interrupt signal to gracefully shut down the application
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	// Shut down the application gracefully
	if err := app.Stop(context.Background()); err != nil {
		log.Fatal("Error stopping application", err)
	}
}
