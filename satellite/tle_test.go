package satellite

import (
	"testing"
)

func TestNewSatellite(t *testing.T) {
	sat, err := NewSatellite("1 65055U 25164A   25329.77566319 -.00034908  00000-0 -22699-2 0 00003",
		"2 65055 041.1321 355.6156 0004144 325.3028 090.7476 15.07206270017711")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", sat)
}
