package main

import (
	"github.com/supwr/uber-fx-demo/v5/httpclient"
	"go.uber.org/fx"
)

func createApp(o ...fx.Option) *fx.App {
	options := []fx.Option{
		fx.Provide(
			httpclient.NewHttpClient,
			fx.Annotate(
				NewExternal,
				fx.As(new(ExternalRepository)),
			),
		),
	}

	return fx.New(append(options, o...)...)
}
