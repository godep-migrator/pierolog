package notifications

import (
	"github.com/Sirupsen/logrus"
	"github.com/andybons/hipchat"
)

type HipChatNotifier struct {
	AuthToken string
	RoomID    string
	From      string

	log *logrus.Logger
}

func NewHipChatNotifier(authToken, roomID, from string, log *logrus.Logger) *HipChatNotifier {
	return &HipChatNotifier{
		AuthToken: authToken,
		RoomID:    roomID,
		From:      from,

		log: log,
	}
}

func (hcn *HipChatNotifier) Notify(message string) error {
	if hcn.AuthToken == "" || hcn.RoomID == "" || hcn.From == "" {
		hcn.log.WithFields(logrus.Fields{
			"auth_token": hcn.AuthToken,
			"room_id":    hcn.RoomID,
			"from":       hcn.From,
		}).Warn("missing required config bits")
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
