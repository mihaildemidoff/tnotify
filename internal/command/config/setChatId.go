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
package config

import (
	"github.com/mihaildemidoff/tnotify/internal/service"
	"github.com/spf13/cobra"
)

var SetChatIdCmd = &cobra.Command{
	Use:   "set-chat-id",
	Short: "Set chat id",
	Long: `Sets chat id. How to get chat id:
1. Create new group and add your bot.
2. Go to https://api.telegram.org/bot<YourBOTToken>/getUpdates and copy result[].message.chat.id
3. If result of previous step is empty try to remove and add your bot to group`,
	Example:            "config set-chat-id <your_chat_id>",
	Args:               cobra.MinimumNArgs(1),
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		service.WriteChatId(args[0])
	},
}

func init() {
}
