
# Slack File Bot
## Project Structure
```
slack-file-bot/
│
├── NEW.txt
├── go.mod
├── go.sum
├── main.go
└── README.md
```
## Introduction
The Slack File Bot is a customizable bot designed to automate file uploads and management tasks in Slack. It simplifies file handling within Slack channels, making your team's workflow more efficient by leveraging Slack's API for seamless file organization and sharing.

## Requirements
To run this bot, you will need:

- Go (Golang) installed on your system.
- A Slack workspace with a bot token.
- The Slack-Go SDK.
- A Slack channel where the files will be uploaded.
## Setup Instructions
### 1. **Clone the Repository**
```bash
git clone https://github.com/MikeOnBoard/slack-file-bot.git
cd slack-file-bot
```
### 2. **Install Dependencies**
```bash
go mod tidy
```
### 3. **Configure Environment Variables**
Before running the bot, configure the following environment variables:

```GO
export SLACK_BOT_TOKEN="your-slack-bot-token"
export CHANNEL_ID="your-channel-id"
```
### 4. **Run the Bot**
To start the bot, simply run the following command:

```bash
go run main.go
```
This command will automatically upload a specified file (in this case, ``NEW.txt``) to your selected Slack channel.

## File Customization
The bot uploads files stored in the ``fileArr`` array located in the ``main.go`` file. You can modify this array to include different files as per your requirements. The bot will iterate through this list and upload each file to the designated Slack channel.

## How It Works
The bot leverages the Slack-Go SDK to interact with Slack's API, utilizing the ``files.upload`` method to upload files. You can automate file uploads to streamline tasks such as sharing documents, logs, or other data with your team directly within Slack channels.

### Slack-Go SDK Installation
The bot uses the Slack-Go SDK. Install it by running:

```bash
go get -u github.com/slack-go/slack
```
## Conclusion
The Slack File Bot streamlines file uploads within Slack channels, providing an easy-to-use interface for automating and managing file sharing. This helps reduce manual work and optimizes team collaboration.

