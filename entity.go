package go_payoucard

// ---------------------------------------------
// 这个是业务上游请求参数
type PayOuCardRechargeReq struct {
	UniqueID string  `json:"uniqueId" mapstructure:"uniqueId"` //合作商用户的唯一ID (merchantId)
	CardNo   string  `json:"cardNo" mapstructure:"cardNo"`     //银行卡号
	Currency string  `json:"currency" mapstructure:"currency"` //币种(默认USDT)
	Amount   float64 `json:"amount" mapstructure:"amount"`     //充值金额
	OrderNo  string  `json:"orderNo" mapstructure:"orderNo"`   //商户订单号
}

// 这个是实际发送给上游api的
type PayOuCardRechargeAPIReq struct {
	RequestID  string               `json:"requestId" mapstructure:"requestId"`   // 请求流水id。20位随机字符
	Data       PayOuCardRechargeReq `json:"data" mapstructure:"data"`             // 业务数据
	MerchantID string               `json:"merchantId" mapstructure:"merchantId"` // 商户ID。PayouCard为商户分配
	Signature  string               `json:"signature" mapstructure:"signature"`   // 签名 这个不用业务传.而是sdk来计算
}

// response
type PayOuCardRechargeRsp struct {
	Code       int                       `json:"code"`
	Message    string                    `json:"message"`
	Success    bool                      `json:"success"` //是否成功。true：成功；false：失败。 当code=0时为true
	Data       *PayOuCardRechargeRspData `json:"data"`
	RequestID  string                    `json:"requestId"`
	MerchantID string                    `json:"merchantId"`
	Signature  string                    `json:"signature"`
}

type PayOuCardRechargeRspData struct {
	Status         int     `json:"status"`         //卡片充值状态。1：成功；2：失败；3：处理中
	CardNo         string  `json:"cardNo"`         //卡号
	OrderNo        string  `json:"orderNo"`        //商户订单号
	Currency       string  `json:"currency"`       //币种
	RechargeAmount float64 `json:"rechargeAmount"` //充值金额
	ReceivedAmount float64 `json:"receivedAmount"` //到账金额
	Fee            float64 `json:"fee"`            //手续费(扣充值金额之外的钱)
	Msg            string  `json:"msg"`
}

//--------------callback------------------------------

type PayOuCardRechargeBackReq struct {
	RequestID  string                       `json:"requestId" mapstructure:"requestId"`   // 请求流水id。20位随机字符
	MerchantID string                       `json:"merchantId" mapstructure:"merchantId"` // 商户ID
	NotifyType int                          `json:"notifyType" mapstructure:"notifyType"` // 通知类型 此通知notifyType = 4
	Data       PayOuCardRechargeBackReqData `json:"data" mapstructure:"data"`             // 提现数据
	Signature  string                       `json:"signature" mapstructure:"signature"`   // 签名
}

type PayOuCardRechargeBackReqData struct {
	CardNo         string  `json:"cardNo" mapstructure:"cardNo"`
	Status         int     `json:"status" mapstructure:"status"`   //卡片充值状态。1：成功；2：失败；
	OrderNo        string  `json:"orderNo" mapstructure:"orderNo"` //商户订单号
	Currency       string  `json:"currency" mapstructure:"currency"`
	RechargeAmount float64 `json:"rechargeAmount" mapstructure:"rechargeAmount"`
	ReceivedAmount float64 `json:"receivedAmount" mapstructure:"receivedAmount"` //option
	Fee            float64 `json:"fee" mapstructure:"fee"`
	Msg            string  `json:"msg" mapstructure:"msg"` //option
}

// 给callback的response
type PayOuCardRechargeBackResp struct {
	Code    int    `json:"code"`    // 响应状态码
	Message string `json:"message"` // 响应消息
}
