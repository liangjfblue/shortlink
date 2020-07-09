/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package convert

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type base62Decimal struct{}

const (
	CODE62     = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CodeLength = 62
)

func newBase62Decimal() IConvert {
	return &base62Decimal{}
}

//Encode 任意进制转换为短码
func (b *base62Decimal) Encode(num interface{}) string {
	number, _ := strconv.ParseUint(fmt.Sprint(num), 10, 64)

	if number == 0 {
		return "0"
	}
	result := make([]byte, 0)
	for number > 0 {
		round := number / CodeLength
		remain := number % CodeLength

		var tmp []byte
		tmp = append(tmp, CODE62[remain])
		result = append(tmp, result...)

		number = round
	}
	return string(result)
}

//Decode 短码转换为任意进制
func (b *base62Decimal) Decode(str62 string) interface{} {
	str62 = strings.TrimSpace(str62)
	var result = 0
	for index, char := range []byte(str62) {
		result += bytes.IndexByte([]byte(CODE62), char) * int(math.Pow(CodeLength, float64(index)))
	}
	return result
}
