package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	// Set environment variables
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-7897382236550-7916763812273-eaXeREgP4XrJMit4c96qRA7p")
	os.Setenv("CHANNEL_ID", "C07SH4LPHK7")

	// Initialize Slack API
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	// Define channel and file to upload
	channelID := os.Getenv("CHANNEL_ID")
	filePath := "/home/mike/Documentos/56ProyectosConGolang/slack-file-bot/NEW.txt"

	// Upload file using the regular UploadFileV2
	uploadResponse, err := api.UploadFile(slack.FileUploadParameters{
		Channels: []string{channelID},
		File:     filePath,
	})

	if err != nil {
		fmt.Printf("Error uploading file: %s\n", err)
		return
	}

	fmt.Printf("File uploaded successfully: ID=%s, Name=%s\n", uploadResponse.ID, uploadResponse.Name)
}
