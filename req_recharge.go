package go_payoucard

import (
	"crypto/tls"
	"github.com/asaka1234/go-payoucard/utils"
	"github.com/mitchellh/mapstructure"
)

// 给银行卡充值
// https://github.com/lifeonearth718/payoucard-doc/blob/main/API-CN.md#%E9%93%B6%E8%A1%8C%E5%8D%A1%E5%85%85%E5%80%BC
func (cli *Client) Recharge(req PayOuCardRechargeReq) (*PayOuCardRechargeRsp, error) {
	//wrap成给上游的req
	apiReq := PayOuCardRechargeAPIReq{
		RequestID:  utils.GenRequestID(),
		MerchantID: cli.MerchantID,
		Data:       req,
	}

	// 1. 拿到请求参数，转为map
	var signDataMap map[string]interface{}
	mapstructure.Decode(apiReq, &signDataMap)

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
