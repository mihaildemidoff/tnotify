package service

import (
	"bufio"
	"io"
	"os"
)

func ReadFromPipeOrFromArguments(args []string) string {
	if len(args) > 0 {
		return args[0]
	} else {
		info, err := os.Stdin.Stat()
		if err != nil {
			panic(err)
		}
		if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
			return ""
		}
		reader := bufio.NewReader(os.Stdin)
		var output []rune

		for {
			input, _, err := reader.ReadRune()
			if err != nil && err == io.EOF {
				break
			}
			output = append(output, input)
		}
		return string(output)
	}
}
