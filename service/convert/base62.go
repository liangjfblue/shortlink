/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package convert

import (
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
	var (
		pos    int
		number int64
	)
	sum := len(str62)
	for i := 0; i < sum; i++ {
		pos = strings.IndexAny(CODE62, str62[i:i+1])
		number = int64(math.Pow(62, float64(sum-i-1))*float64(pos)) + number
	}
	return number
}
