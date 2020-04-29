package message

import "encoding/json"

type textMessage struct {
	Options *textOptions
}

func (m *textMessage) Get() ([]byte, error) {
	bytes, err := json.Marshal(m.Options)
	return bytes, err
}

func NewTextMessage(handles... optionHandle) *textMessage {
	options := &textOptions{
		MsgType: "text",
		Text:    Text{Content: ""},
		At:      Mobiles{
			IsAtAll: false,
			AtMobiles: []string{},
		},
	}
	for _, i := range handles {
		i(options)
	}

	return &textMessage{
		Options: options,
	}
}
