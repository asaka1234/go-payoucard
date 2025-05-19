package go_payoucard

import (
	"bytes"
	"encoding/json"
	"github.com/asaka1234/go-payoucard/utils"
	"github.com/spf13/cast"
	"log"
	"net/http"
)

// 提现
func (cli *Client) Withdraw(req PayOuCardWithdrawReq) (*PayOuCardWithdrawRsp, error) {
	// Prepare request data
	requestData := map[string]interface{}{
		"uniqueId": req.Data.UniqueID,
		"cardNo":   req.Data.CardNo,
		"currency": req.Data.Currency,
		"amount":   cast.ToString(req.Data.Amount),
		"orderNo":  req.Data.OrderNo,
	}

	// Generate signed request
	pyCard := utils.PayOuCardUtil{
		Logger: cli.logger,
	}
	jsonStr, err := pyCard.InitCommonRequest(requestData, req.MerchantID, cli.AccessKey)
	if err != nil {
		log.Printf("PayOuCardService#signature#error: %v", err)
		return nil, err
	}

	log.Printf("PayOuCardService#withdraw#json: %s", jsonStr)

	// Create HTTP request
	httpReq, err := http.NewRequest("POST", cli.WithdrawURL, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	var result PayOuCardWithdrawRsp
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
