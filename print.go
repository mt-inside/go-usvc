package usvc

import (
	"fmt"
	"os"

	"github.com/mattn/go-isatty"
)

// PrintUpdateLn prints s to stdout.
// If stdout is connected to a terminal, it leaves the cursor at the end of the line and overwrites that line with any subsequent messages
// If stdout is not connected to a terminal, message are printed with newlines, just like fmt.Println()
// s should not contain any newline or other escape characters
// Calls to PrintUpdateLn should not be mixed with anything else that interacts with stdout
func PrintUpdateLn(s string) {
	if isatty.IsTerminal(os.Stdout.Fd()) {
		// Clear line, move cursor to beginning
		// These are ANSI escape sequences. \033[0G is equiv to \r
		// \r doesn't work - the cursor moves home, but nothing is cleared, so there's glitches if the new string is shorter
		// likewise repeated \b doesn't work; "backspace" just means move the cursor back; doesn't clear anything
		fmt.Printf("\033[K\033[0G")
	}

	fmt.Print(s)

	if !isatty.IsTerminal(os.Stdout.Fd()) {
		fmt.Println()
	}
}
