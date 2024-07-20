package fxloader

import (
	"ap_sell_products/api/controllers"
	"ap_sell_products/api/middleware"
	"ap_sell_products/api/routers"
	"ap_sell_products/core/infra"
	"ap_sell_products/core/infra/pgsql"
	"ap_sell_products/core/usecase"

	"go.uber.org/fx"
)

func Load() []fx.Option {
	return []fx.Option{
		fx.Options(loadAdapter()...),
		fx.Options(loadUseCase()...),
		fx.Options(loadEngine()...),
	}
}
func loadUseCase() []fx.Option {
	return []fx.Option{
		fx.Provide(usecase.NewJwtUseCasee),
		fx.Provide(usecase.NewUserUseCase),
		fx.Provide(usecase.NewOrderUseCase),
		fx.Provide(usecase.NewProductUseCase),
	}
}

func loadEngine() []fx.Option {
	return []fx.Option{
		fx.Provide(routers.NewApiRouter),
		fx.Provide(controllers.NewJwtController),
		fx.Provide(controllers.NewUserController),
		fx.Provide(middleware.NewMiddleware),
		fx.Provide(controllers.NewOrderController),
	}
}
func loadAdapter() []fx.Option {
	return []fx.Option{
		fx.Provide(infra.NewpostgreDb),
		fx.Provide(pgsql.NewCollectionUser),
		fx.Provide(pgsql.NewCollectionOrder),
		fx.Provide(pgsql.NewCollectionProduct),
		fx.Provide(pgsql.NewCollectionTransaction),
	}
}
