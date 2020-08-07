package dtrobot

import (
	"fmt"
	"github.com/sockstack/dtrobot/message"
	"testing"
)

func TestNewRobot(t *testing.T) {
	accessToken := "79a01c796146d462b59bd6befc8d43e2c87dc446218d8757acca10d752c4fa03"
	robot, err := NewRobot(accessToken, WithSecret("SEC5a7d0e259fd08f1f19573101617713dcb19e5733b69f158969beaa723648410d"))

	if err != nil {
		t.Error(err)
	}

	textMessage := message.NewTextMessage(message.WithText(message.Text{Content: "ok"}))
	send, err := robot.Send(textMessage)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(send)
}
