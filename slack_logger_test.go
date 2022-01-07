package slack_logger

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	// t.Logf("start!")

	Init(&ClientConfig{
		ApiRoot:    os.Getenv("SLACK_API_ROOT"),
		Channel:    os.Getenv("SLACK_CHANNEL"),
		OAuthToken: os.Getenv("SLACK_OAUTH_TOKEN"),
	})

	logger := log.New()
	logger = Logger()

	logger.WithFields(log.Fields{"message": "1"}).Info("test log message")
	logger.WithFields(log.Fields{"message": "2"}).Info("test log message")
	logger.WithFields(log.Fields{"message": "3", "_new_thread": ""}).Info("test log message")
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
	// log.SetFormatter(&log.TextFormatter{})
}
