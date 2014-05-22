package notifications

import (
	"log"

	"github.com/andybons/hipchat"
)

type HipChatNotifier struct {
	AuthToken string
	RoomID    string
	From      string
}

func NewHipChatNotifier(authToken, roomID, from string) *HipChatNotifier {
	return &HipChatNotifier{
		AuthToken: authToken,
		RoomID:    roomID,
		From:      from,
	}
}

func (hcn *HipChatNotifier) Notify(message string) error {
	if hcn.AuthToken == "" || hcn.RoomID == "" || hcn.From == "" {
		log.Printf("missing required config bits: %#v\n", hcn)
		return nil
	}

	client := hipchat.Client{AuthToken: hcn.AuthToken}
	req := hipchat.MessageRequest{
		RoomId:        hcn.RoomID,
		From:          hcn.From,
		Message:       message,
		MessageFormat: hipchat.FormatText,
	}

	return client.PostMessage(req)
}
