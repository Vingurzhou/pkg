package satellite

import (
	"testing"
)

// func TestNewSatellite(t *testing.T) {
// 	type args struct {
// 		line1 string
// 		line2 string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *Satellite
// 		wantErr bool
// 	}{{
// 		name: "",
// 		args: args{
// 			line1: "1 65055U 25164A   25329.77566319 -.00034908  00000-0 -22699-2 0 00003",
// 			line2: "2 65055 041.1321 355.6156 0004144 325.3028 090.7476 15.07206270017711",
// 		},
// 		want: &Satellite{
// 			Line1:             "1 65055U 25164A   25329.77566319 -.00034908  00000-0 -22699-2 0 00003",
// 			Line2:             "2 65055 041.1321 355.6156 0004144 325.3028 090.7476 15.07206270017711",
// 			MeanMotion:        15.0720627,
// 			SemiMajorAxis:     6922.8785581607035,
// 			Eccentricity:      "0.0004144",
// 			Inclination:       "041.1321",
// 			RAAN:              "355.6156",
// 			ArgumentOfPerigee: "325.3028",
// 			MeanAnomaly:       "090.7476",
// 			EpochTime:         time.Date(2025, 11, 25, 18, 36, 57, 299615999, time.UTC),
// 		},
// 		wantErr: false,
// 	}}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := NewSatellite(tt.args.line1, tt.args.line2)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("NewSatellite() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewSatellite() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestGetOsculatingElements(t *testing.T) {
	t.Log(NewClassicalOrbitalElements(
		"1 65055U 25164A   25329.77566319 -.00034908  00000-0 -22699-2 0 00003",
		"2 65055 041.1321 355.6156 0004144 325.3028 090.7476 15.07206270017711"))
}
