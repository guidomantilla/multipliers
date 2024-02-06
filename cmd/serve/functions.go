package serve

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	feather_boot "github.com/guidomantilla/go-feather-boot/pkg/boot"
	feather_commons_config "github.com/guidomantilla/go-feather-commons/pkg/config"
	feather_commons_log "github.com/guidomantilla/go-feather-commons/pkg/log"
	"github.com/spf13/cobra"

	"multipliers/pkg/config"
	"multipliers/pkg/endpoints"
	"multipliers/pkg/services"
)

func ExecuteCmdFn(_ *cobra.Command, args []string) {

	ctx := context.Background()
	logger := feather_commons_log.Custom()
	appName, version := config.Application, config.Version
	enablers := &feather_boot.Enablers{
		HttpServerEnabled: true,
		GrpcServerEnabled: false,
		DatabaseEnabled:   false,
	}

	builder := feather_boot.NewBeanBuilder(ctx)
	builder.Config = func(appCtx *feather_boot.ApplicationContext) {
		var cfg config.Config
		if err := feather_commons_config.Process(ctx, appCtx.Environment, &cfg); err != nil {
			feather_commons_log.Fatal(fmt.Sprintf("starting up - error setting up configuration: %s", err.Error()))
		}

		appCtx.HttpConfig = &feather_boot.HttpConfig{
			Host: cfg.Host,
			Port: cfg.HttpPort,
		}
	}
	err := feather_boot.Init(appName, version, args, logger, enablers, builder, func(appCtx feather_boot.ApplicationContext) error {

		numbersService := services.NewDefaultNumbersService()
		numbersEndpoint := endpoints.NewDefaultNumbersEndpoint(numbersService)

		appCtx.PublicRouter.POST("/numbers", numbersEndpoint.Save)
		appCtx.PublicRouter.GET("/numbers", numbersEndpoint.GetAll)
		appCtx.PublicRouter.GET("/numbers/:number", numbersEndpoint.Get)

		return nil
	})
	if err != nil {
		feather_commons_log.Fatal(err.Error())
	}
}
