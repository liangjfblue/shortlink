/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package core

type ICore interface {
	GetShortLinkByLongLink(string) (string, error)
	GetLongLinkByShortLink(string) (string, error)
	CreateShortLinkByCustomizeShortCode(string, string) (string, error)
}

var (
	defaultCore = newRedisWay()
	_coreM      map[string]ICore
)

func init() {
	RegisterShortlink("default", defaultCore)
}

func RegisterShortlink(name string, Shortlink ICore) {
	if _coreM == nil {
		_coreM = make(map[string]ICore)
	}
	_coreM[name] = Shortlink
}

//GetCore 获取某种处理类型
func GetCore(name string) ICore {
	if _, ok := _coreM[name]; ok {
		return _coreM[name]
	}
	return defaultCore
}

//SetCore 设置默认core
func SetCore(name string) bool {
	if _, ok := _coreM[name]; ok {
		defaultCore = _coreM[name]
		return true
	}
	return false
}

//GetShortLinkByLongLink 长连接转短连接
func GetShortLinkByLongLink(longLink string) (string, error) {
	return defaultCore.GetShortLinkByLongLink(longLink)
}

//GetLongLinkByShortLink 短连接转长连接
func GetLongLinkByShortLink(shortLink string) (string, error) {
	return defaultCore.GetLongLinkByShortLink(shortLink)
}

//CreateShortLinkByCustomizeShortCode 自定义短连接转长连接
func CreateShortLinkByCustomizeShortCode(shortCode, longLink string) (string, error) {
	return defaultCore.CreateShortLinkByCustomizeShortCode(shortCode, longLink)
}
