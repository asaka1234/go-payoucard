package utils

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"sort"
	"strings"
)

func Sign(paramMap map[string]interface{}, privateKey string) string {

	keys := lo.Keys(paramMap)
	// 1. 参数名按照ASCII码表升序排序
	sort.Strings(keys)

	//拼接签名原始字符串
	//rawString := ""
	var pairs []string
	lo.ForEach(keys, func(x string, index int) {
		value := ""
		if x == "data" {
			//data字段(是一个map.先转成string)
			valueByte, _ := json.Marshal(paramMap[x])
			value = string(valueByte)
		} else {
			value = cast.ToString(paramMap[x])
		}

		if value != "" {
			pairs = append(pairs, x+"="+value)
		}
	})
	queryString := strings.Join(pairs, "&") // values.Encode()

	fmt.Printf("[rawString]%s\n", queryString)

	//4. 计算RSA签名
	signResult, err := SignSHA256RSA([]byte(queryString), privateKey)
	if err != nil {
		fmt.Printf("==sign===>%s\n", err.Error())
	}
	return signResult
}

func VerifySign(paramMap map[string]interface{}, publicKey string, sign string) bool {

	keys := lo.Keys(paramMap)

	// 1. 参数名按照ASCII码表升序排序
	sort.Strings(keys)

	//拼接签名原始字符串
	//rawString := ""
	var pairs []string
	lo.ForEach(keys, func(x string, index int) {
		value := ""
		if x == "data" {
			//data字段(是一个map.先转成string)
			valueByte, _ := json.Marshal(paramMap[x])
			value = string(valueByte)
		} else {
			value = cast.ToString(paramMap[x])
		}
		if value != "" {
			pairs = append(pairs, x+"="+value)
		}
	})
	queryString := strings.Join(pairs, "&")

	fmt.Printf("[rawString]%s\n", queryString)

	//4. 验证
	verifyResult, err := VerifySHA256RSA([]byte(queryString), publicKey, sign)
	if err != nil {
		return false
	}

	return verifyResult
}

//--------------------------------------------------

// GenRequestID generates request ID
func GenRequestID() string {
	return strings.Join(strings.Split(uuid.NewString(), "-"), "")[0:20]
	//return prefix + time.Now().Format("20060102150405.000")
}
