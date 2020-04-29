package message

type optionHandle func(o *textOptions)

type Text struct {
	Content string `json:"content"`
}

type Mobiles struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll bool `json:"isAtAll"`
}

type textOptions struct {
	MsgType	string `json:"msgtype"`
	Text Text `json:"text"`
	At Mobiles
}

func WithMobiles(mobiles Mobiles) optionHandle {
	return func(o *textOptions) {
		o.At = mobiles
	}
}

func WithText(text Text) optionHandle {
	return func(o *textOptions) {
		o.Text = text
	}
}
