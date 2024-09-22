package models

// SlackRes represents the main structure of the JSON response.
type SlackRes struct {
	Ok               bool             `json:"ok"`
	Channel          string           `json:"channel"`
	Ts               string           `json:"ts"`
	Message          SlackMessage          `json:"message"`
	Warning          string           `json:"warning"`
	ResponseMetadata SlackResponseMetadata `json:"response_metadata"`
}

// SlackMessage represents the message object within the main response.
type SlackMessage struct {
	User       string     `json:"user"`
	Type       string     `json:"type"`
	Ts         string     `json:"ts"`
	BotID      string     `json:"bot_id"`
	AppID      string     `json:"app_id"`
	Text       string     `json:"text"`
	Team       string     `json:"team"`
	BotProfile SlackBotProfile `json:"bot_profile"`
	Blocks     []SlackBlock    `json:"blocks"`
}

// SlackBotProfile represents the profile of the bot that sent the message.
type SlackBotProfile struct {
	ID      string `json:"id"`
	AppID   string `json:"app_id"`
	Name    string `json:"name"`
	Icons   SlackIcons  `json:"icons"`
	Deleted bool   `json:"deleted"`
	Updated int64  `json:"updated"`
	TeamID  string `json:"team_id"`
}

// SlackIcons represents the bot's profile images of different sizes.
type SlackIcons struct {
	Image36 string `json:"image_36"`
	Image48 string `json:"image_48"`
	Image72 string `json:"image_72"`
}

// Block represents the block object in the message.
type SlackBlock struct {
	Type     string    `json:"type"`
	BlockID  string    `json:"block_id"`
	Elements []SlackElement `json:"elements"`
}

// SlackElement represents a section within the block.
type SlackElement struct {
	Type     string        `json:"type"`
	Elements []SlackInnerElement `json:"elements"`
}

// SlackInnerElement represents the inner elements within a rich text section.
type SlackInnerElement struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// SlackResponseMetadata represents additional metadata about the response.
type SlackResponseMetadata struct {
	Warnings []string `json:"warnings"`
}
