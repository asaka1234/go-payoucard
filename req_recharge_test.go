package go_payoucard

import (
	"fmt"
	"testing"
)

func TestRecharge(t *testing.T) {

	//构造client
	cli := NewClient(nil, MERCHANT_ID, RAS_PUBLIC_KEY, RAS_PRIVATE_KEY, WITHDRAW_URL)

	//发请求
	resp, err := cli.Recharge(GenRechargeRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenRechargeRequestDemo() PayOuCardRechargeReq {
	return PayOuCardRechargeReq{
		UniqueID: "123", //商户uid
		CardNo:   "30779639363",
		Currency: "USDT",
		Amount:   600.00,
		OrderNo:  "807936863", //商户订单号
	}
}
