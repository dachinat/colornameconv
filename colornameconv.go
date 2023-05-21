package colornameconv

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/crazy3lf/colorconv"
	"io"
	"math"
	"os"
	"path"
	"runtime"
	"strings"
)

type Colors struct {
	Colors []Color `json:"colors"`
}

type Color struct {
	Hex  string  `json:"hex"`
	Name string  `json:"name"`
	R    uint8   `json:"r"`
	G    uint8   `json:"g"`
	B    uint8   `json:"b"`
	H    float64 `json:"h"`
	S    float64 `json:"s"`
	L    float64 `json:"l"`
}

func New(hex string) (string, error) {
	hex = normalizeHex(hex)

	color, _ := colorconv.HexToColor(hex)
	h, s, l := colorconv.ColorToHSL(color)
	r, g, b, _ := colorconv.HexToRGB(hex)

	ndf1, ndf2, ndf := 0, 0, 0
	df, cl := -1, -1
	colors := readJson()

	for i, color := range colors.Colors {

		if hex == "#"+color.Hex {
			return color.Name, nil
		}

		ndf1 = int(math.Pow(float64(r-color.R), 2) + math.Pow(float64(g-color.G), 2) + math.Pow(float64(b-color.B), 2))
		ndf2 = int(math.Pow(h-color.H, 2) + math.Pow(s-color.S, 2) + math.Pow(l-color.L, 2))

		ndf = ndf1 + ndf2*2

		if df < 0 || df > ndf {
			df = ndf
			cl = i
		}
	}

	if cl < 0 {
		return "", errors.New("Invalid color definition")
	} else {
		return colors.Colors[cl].Name, nil
	}
}

func normalizeHex(hex string) string {
	hex = strings.ToUpper(hex)

	if !strings.HasPrefix(hex, "#") {
		hex = "#" + hex
	}

	return hex
}

func readJson() Colors {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	in, err := os.Open(path.Dir(filename) + "/colors.json")

	if err != nil {
		panic(fmt.Sprintf("Failed to read JSON file: %s", err))
	}

	byteValue, _ := io.ReadAll(in)

	var colors Colors

	err = json.Unmarshal(byteValue, &colors)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal JSON: %s", err))
	}

	return colors
}
