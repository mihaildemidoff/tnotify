# Tnotify
Simple CLI application that allows you to send messages to telegram group from terminal.
# Build
## Prerequisites
- go 1.16+
- make
- docker
## Build and install
- Vendor with `go mod vendor`
- Build application for your current OS and architecture with `make build`
- Install application with `cd cmd/tnotify && go install`
# Setup application
You need to setup application before usage. You need to provide:
- Your bot access token
- Chat id where you want to receive messages from application
## Register bot and get bot token 
Follow instruction on https://core.telegram.org/bots#6-botfather and set bot token with `tnotify config set-bot-token <your_bot_token>`
## Add bot to chat and get chat id
In order to get messages you need to create new group and add your bot to this group.
Go to https://api.telegram.org/bot<YourBOTToken>/getUpdates and copy chat id.
Set chat id to tnotify with `tnotify config set-chat-id <your_chat_id>`
# Usage
- `tnotify 'your message'` - send text message to chat
- `tnotify 'your message' --file` - send text message to chat as file
- `tnotify send-file file.txt` - send file to chat
- `tnotify config set-bot-token <your_bot_token>` - set bot token to configuration
- `tnotify config set-chat-id <your_chat_id>` - set chat id to configuration
- `tnotify config get-bot-token` - get bot token from configuration
- `tnotify config get-chat-id` - get chat id from configuration
