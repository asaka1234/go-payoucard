package go_payoucard

//---------------------------------------------

type PayOuCardWithdrawReq struct {
	RequestID  string         `json:"requestId"`
	MerchantID string         `json:"merchantId"`
	Data       *PayOuCardData `json:"data"`
}

type PayOuCardData struct {
	UniqueID string  `json:"uniqueId"`
	CardNo   string  `json:"cardNo"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"` // Using big.Float for precise decimal representation
	OrderNo  string  `json:"orderNo"`
}

type PayOuCardWithdrawRsp struct {
	RequestID  string                 `json:"requestId"`
	MerchantID string                 `json:"merchantId"`
	Signature  string                 `json:"signature"`
	Success    bool                   `json:"success"`
	Message    string                 `json:"message"`
	Code       int                    `json:"code"`
	Data       *PayOuCardWithdrawData `json:"data"`
}

type PayOuCardWithdrawData struct {
	CardNo         string  `json:"cardNo"`
	OrderNo        string  `json:"orderNo"`
	Currency       string  `json:"currency"`
	Status         int     `json:"status"`
	Msg            string  `json:"msg"`
	RechargeAmount float64 `json:"rechargeAmount"`
	ReceivedAmount float64 `json:"receivedAmount"`
	Fee            float64 `json:"fee"`
}
