package main

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"
)

func createApp(o ...fx.Option) *fx.App {
	options := []fx.Option{
		fx.Provide(
			NewHttpClient,
			fx.Annotate(
				NewExternal,
				fx.As(new(ExternalRepository)),
			),
		),
	}

	return fx.New(append(options, o...)...)
}

func NewHttpClient() *resty.Client {
	return resty.New()
}
