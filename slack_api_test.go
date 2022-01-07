package slack_logger

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestAPI(t *testing.T) {
	// t.Logf("start!")

	Init(&ClientConfig{
		ApiRoot: os.Getenv("SLACK_API_ROOT"),
		Channel: os.Getenv("SLACK_CHANNEL"),
		OAuthToken: os.Getenv("SLACK_OAUTH_TOKEN"),
	})
	Message("test api message 1", nil)
	Message("test api message 2", nil)
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
	// log.SetFormatter(&log.TextFormatter{})
}
