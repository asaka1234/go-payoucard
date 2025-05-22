package go_payoucard

import (
	"github.com/asaka1234/go-payoucard/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	MerchantID    string // merchantId
	RSAPublicKey  string // 公钥
	RSAPrivateKey string // 私钥

	WithdrawURL string

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, merchantID string, rsaPublicKey, rsaPrivateKey, withdrawURL string) *Client {
	return &Client{
		MerchantID:    merchantID,
		RSAPublicKey:  rsaPublicKey,
		RSAPrivateKey: rsaPrivateKey,
		WithdrawURL:   withdrawURL,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
