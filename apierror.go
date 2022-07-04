package discordwebhook

import "encoding/json"

// APIError is the Discord error response format based on information found at
// https://discord.com/developers/docs/reference#error-messages
type APIError struct {
	// Code is a unique identifier for the type of message returned. For further
	// reference visit the documentation at
	// https://discord.com/developers/docs/topics/opcodes-and-status-codes#json
	Code int `json:"code"`
	// Errors is a map of details about the violation. Not yet implemented.
	Errors map[string]*json.RawMessage `json:"errors,omitempty"`
	// Message is the short description of the error
	Message string `json:"message"`
}
