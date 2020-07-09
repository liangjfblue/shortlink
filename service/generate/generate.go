/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package generate

type IGenerate interface {
	//生成短码id
	Create(opts ...Option) (int64, error)
}

type Option func(*Options)

var (
	defaultGenerate = NewRedisGenerate()
	_generateM      map[string]IGenerate
)

func RegisterGenerate(name string, generate IGenerate) {
	if _generateM == nil {
		_generateM = make(map[string]IGenerate)
	}
	_generateM[name] = generate
}

func GetConvert(name string) IGenerate {
	if _, ok := _generateM[name]; ok {
		return _generateM[name]
	}
	return defaultGenerate
}

//SetConvert 设置默认 generate
func SetDefaultGenerate(name string) bool {
	if _, ok := _generateM[name]; ok {
		defaultGenerate = _generateM[name]
		return true
	}
	return false
}

//Create 生成短码id
func Create(opts ...Option) (int64, error) {
	return defaultGenerate.Create(opts...)
}
