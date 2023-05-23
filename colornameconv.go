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
	Hex  string `json:"hex"`
	Name string `json:"name"`
	R    int    `json:"r"`
	G    int    `json:"g"`
	B    int    `json:"b"`
	H    int    `json:"h"`
	S    int    `json:"s"`
	L    int    `json:"l"`
}

func New(hex string) (string, error) {
	hex = normalizeHex(hex)

	r, g, b, _ := colorconv.HexToRGB(hex)

	h, s, l := HexToHSL(hex)

	ndf1, ndf2, ndf := 0, 0, 0
	cl, df := -1, -1
	colors := readJson()

	for i := 0; i < len(colors.Colors); i++ {
		color := colors.Colors[i]

		if hex == "#"+color.Hex {
			return color.Name, nil
		}

		ndf1 = int(math.Pow(float64(int(r)-color.R), 2) + math.Pow(float64(int(g)-color.G), 2) + math.Pow(float64(int(b)-color.B), 2))
		ndf2 = int(math.Pow(float64(h-color.H), 2) + math.Pow(float64(s-color.S), 2) + math.Pow(float64(l-color.L), 2))
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
