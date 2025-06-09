package go_payoucard

import (
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestRecharge(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &PayOuCardInitParams{MERCHANT_ID, RAS_PUBLIC_KEY, RAS_PRIVATE_KEY, WITHDRAW_URL})

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
