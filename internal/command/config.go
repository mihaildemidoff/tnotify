/*
Copyright © 2021 Mikhail Demidov <mihaildemidoff@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package command

import (
	"github.com/mihaildemidoff/tnotify/internal/command/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Base configuration command",
	Long: `Base configuration command that allows you to setup application.
Application should be configured before usage. To properly setup tnotify app you need:
1. Register your bot at @BotFather and set your bot id with command 'tnotify config set-bot-token <your_bot_token>'
2. Create channel where you want to receive messages from tnotify. Add your bot to this channel and get and setup chat id with command 'tnotify config set-chat-id <your_chat_id>'
How to get chat id:
1. Create group and add your bot this group
2. Go to https://api.telegram.org/bot<YourBOTToken>/getUpdates and copy result[].message.chat.id
3. If previous link returned empty result try to remove and add your bot to group and try again previous step`,
	Example: `tnotify config set-bot-token <your_bot_token>
tnotify config set-chat-id <your_chat_id>
tnotify config get-bot-token
tnotify config get-chat-id`,
}

func init() {
	configCmd.AddCommand(config.SetBotTokenCmd)
	configCmd.AddCommand(config.SetChatIdCmd)
	configCmd.AddCommand(config.GetBotTokenCmd)
	configCmd.AddCommand(config.GetChatIdCmd)
}
