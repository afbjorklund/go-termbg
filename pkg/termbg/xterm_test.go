package termbg

import (
	"testing"

	"gotest.tools/v3/assert"
)

func decodeX11ColorValues(s string) []uint16 {
	r, g, b, err := decodeX11Color(s)
	if err != nil {
		return nil
	}
	return []uint16{r, g, b}
}

func TestDecodeX11Color(t *testing.T) {
	s := "0000/0000/0000"
	assert.DeepEqual(t, decodeX11ColorValues(s), []uint16{0, 0, 0})

	s = "1111/2222/3333"
	assert.DeepEqual(t, decodeX11ColorValues(s), []uint16{0x1111, 0x2222, 0x3333})

	s = "111/222/333"
	assert.DeepEqual(t, decodeX11ColorValues(s), []uint16{0x1110, 0x2220, 0x3330})

	s = "11/22/33"
	assert.DeepEqual(t, decodeX11ColorValues(s), []uint16{0x1100, 0x2200, 0x3300})

	s = "1/2/3"
	assert.DeepEqual(t, decodeX11ColorValues(s), []uint16{0x1000, 0x2000, 0x3000})
}
