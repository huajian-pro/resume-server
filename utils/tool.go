package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/nilorg/sdk/convert"
)

// InterfaceToString 类型转换工具，去吧：字符串
func InterfaceToString(src interface{}) string {
	if src == nil {
		fmt.Println("src为空")
	}
	switch src.(type) {
	case string:
		return src.(string)
	case int, int8, int32, int64:
	case uint8, uint16, uint32, uint64:
	case float32, float64:
		return convert.ToString(src)
	}
	data, err := json.Marshal(src)
	if err != nil {
		panic(any(err))
	}
	return string(data)
}

// GetRandomCode 随机生成 6 位数验证码
func GetRandomCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return code
}
