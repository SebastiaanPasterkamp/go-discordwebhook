package discordwebhook

import "time"

// Embed is an object describing embedded content.This follows the documentation
// found at
// https://discord.com/developers/docs/resources/channel#embed-object
//
// For sake of simplicity this structure does not support the additional embed
// structures like footer, image, thumbnail, video, provider, author, and
// fields yet.
type Embed struct {
	// Title is the title of embed
	Title string `json:"title,omitempty"`
	// Type is the type of embed (always "rich" for webhook embeds)
	Type EmbedType `json:"type,omitempty"`
	// Description is the description of embed
	Description string `json:"description,omitempty"`
	// URL is the original url of embed
	URL string `json:"url,omitempty"`
	// Timestamp is ISO8601 timestamp of the of embed content
	Timestamp time.Time `json:"timestamp,omitempty"`
	// Color is integer color code of the embed
	Color Color `json:"color,omitempty"`
}

// EmbedType is a string type for an embed enum
type EmbedType string

const (
	// EmbedRich is a embed type for an generic embed rendered from embed
	// attributes
	EmbedRich EmbedType = "rich"
	// EmbedImage is a embed type for an embedded image
	EmbedImage EmbedType = "image"
	// EmbedVideo is a embed type for an embedded video
	EmbedVideo EmbedType = "video"
	// EmbedGifv is a embed type for an embedded animated gif image rendered as
	// a video embed
	EmbedGifv EmbedType = "gifv"
	// EmbedArticle is a embed type for an embedded article
	EmbedArticle EmbedType = "article"
	// EmbedLink is a embed type for an embedded link
	EmbedLink EmbedType = "link"
)
