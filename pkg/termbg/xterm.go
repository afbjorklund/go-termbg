// +build !windows

package termbg

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	termx "golang.org/x/term"
)

func fromXterm(term Terminal, timeout time.Duration) (RGB, error) {
	// Query by XTerm control sequence
	var query string
	if term == Tmux {
		query = "\x1bPtmux;\x1b\x1b]11;?\x07\x1b\\\x03"
	} else if term == Screen {
		query = "\x1bP\x1b]11;?\x07\x1b\\\x03"
	} else {
		query = "\x1b]11;?\x1b\\"
	}

	oldState, err := termx.MakeRaw(int(os.Stderr.Fd()))
	if err != nil {
		panic(err)
	}
	defer termx.Restore(int(os.Stderr.Fd()), oldState)
	fmt.Fprint(os.Stderr, query)

	buffer := []byte{}
	start := false
	var buf []byte = make([]byte, 1)
	for {
		os.Stdin.Read(buf)
		// response terminated by BEL(0x7)
		if start && (buf[0] == 0x7) {
			break
		}
		// response terminated by ST(0x1b 0x5c)
		if start && (buf[0] == 0x1b) {
			// consume last 0x5c
			os.Stdin.Read(buf)
			break
		}
		if start {
			buffer = append(buffer, buf[0])
		}
		if buf[0] == ':' {
			start = true
		}
	}

	s := string(buffer)
	r, g, b, err := decodeX11Color(s)
	if err != nil {
		return RGB{}, err
	}
	return RGB{r, g, b}, nil
}

func decodeHex(s string) (uint16, error) {
	ret, err := strconv.ParseUint(s, 16, 16)
	if err != nil {
		return 0, err
	}
	ret = ret << ((4 - len(s)) * 4)
	return uint16(ret), nil
}

func decodeX11Color(s string) (uint16, uint16, uint16, error) {
	rgb := strings.Split(s, "/")

	r, err := decodeHex(rgb[0])
	if err != nil {
		return 0, 0, 0, err
	}
	g, err := decodeHex(rgb[1])
	if err != nil {
		return 0, 0, 0, err
	}
	b, err := decodeHex(rgb[2])
	if err != nil {
		return 0, 0, 0, err
	}

	return r, g, b, nil
}
