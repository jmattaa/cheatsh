package formatter

import (
	"os"
	"os/exec"
	"strings"
)

func Print(text string) {
	// Check if "mdcat" is available
	if _, err := exec.LookPath("glow"); err != nil {
		// Fallback: just print the plain text.
		println(text)
		return
	}

	cmd := exec.Command("glow")
	cmd.Stdin = strings.NewReader(text)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		println(text)
	}
}

