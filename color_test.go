package discordwebhook_test

import (
	"errors"
	"testing"

	discord "github.com/SebastiaanPasterkamp/go-discordwebhook"
)

func TestColorToString(t *testing.T) {
	testCases := []struct {
		name     string
		color    discord.Color
		expected string
	}{
		{"Black", 0, "#000000"},
		{"White", 16777215, "#FFFFFF"},
		{"Green", 3066993, "#2ECC71"},
		{"Blue", 3447003, "#3498DB"},
		{"Dark Purple", 7419530, "#71368A"},
		{"Orange", 15105570, "#E67E22"},
		{"Red", 15158332, "#E74C3C"},
		{"Grey", 9807270, "#95A5A6"},
		{"Yellow", 16776960, "#FFFF00"},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			result := tt.color.String()

			if result != tt.expected {
				t.Errorf("Wrong color. Expected %q, got %q.",
					tt.expected, result)
			}
		})
	}
}

func TestColorFromString(t *testing.T) {
	testCases := []struct {
		name          string
		hex           string
		expectedColor discord.Color
		expectedError error
	}{
		{"Black", "#000000", 0, nil},
		{"White", "FFFFFF", 16777215, nil},
		{"Green", "#2ecc71", 3066993, nil},
		{"Blue", "3498db", 3447003, nil},
		{"Dark Purple", "71368A", 7419530, nil},
		{"Orange", "#e67e22", 15105570, nil},
		{"Red", "E74C3C", 15158332, nil},
		{"Grey", "#95A5A6", 9807270, nil},
		{"Yellow", "FFFF00", 16776960, nil},
		{"Too short", "fffff", 0, discord.ErrColorTooShort},
		{"Too long", "fffffff", 0, discord.ErrColorTooLong},
		{"Malformed", "foobar", 0, discord.ErrColorMalformed},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			result, err := discord.ColorFromString(tt.hex)
			if !errors.Is(err, tt.expectedError) {
				t.Errorf("Unexpected error. Expected %d, got %d.",
					tt.expectedError, err)
			}

			if result != tt.expectedColor {
				t.Errorf("Wrong color. Expected %d, got %d.",
					tt.expectedColor, result)
			}
		})
	}
}
