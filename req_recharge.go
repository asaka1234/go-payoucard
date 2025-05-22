package go_payoucard

import (
	"crypto/tls"
	"github.com/asaka1234/go-payoucard/utils"
	"github.com/mitchellh/mapstructure"
)

// 给银行卡充值
func (cli *Client) Recharge(req PayOuCardRechargeReq) (*PayOuCardRechargeRsp, error) {
	// 1. 拿到请求参数，转为map
	var signDataMap map[string]interface{}
	mapstructure.Decode(req, &signDataMap)

	// 2. 计算签名,补充参数
	signStr := utils.Sign(signDataMap, cli.RSAPrivateKey) //私钥加密
	signDataMap["signature"] = signStr

	var result PayOuCardRechargeRsp

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(signDataMap).
		SetHeaders(getHeaders()).
		SetResult(&result).
		SetError(&result).
		Post(cli.WithdrawURL)

	//fmt.Printf("result: %s\n", string(resp.Body()))

	if err != nil {
		return nil, err
	}

	return &result, nil
}
