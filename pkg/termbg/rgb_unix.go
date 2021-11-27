package termbg

import (
	"fmt"
	"time"
)

func NewRGB(timeout time.Duration) (RGB, error) {
	term, err := NewTerminal()
	if err != nil {
		return RGB{}, err
	}
	switch term {
	case VSCode:
		return RGB{}, fmt.Errorf("unsupported")
	case RxvtCompatible:
		return fromRxvt()
	case XtermCompatible:
		return fromXterm(term, timeout)
	default:
		return RGB{}, fmt.Errorf("unknown")
	}
}
