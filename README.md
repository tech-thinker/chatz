# chatz
Chatz is a cli application that allows you to send messages to google and slack.

## Installation
Download and install executable binary from GitHub releases page.

### Linux Installation
```sh
curl -sL https://github.com/tech-thinker/chatz/releases/download/v1.0.0/chatz-linux-amd64 -o chatz
chmod +x chatz
sudo mv chatz /usr/bin
```

### MacOS Installation
```sh
curl -sL https://github.com/tech-thinker/chatz/releases/download/v1.0.0/chatz-darwin-amd64 -o chatz
chmod +x chatz
sudo mv chatz /usr/bin
```

### Windows Installation
```sh
curl -sL https://github.com/tech-thinker/chatz/releases/download/v1.0.0/chatz-windows-amd64.exe -o chatz.exe
chatz.exe
```

## Setup
- Create config file at home directory. `.chatz.ini`
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
chatz "hello"
```

- Reply message using default provider
```sh
chatz -t="<thread-id>" "Hello"
```

- Send message using slack provider
```sh
chatz --provider=slack "hello"
```

- Reply message using slack provider
```sh
chatz --provider=slack -t="<thread-id>" "Hello"
```

- Send message using google provider
```sh
chatz --provider=google "hello"
```

- Reply message using google provider
```sh
chatz --provider=google -t="<thread-id>" "Hello"
```

- See output
```sh
chatz -o "Hi"
```

- See help
```sh
chatz --help
```
