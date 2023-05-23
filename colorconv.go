package colornameconv

import (
	"math"
	"strconv"
)

func HexToHSL(color string) (int, int, int) {
	r, _ := strconv.ParseInt(color[1:3], 16, 64)
	g, _ := strconv.ParseInt(color[3:5], 16, 64)
	b, _ := strconv.ParseInt(color[5:7], 16, 64)

	rgb := [3]float64{float64(r) / 255, float64(g) / 255, float64(b) / 255}

	min := math.Min(math.Min(rgb[0], rgb[1]), rgb[2])
	max := math.Max(math.Max(rgb[0], rgb[1]), rgb[2])
	delta := max - min
	l := (min + max) / 2

	s := 0.0;
	d := 0.0;
	if l > 0 && l < 1 {

		if l < 0.5 {
			d = 2 * l
		} else {
			d = 2 - 2 * l
		}

		s = delta / d
	}

	h := 0.0
	if delta > 0 {
		if max == rgb[0] && max != rgb[1] {
			h += (rgb[1] - rgb[2]) / delta
		}
		if max == rgb[1] && max != rgb[2] {
			h += (2 + (rgb[2]-rgb[0])/delta)
		}
		if max == rgb[2] && max != rgb[0] {
			h += (4 + (rgb[0]-rgb[1])/delta)
		}
		h /= 6
	}

	return int(h * 255), int(s * 255), int(l * 255)
}
