package slack_logger

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestInit(t *testing.T) {
	// t.Logf("start!")

	Init(&ClientConfig{
		HookRoot:    os.Getenv("SLACK_HOOK_ROOT"),
		HookUrl:     os.Getenv("SLACK_HOOK_URL"),
		Channel:     os.Getenv("SLACK_CHANNEL"),
		BotUsername: os.Getenv("SLACK_BOT_USERNAME"),
		BotIcon:     os.Getenv("SLACK_BOT_ICON"),
	})
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
	// log.SetFormatter(&log.TextFormatter{})
}
