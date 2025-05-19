package go_payoucard

import (
	"github.com/asaka1234/go-payoucard/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	MerchantID string // merchantId
	AccessKey  string // accessKey

	WithdrawURL string

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, merchantID string, accessKey, withdrawURL string) *Client {
	return &Client{
		MerchantID:  merchantID,
		AccessKey:   accessKey,
		WithdrawURL: withdrawURL,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
