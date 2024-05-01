# chat
Chat is a cli application that allows you to send messages to google and slack.

## Install
Download binaries from github release and then install.

## Setup
- Create config file at home directory. `.chat.ini`
```ini
[default]
PROVIDER=slack
SLACK_TOKEN=
CHANNEL_ID=

[slack]
PROVIDER=slack
SLACK_TOKEN=
CHANNEL_ID=

[google]
PROVIDER=google
WEB_HOOK_URL=
```

## Usage
- Send message using default provider
```sh
chat "hello"
```

- Reply message using default provider
```sh
chat -t="<thread-id>" "Hello"
```

- Send message using slack provider
```sh
chat --provider=slack "hello"
```

- Reply message using slack provider
```sh
chat --provider=slack -t="<thread-id>" "Hello"
```

- Send message using google provider
```sh
chat --provider=google "hello"
```

- Reply message using google provider
```sh
chat --provider=google -t="<thread-id>" "Hello"
```

- See help
```sh
chat --help
```
