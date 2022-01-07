package slack_logger

import (
	// "fmt"
	"net/http"
	log "github.com/sirupsen/logrus"
)

type ClientConfig struct {
	Logger *log.Logger

	ApiRoot  string
	HookRoot string
	HookUrl  string

	Channel string

	BotUsername string
	BotIcon     string

	OAuthToken string
}

var (
	client = &http.Client{}
	config *ClientConfig
)

func Init(cfg *ClientConfig) error {
	config = cfg
	return nil
}

type slackReq struct {
	Channel   string  `json:"channel"`
	Text      string  `json:"text"`
	Username  string  `json:"username,omitempty"`
	IconEmoji string  `json:"icon_emoji,omitempty"`
	ThreadTs  *string `json:"thread_ts,omitempty"` //"thread_ts": "PARENT_MESSAGE_TS",
}

type slackThread struct {
	channel string
	rootTs  *string
}
