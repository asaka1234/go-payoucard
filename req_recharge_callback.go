package go_payoucard

import (
	"errors"
	"fmt"
	"github.com/asaka1234/go-payoucard/utils"
	"github.com/mitchellh/mapstructure"
)

// https://github.com/lifeonearth718/payoucard-doc/blob/main/API-CN.md#%E9%93%B6%E8%A1%8C%E5%8D%A1%E5%8D%A1%E7%89%87%E5%85%85%E5%80%BC%E7%BB%93%E6%9E%9C%E5%9B%9E%E8%B0%83%E9%80%9A%E7%9F%A5
func (cli *Client) RechargeCallback(req PayOuCardRechargeBackReq, processor func(PayOuCardRechargeBackReq) error) error {
	//验证
	sign := req.Signature //收到的签名
	//再自己计算一个
	var paramMap map[string]interface{}
	mapstructure.Decode(req, &paramMap)
	delete(paramMap, "signature") //去掉，用余下的来计算签名

	verifyResult := utils.VerifySign(paramMap, cli.RSAPublicKey, sign)
	if !verifyResult {
		fmt.Println("签名验证失败")
		return errors.New("sign verify failed!")
	}

	//开始处理
	return processor(req)
}
