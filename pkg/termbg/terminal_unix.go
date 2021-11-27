package termbg

import (
	"os"
	"strings"
)

func NewTerminal() (Terminal, error) {
	program, ok := os.LookupEnv("TERM_PROGRAM")
	if ok {
		if program == "vscode" {
			return VSCode, nil
		}
	}

	if _, ok := os.LookupEnv("TMUX"); ok {
		return Tmux, nil
	} else {
		term := os.Getenv("TERM")
		if strings.HasPrefix(term, "screen") {
			return Screen, nil
			//} else if strings.HasPrefix(term, "rxvt") {
			//	return RxvtCompatible, nil
		} else {
			return XtermCompatible, nil
		}
	}
}
