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
	"fmt"
	"github.com/mihaildemidoff/tnotify/internal/service"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "tnotify",
	Short: "Telegram notification app",
	Long:  `Application that allows you to send messages to telegram chat from terminal.`,
	Example: `tnotify 'message'
tnotify 'message' --file
tnotify send-file test.txt
tnotify config set-bot-token <your_bot_token>
tnotify config set-chat-id <your_chat_id>
cat test.txt | tnotify
cat test.txt | tnotify --file`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		message := service.ReadFromPipeOrFromArguments(args)
		if len(message) == 0 {
			_ = cmd.Help()
			os.Exit(0)
		}
		configured := service.CheckConfigured()
		if !configured {
			fmt.Printf("Tnotify is not properly configured! Set bot token and chat id.\n")
			configCmd.Help()
			os.Exit(0)
		}
		isFile, err := cmd.Flags().GetBool("file")
		if err != nil {
			isFile = false
		}
		service.SendMessage(message, isFile)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.AddCommand(configCmd)
	RootCmd.AddCommand(sendFileCmd)
	RootCmd.Flags().Bool("file", false, "Send message as a file")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		service.InitConfigFile()
	}

	viper.AutomaticEnv()

	viper.ReadInConfig()
}
