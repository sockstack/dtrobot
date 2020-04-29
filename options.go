package dtrobot

type optionHandle func(o *options)

type options struct {
	Host string
	Timestamp int64
	Sign string
	Secret string
	AccessToken string
	Path string
}

//WithSecret 设置secret
func WithSecret(secret string) optionHandle {
	return func(o *options) {
		o.Secret = secret
	}
}
