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
	"github.com/mihaildemidoff/tnotify/internal/service"
	"github.com/spf13/cobra"
)

// sendFileCmd represents the sendFile command
var sendFileCmd = &cobra.Command{
	Use:   "send-file filename ...",
	Short: "Send file(s) from filesystem to telegram",
	Long:  "Sends files from filesystem to selected telegram chat. You can send several files just by enumerating them as arguments.",
	Example: `tnotify send-file test.txt
tnotify send-file test1.txt ../test2.txt`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, s := range args {
			service.SendFile(s)
		}
	},
}

func init() {
	RootCmd.AddCommand(sendFileCmd)
}
