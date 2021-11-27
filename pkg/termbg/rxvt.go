// +build !windows

package termbg

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fromRxvt() (RGB, error) {
	env := os.Getenv("COLORFGBG")
	fgbg := strings.Split(env, ";")
	if len(fgbg) < 2 {
		return RGB{}, fmt.Errorf("unknown fgbg %s", env)
	}
	// fg;bg
	bg, err := strconv.ParseUint(fgbg[1], 10, 8)
	if err != nil {
		// fg;xpm;bg
		bg, err = strconv.ParseUint(fgbg[2], 10, 8)
		if err != nil {
			return RGB{}, err
		}
	}
	r, g, b := colorTable(uint(bg))
	return RGB{
		R: r * 256,
		G: g * 256,
		B: b * 256,
	}, nil
}

func colorTable(bg uint) (uint16, uint16, uint16) {
	switch bg {
	// black
	case 0:
		return 0, 0, 0
	// red
	case 1:
		return 205, 0, 0
	// green
	case 2:
		return 0, 205, 0
	// yellow
	case 3:
		return 205, 205, 0
	// blue
	case 4:
		return 0, 0, 238
	// magenta
	case 5:
		return 205, 0, 205
	// cyan
	case 6:
		return 0, 205, 205
	// white
	case 7:
		return 229, 229, 229
	// bright black
	case 8:
		return 127, 127, 127
	// bright red
	case 9:
		return 255, 0, 0
	// bright green
	case 10:
		return 0, 255, 0
	// bright yellow
	case 11:
		return 255, 255, 0
	// bright blue
	case 12:
		return 92, 92, 255
	// bright magenta
	case 13:
		return 255, 0, 255
	// bright cyan
	case 14:
		return 0, 255, 255
	// bright white
	case 15:
		return 255, 255, 255
	default:
		return 0, 0, 0
	}
}
