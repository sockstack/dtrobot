package dtrobot

import (
	"cn.sockstack/dtrobot/message"
	"cn.sockstack/dtrobot/pkg/encode"
	"fmt"
	"time"
)

var r *robot

type robot struct {
	Options *options
}

func (r *robot) Send(message message.Message) (Response, error) {
	data, err := message.Get()
	if err != nil {
		return nil, err
	}

	url := ""
	if r.Options.Secret != "" {
		url = fmt.Sprintf("%s%s?access_token=%s&timestamp=%d&sign=%s", r.Options.Host, r.Options.Path,
			r.Options.AccessToken, r.Options.Timestamp, r.Options.Sign)
	} else {
		url = fmt.Sprintf("%s%s?access_token=%s", r.Options.Host, r.Options.Path, r.Options.AccessToken)
	}

	resp, err := Post(url, data, WithHeader("Content-Type", "application/json"))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NewRobot(accessToken string, handles... optionHandle) (*robot, error) {
	o := &options{
		Host:        "https://oapi.dingtalk.com",
		Path:		 "/robot/send",
		AccessToken: accessToken,
		Timestamp: time.Now().Unix() * 1000,
	}
	for _, handle := range handles {
		handle(o)
	}

	if o.Secret != "" {
		sign := encode.CalcSign(o.Timestamp, o.Secret)
		o.Sign = sign
	}

	if r == nil {
		r = &robot{
			Options: o,
		}
	}

	return r, nil
}
