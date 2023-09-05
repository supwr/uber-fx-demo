package httpclient

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"
)

type Client struct {
	fx.Out

	PersonClient *resty.Client `name:"person"`
}

func NewHttpClient() Client {
	return Client{
		PersonClient: resty.New(),
	}
}
