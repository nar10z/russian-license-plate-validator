package russian_license_plate_validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateLicense(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		license string
		want    bool
	}{
		{
			name:    "Empty",
			license: "",
			want:    false,
		},
		{
			name:    "Invalid",
			license: "dasd as3",
			want:    false,
		},
		{
			name:    "Auto",
			license: "С227НА69",
			want:    true,
		},
		{
			name:    "Trailer",
			license: "АН733157",
			want:    true,
		},
		{
			name:    "Motorcycle",
			license: "8776АЕ64",
			want:    true,
		},
		{
			name:    "Scooter",
			license: "ММ55АА23",
			want:    true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, tt.want, ValidateLicense(tt.license), "ValidateLicense(%v)", tt.license)
		})
	}
}
