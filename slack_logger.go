package slack_logger

import (
	"bytes"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	levelPrefixes = map[log.Level]string{
		log.PanicLevel: ":exclamation:",
		log.FatalLevel: ":no_entry:",
		log.ErrorLevel: ":sos:",
		log.WarnLevel:  ":warning:",
		log.InfoLevel:  ":information_source:",
		log.DebugLevel: ":spider:",
		log.TraceLevel: ":spider:",
	}
)

type slackLogger struct {
	channel  string
	threadTs *string `json:"thread_ts,omitempty"` //"thread_ts": "PARENT_MESSAGE_TS",
}

func (s *slackLogger) Levels() []log.Level {
	return log.AllLevels
}

func (s *slackLogger) Fire(entry *log.Entry) error {
	message := bytes.Buffer{}
	message.WriteString(fmt.Sprintf("%v ", levelPrefixes[entry.Level]))
	message.WriteString(fmt.Sprintf("%v ", entry.Message))

	thread := s.threadTs
	if _, ok := entry.Data["_new_thread"]; ok {
		thread = nil
		delete(entry.Data, "_new_thread")
	}

	data := []string{}
	for k, v := range entry.Data {
		data = append(data, fmt.Sprintf("%s=%v", k, v))
	}
	if len(data) > 0 {
		message.WriteString(fmt.Sprintf("`%v`", strings.Join(data, " ")))
	}
	resp, err := Message(message.String(), thread)
	if err != nil {
		log.WithFields(log.Fields{"resp": resp, "err": err}).Error("slack.Message")
	}
	if thread == nil {
		body, _ := DecodeBody(resp)
		ts := fmt.Sprintf("%s", body["ts"])
		s.threadTs = &ts
	}

	return nil
}

func Logger() *log.config.Logger {
	if config.Logger == nil {
		config.Logger = log.New()
	}
	logger := log.New()
	hook := &slackLogger{
		channel: config.Channel,
	}
	logger.AddHook(hook)
	return logger
}
