package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
)

const (
	Success       = "success"
	Code          = "code"
	Message       = "message"
	Data          = "data"
	Signature     = "signature"
	RequestID     = "requestId"
	MerchantID    = "merchantId"
	NotifyType    = "notifyType"
	RequestPrefix = "PYC"
)

type PayOuCardUtil struct {
	Logger Logger
}

// InitCommonRequest initializes request with signature
func (ut *PayOuCardUtil) InitCommonRequest(data interface{}, merchantId, privateKey string) (string, error) {
	commonParamMap, err := ut.buildCommonParam(nil, nil, nil, data, nil, merchantId)
	if err != nil {
		return "", fmt.Errorf("build common params failed: %v", err)
	}

	signDataStr, err := ut.getTreeValue(commonParamMap)
	if err != nil {
		return "", fmt.Errorf("get tree value failed: %v", err)
	}
	ut.Logger.Infof("signDataStr: %s", signDataStr)

	// Sign the data
	signature, err := SignSHA256([]byte(signDataStr), privateKey)
	if err != nil {
		return "", fmt.Errorf("signature failed: %v", err)
	}

	commonParamMap[Signature] = signature
	commonParamStr, err := json.Marshal(commonParamMap)
	if err != nil {
		return "", fmt.Errorf("marshal json failed: %v", err)
	}

	ut.Logger.Infof("commonRequest: %s", commonParamStr)
	return string(commonParamStr), nil
}

// VerifySignature verifies the response signature
func (ut *PayOuCardUtil) VerifySignature(responseStr, publicKey string) (bool, error) {
	var requestMap map[string]interface{}
	if err := json.Unmarshal([]byte(responseStr), &requestMap); err != nil {
		return false, fmt.Errorf("unmarshal response failed: %v", err)
	}

	signature, ok := requestMap[Signature].(string)
	if !ok {
		return false, fmt.Errorf("signature not found or invalid")
	}

	delete(requestMap, Signature)

	signDataStr, err := ut.getTreeValue(requestMap)
	if err != nil {
		return false, fmt.Errorf("get tree value failed: %v", err)
	}
	ut.Logger.Infof("verify signDataStr: %s", signDataStr)

	return Verify([]byte(signDataStr), publicKey, signature)
}

// GenRequestID generates request ID with prefix
func (ut *PayOuCardUtil) GenRequestID(prefix string) string {
	if prefix == "" {
		prefix = RequestPrefix
	}
	return prefix + time.Now().Format("20060102150405.000")
}

// getTreeValue constructs signature string from sorted params
func (ut *PayOuCardUtil) getTreeValue(paramMap map[string]interface{}) (string, error) {
	// Sort keys
	keys := make([]string, 0, len(paramMap))
	for k := range paramMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var origin bytes.Buffer
	first := true

	for _, key := range keys {
		value := paramMap[key]
		if value == nil {
			continue
		}

		valueStr := fmt.Sprintf("%v", value)
		if valueStr == "" {
			continue
		}

		if !first {
			origin.WriteString("&")
		} else {
			first = false
		}

		origin.WriteString(fmt.Sprintf("%s=%s", key, valueStr))
	}

	return origin.String(), nil
}

func (ut *PayOuCardUtil) buildCommonParam(code, success, message interface{}, data interface{},
	notifyType *int, merchantId string) (map[string]interface{}, error) {

	commonParamMap := make(map[string]interface{})
	commonParamMap[RequestID] = ut.GenRequestID("")
	commonParamMap[MerchantID] = merchantId
	commonParamMap[NotifyType] = notifyType
	commonParamMap[Code] = code
	commonParamMap[Success] = success
	commonParamMap[Message] = message

	if data != nil {
		switch v := data.(type) {
		case string:
			if strings.TrimSpace(v) != "" {
				var jsonData interface{}
				if err := json.Unmarshal([]byte(v), &jsonData); err == nil {
					commonParamMap[Data] = jsonData
				} else {
					commonParamMap[Data] = v
				}
			}
		default:
			commonParamMap[Data] = v
		}
	}

	return commonParamMap, nil
}
