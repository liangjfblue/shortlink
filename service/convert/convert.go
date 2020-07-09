/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package convert

type IConvert interface {
	//任意进制转换为短码
	Encode(interface{}) string
	//短码转换为任意进制
	Decode(string) interface{}
}

var (
	defaultConvert = newBase62Decimal()
	_convertM      map[string]IConvert
)

func init() {
	RegisterConvert("default", defaultConvert)
}

func RegisterConvert(name string, convert IConvert) {
	if _convertM == nil {
		_convertM = make(map[string]IConvert)
	}
	_convertM[name] = convert
}

func GetConvert(name string) IConvert {
	if _, ok := _convertM[name]; ok {
		return _convertM[name]
	}
	return defaultConvert
}

//SetConvert 设置默认convert
func SetDefaultConvert(name string) bool {
	if _, ok := _convertM[name]; ok {
		defaultConvert = _convertM[name]
		return true
	}
	return false
}

//Encode 任意进制转换为短码
func Encode(id interface{}) string {
	return defaultConvert.Encode(id)
}

//Decode 短码转换为任意进制
func Decode(code string) interface{} {
	return defaultConvert.Decode(code)
}
