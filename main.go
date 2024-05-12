package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/01zulfi/jf/utils"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No arguments provided")
	}

	text := args[0]

	var out bytes.Buffer
	err := json.Indent(&out, []byte(text), "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	err = utils.SetClipboard(out.String())
	if err != nil {
		fmt.Println(out.String())
		return
	}

	fmt.Println(out.String())
	fmt.Println("")
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, "^ copied to clipboard")
	fmt.Println(colored)
}
