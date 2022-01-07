package slack_logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func NewSlackThread(message string) (*slackThread, error) {
	// if channel == "" {
	// 	return nil, nil
	// }
	thread := &slackThread{
		// channel: config.Channel,
	}
	thread.Send(message)
	return thread, nil
}

func (s *slackThread) Send(message string) (*http.Response, error) {
	resp, err := Message(message, s.rootTs)
	if resp == nil {
		resp = &http.Response{}
	}

	body, _ := DecodeBody(resp)
	if s.rootTs == nil {
		ts := fmt.Sprintf("%s", body["ts"])
		s.rootTs = &ts
	}

	return resp, err
}

func DecodeBody(resp *http.Response) (map[string]interface{}, error) {
	body := make(map[string]interface{})
	if resp.Body != nil {
		b, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		json.Unmarshal(b, &body)
	}
	return body, nil
}

func Message(message string, rootTs *string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", config.ApiRoot, "chat.postMessage")

	req := slackReq{
		Channel:  config.Channel,
		Text:     message,
		ThreadTs: rootTs,
	}

	reqStr, _ := json.Marshal(req)
	l := log.WithFields(log.Fields{
		"request": string(reqStr),
	})

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(req)

	request, err := http.NewRequest("POST", url, b)
	request.Close = true
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.OAuthToken))
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
