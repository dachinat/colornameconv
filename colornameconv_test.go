package colornameconv

import "testing"

type addTest struct {
	arg1, expected string
}

var addTests = []addTest{
	addTest{"03AF1E", "Malachite"},
	addTest{"3EB831", "Apple"},
	addTest{"55EE20", "Bright Green"},
	addTest{"32C270", "Mountain Meadow"},
	addTest{"416A46", "Killarney"},
}

func TestAdd(t *testing.T){
	for _, test := range addTests {
		output, _ := New(test.arg1)
		if output != test.expected {
			t.Errorf("Color %q not equal to expected %q", output, test.expected)
		}
	}
}
