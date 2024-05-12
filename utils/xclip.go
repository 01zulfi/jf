package utils

import (
	"os/exec"
	"strings"
)

func SetClipboard(value string) error {
	cmd := exec.Command("xclip", "-selection", "c")
	cmd.Stdin = strings.NewReader(value)
	return cmd.Run()
}
