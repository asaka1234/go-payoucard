package go_payoucard

import (
	"github.com/asaka1234/go-payoucard/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Params PayOuCardInitParams

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, params PayOuCardInitParams) *Client {
	return &Client{
		Params: params,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
