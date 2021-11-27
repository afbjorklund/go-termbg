package termbg

import (
	"time"
)

type Theme string

const (
	Light Theme = "Light"
	Dark  Theme = "Dark"
)

func NewTheme(timeout time.Duration) (Theme, error) {
	rgb, err := NewRGB(timeout)
	if err != nil {
		return "", err
	}

	// ITU-R BT.601
	y := float64(rgb.R)*0.299 + float64(rgb.G)*0.587 + float64(rgb.B)*0.114

	if y > 32768.0 {
		return Light, nil
	} else {
		return Dark, nil
	}
}
