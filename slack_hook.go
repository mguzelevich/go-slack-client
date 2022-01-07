package slack_logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

/*
POST https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX
Content-type: application/json
{
    "text": "Hello, world."
  "thread_ts": "PARENT_MESSAGE_TS",
}

curl -X POST
	--data-urlencode "payload={\"channel\": \"#$(slack_channel)\", \"username\": \"postman\", \"text\": \"$(msg)\", \"icon_emoji\": \":ghost:\"}"
	$(slack_hook)

*/
func SendHook(message string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", config.HookRoot, config.HookUrl)

	req := slackReq{
		Channel:   config.Channel,
		Username:  config.BotUsername,
		Text:      message,
		IconEmoji: config.BotIcon,
	}

	reqStr, _ := json.Marshal(req)
	l := log.WithFields(log.Fields{
		"request": string(reqStr),
	})

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(req)

	request, err := http.NewRequest("POST", url, b)
	request.Close = true
	// request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.UserOauthToken))
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	if resp == nil {
		resp = &http.Response{}
	}

	body, _ := DecodeBody(resp)

	l = l.WithFields(log.Fields{
		"response":      resp,
		"response.body": body,
	})

	if err != nil || resp.StatusCode != 200 {
		l.WithFields(log.Fields{"error": err}).Errorf("request failed")
	} else {
		l.Trace("request done")
	}

	return resp, err
}
