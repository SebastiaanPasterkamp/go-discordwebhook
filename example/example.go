package main

import (
	"encoding/json"
	"log"
	"os"

	discord "github.com/SebastiaanPasterkamp/go-discordwebhook"
)

func main() {
	d, err := discord.NewFromEnv("DISCORD_URL")
	if err != nil {
		log.Fatalf("Failed to initialize discord webhook client: %v", err)
	}

	m := discord.Message{
		Content:   "This is the body",
		Username:  "example",
		AvatarURL: "https://avatars.githubusercontent.com/u/26205277?s=64&v=4",
		Embeds: []discord.Embed{
			{
				Title:       "discord webhook client",
				URL:         "https://github.com/SebastiaanPasterkamp/go-discordwebhook",
				Type:        discord.EmbedRich,
				Description: "Hello *discord*",
				Color:       discord.ColorBlue,
			},
		},
	}

	stdout := json.NewEncoder(os.Stdout)

	reply, fail, err := d.Send(m, true)
	if err != nil {
		stdout.Encode(fail)
		log.Fatalf("Failed to send message to discord: %v", err)
	}

	if fail != nil {
		stdout.Encode(fail)
		log.Println("Error response from discord")
		os.Exit(fail.Code)
	}

	if reply != nil {
		stdout.Encode(reply)
	}
}
