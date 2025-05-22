package utils

import (
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"net/url"
	"sort"
	"time"
)

func Sign(paramMap map[string]interface{}, accessKey string) string {

	keys := lo.Keys(paramMap)

	// 1. 参数名按照ASCII码表升序排序
	sort.Strings(keys)

	//拼接签名原始字符串
	//rawString := ""
	values := url.Values{}
	lo.ForEach(keys, func(x string, index int) {
		value := ""
		if x == "data" {
			//data字段(是一个map.先转成string)
			valueByte, _ := json.Marshal(paramMap[x])
			value = string(valueByte)
		} else {
			value = cast.ToString(paramMap[x])
		}
		values.Add(x, value)
	})
	queryString := values.Encode()

	fmt.Printf("[rawString]%s\n", queryString)

	//4. 计算RSA签名
	signResult, _ := SignSHA256([]byte(queryString), accessKey)
	return signResult
}

func VerifySign(paramMap map[string]interface{}, publicKey string, sign string) bool {

	keys := lo.Keys(paramMap)

	// 1. 参数名按照ASCII码表升序排序
	sort.Strings(keys)

	//拼接签名原始字符串
	//rawString := ""
	values := url.Values{}
	lo.ForEach(keys, func(x string, index int) {
		value := ""
		if x == "data" {
			//data字段(是一个map.先转成string)
			valueByte, _ := json.Marshal(paramMap[x])
			value = string(valueByte)
		} else {
			value = cast.ToString(paramMap[x])
		}
		values.Add(x, value)
	})
	queryString := values.Encode()

	fmt.Printf("[rawString]%s\n", queryString)

	//4. 验证
	verifyResult, _ := Verify([]byte(queryString), publicKey, sign)
	return verifyResult
}

//--------------------------------------------------

// GenRequestID generates request ID with prefix
func GenRequestID(prefix string) string {
	if prefix == "" {
		prefix = "PYC"
	}
	return prefix + time.Now().Format("20060102150405.000")
}
