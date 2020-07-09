/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package generate

var (
	defaultOptions = Options{
		CustomizeShortCode: "",
	}
)

type Options struct {
	CustomizeShortCode string
}

func withCustomizeShortCode(customizeShortCode string) Option {
	return func(o *Options) {
		o.CustomizeShortCode = customizeShortCode
	}
}
