# go-slack-client

minimal [Slack](https://slack.com/) integration

## slack api integration

```golang
    slack.Init(&slack.ClientConfig{
        ApiRoot: "https://slack.com/api/",
        Channel: "Uxxxxxxxx",
        OAuthToken: "xoxb-xxxxxxxxxxx-xxxxxxxxxxxxx-xxxxxxxxxxxxxxxxxxxxxxxx",
    })
    slack.Message("test api message", nil)
```

## slack hooks integration

simple

```golang
import (
    ...
    slack "github.com/mguzelevich/go-slack-client"
    ...
)

...
    slack.Init(&slack.ClientConfig{
        HookRoot:       "https://hooks.slack.com/services",
        HookUrl:        "/T0xxxxxxx/xxxxxxxxxxx/xxxxxxxxxxxxxxxxxxxxxxxx",
    })
    slack.SendHook("test")
```

advanced

```golang
import (
    ...
    slack "github.com/mguzelevich/go-slack-client"
    ...
)

...
    slack.Init(&slack.ClientConfig{
        HookRoot:       "https://hooks.slack.com/services",
        HookUrl:        "/T0xxxxxxx/xxxxxxxxxxx/xxxxxxxxxxxxxxxxxxxxxxxx",
        Channel:        "Uxxxxxxxx",
        BotUsername:    "slack-test",
        BotIcon:        ":ghost:",
    })
    slack.SendHook("test")
```

## logrus slack logging hook implementation

```golang
import (
    ...
    log "github.com/sirupsen/logrus"
    slack "github.com/mguzelevich/go-slack-client"
    ...
)

...
    slack.Init(&slack.ClientConfig{
        ApiRoot: "https://slack.com/api/",
        Channel: "Uxxxxxxxx",
        OAuthToken: "xoxb-xxxxxxxxxxx-xxxxxxxxxxxxx-xxxxxxxxxxxxxxxxxxxxxxxx",
    })

    logger := log.New()
    logger = slack.Logger()

    logger.WithFields(log.Fields{"message": "1"}).Info("test log message")
    logger.WithFields(log.Fields{"message": "2"}).Info("test log message")
    logger.WithFields(log.Fields{"message": "3", "_new_thread": ""}).Info("test log message")
}

```

## Makefile

```bash
.PHONY: run test clean cleanup prepare

SLACK_API_ROOT = https://slack.com/api/
SLACK_HOOK_ROOT = https://hooks.slack.com/services
SLACK_HOOK_URL = <...>
SLACK_CHANNEL = <...>
SLACK_BOT_USERNAME = slack-test
SLACK_BOT_ICON = :ghost:
SLACK_OAUTH_TOKEN = <...>

dummy:
    @echo "hello"

test:
    SLACK_API_ROOT=$(SLACK_API_ROOT) \
    SLACK_HOOK_ROOT=$(SLACK_HOOK_ROOT) \
    SLACK_HOOK_URL=$(SLACK_HOOK_URL) \
    SLACK_CHANNEL=$(SLACK_CHANNEL) \
    SLACK_BOT_USERNAME=$(SLACK_BOT_USERNAME) \
    SLACK_BOT_ICON=$(SLACK_BOT_ICON) \
    SLACK_OAUTH_TOKEN=$(SLACK_OAUTH_TOKEN) \
        go test -v ./...   
```
