package go_payoucard

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-payoucard/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
)

// 给银行卡充值
// https://github.com/lifeonearth718/payoucard-doc/blob/main/API-CN.md#%E9%93%B6%E8%A1%8C%E5%8D%A1%E5%85%85%E5%80%BC
func (cli *Client) Recharge(req PayOuCardRechargeReq) (*PayOuCardRechargeRsp, error) {
	//wrap成给上游的req
	apiReq := PayOuCardRechargeAPIReq{
		RequestID:  utils.GenRequestID(),
		MerchantID: cli.Params.MerchantId,
		Data:       req,
	}

	// 1. 拿到请求参数，转为map
	var signDataMap map[string]interface{}
	mapstructure.Decode(apiReq, &signDataMap)

	// 2. 计算签名,补充参数
	signStr := utils.Sign(signDataMap, cli.Params.RSAPrivateKey) //私钥加密
	signDataMap["signature"] = signStr

	var result PayOuCardRechargeRsp

	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(signDataMap).
		SetHeaders(getHeaders()).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(cli.Params.WithdrawUrl)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2))
	cli.logger.Infof("PSPResty#payoucard#withdraw->%+v", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp2.StatusCode() != 200 {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("status code: %d", resp2.StatusCode())
	}

	if resp2.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v, body:%s", resp2.Error(), resp2.Body())
	}

	return &result, nil
}
