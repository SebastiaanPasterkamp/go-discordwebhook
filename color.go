package discordwebhook

import (
	"fmt"
	"strconv"
)

// Color is a 16 bit unsigned integer used by Discord to describe a color. This
// directly maps to the common hexadecimal RGB color representation where #
// 000000 is black and #ffffff is white. The color type comes with conversion
// functions between the integer and hexadecimal format.
type Color uint32

var (
	// ErrColorTooShort is the error returned when the hexadecimal string to
	// convert to a Color does not contain 6 hexadecimal characters.
	ErrColorTooShort = fmt.Errorf("the hex string is shorter than 6 characters")
	// ErrColorTooLong is the error returned when the hexadecimal string to
	// convert to a Color is more than 6 hexadecimal characters, even after
	// removing the optional # prefix..
	ErrColorTooLong = fmt.Errorf("the hex string is longer than 6 characters")
	// ErrColorMalformed is the error returned when the hexadecimal string to
	// convert to a Color contains other characters than the optional # prefix
	// followed by 0-9, a-f, or A-F.
	ErrColorMalformed = fmt.Errorf("the string does not look hexadecimal")
)

const (
	// ColorBlack is a predefined Discord color
	ColorBlack = Color(0)
	// ColorAqua is a predefined Discord color
	ColorAqua = Color(1752220)
	// ColorDarkAqua is a predefined Discord color
	ColorDarkAqua = Color(1146986)
	// ColorGreen is a predefined Discord color
	ColorGreen = Color(3066993)
	// ColorDarkGreen is a predefined Discord color
	ColorDarkGreen = Color(2067276)
	// ColorBlue is a predefined Discord color
	ColorBlue = Color(3447003)
	// ColorDarkBlue is a predefined Discord color
	ColorDarkBlue = Color(2123412)
	// ColorPurple is a predefined Discord color
	ColorPurple = Color(10181046)
	// ColorDarkPurple is a predefined Discord color
	ColorDarkPurple = Color(7419530)
	// ColorPink is a predefined Discord color
	ColorPink = Color(15277667)
	// ColorDarkPink is a predefined Discord color
	ColorDarkPink = Color(11342935)
	// ColorGold is a predefined Discord color
	ColorGold = Color(15844367)
	// ColorDarkGold is a predefined Discord color
	ColorDarkGold = Color(12745742)
	// ColorOrange is a predefined Discord color
	ColorOrange = Color(15105570)
	// ColorDarkOrange is a predefined Discord color
	ColorDarkOrange = Color(11027200)
	// ColorRed is a predefined Discord color
	ColorRed = Color(15158332)
	// ColorDarkRed is a predefined Discord color
	ColorDarkRed = Color(10038562)
	// ColorGrey is a predefined Discord color
	ColorGrey = Color(9807270)
	// ColorDarkGrey is a predefined Discord color
	ColorDarkGrey = Color(9936031)
	// ColorDarkerGrey is a predefined Discord color
	ColorDarkerGrey = Color(8359053)
	// ColorLightGrey is a predefined Discord color
	ColorLightGrey = Color(12370112)
	// ColorNavy is a predefined Discord color
	ColorNavy = Color(3426654)
	// ColorDarkNavy is a predefined Discord color
	ColorDarkNavy = Color(2899536)
	// ColorYellow is a predefined Discord color
	ColorYellow = Color(16776960)
)

// String converts a Color to a hexadecimal string
func (c Color) String() string {
	return fmt.Sprintf("#%06X", uint32(c))
}

// FromString updates the Color using the hexadecimal color value.
func (c *Color) FromString(hex string) error {
	if len(hex) < 6 {
		return fmt.Errorf("%w: %q", ErrColorTooShort, hex)
	}

	// Removing optional # prefix
	if hex[0] == '#' {
		hex = hex[1:]
	}

	if len(hex) > 6 {
		return fmt.Errorf("%w: %q", ErrColorTooLong, hex)
	}

	out, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrColorMalformed, err)
	}

	*c = Color(out)

	return nil
}

// ColorFromString creates a Color object using the hexadecimal color value.
func ColorFromString(hex string) (Color, error) {
	var c Color
	return c, c.FromString(hex)
}
