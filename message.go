package discordwebhook

// Message is the root structure for a Discord Webhook API payload. This follows
// the documentation found at
// https://discord.com/developers/docs/resources/webhook#execute-webhook-jsonform-params
//
// This format does not support file uploads or message components.
//
// To do: Support file uploads following the instructions found at
// https://discord.com/developers/docs/reference#uploading-files
//
// To do: Support message components following the instructions found at
// https://discord.com/developers/docs/interactions/message-components
type Message struct {
	// Content is the message contents (up to 2000 characters). Either content
	// or embeds is required.
	Content string `json:"content,omitempty"`
	// Username can be used to override the default username of the webhook
	Username string `json:"username,omitempty"`
	// AvatarURL can be used to override the default avatar of the webhook
	AvatarURL string `json:"avatar_url,omitempty"`
	// TTS can be set to true if this is a TTS message. Defaults to false.
	TTS bool `json:"tts,omitempty"`
	// Embeds is an array of up to 10 embed objects	using the rich content
	// format. One of content, file, or embeds is required.
	Embeds []Embed `json:"embeds,omitempty"`
	// AllowedMentions restricts which mentions and/or types of mentions are
	// used and allowed.
	AllowedMentions AllowedMention `json:"allowed_mentions,omitempty"`
	// ThreadName is the name of thread to create (requires the webhook channel
	// to be a forum channel)
	ThreadName string `json:"thread_name,omitempty"`
}

// AllowedMention is an object describing which types of mentions are permitted
// in the Message. This follows the documentation found at
// https://discord.com/developers/docs/resources/channel#allowed-mentions-object
type AllowedMention struct {
	// Parse is a array of allowed mention types to parse from the content.
	Parse []MentionType `json:"parse,omitempty"`
	// Roles is a list of snowflake role IDs to mention (Max size of 100)
	Roles []Snowflake `json:"roles,omitempty"`
	// Users is a list of Snowflake user IDs to mention (Max size of 100)
	Users []Snowflake `json:"users,omitempty"`
	// RepliedUser is a boolean	indicating whether to mention the author of the
	// message being replied to (default false)
	RepliedUser bool `json:"replied_user,omitempty"`
}

// Snowflake is an int64 type for unique numeric identifiers used by Discord.
type Snowflake int64

// MentionType is a string type for a mention enum
type MentionType string

const (
	// MentionRole permits mentions of specific roles
	MentionRole MentionType = "roles"
	// MentionUser permits mentions of specific users
	MentionUser MentionType = "users"
	// MentionEveryone permits mentions of @everyone or @here
	MentionEveryone MentionType = "everyone"
)
