package russian_license_plate_validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLicense(t *testing.T) {
	t.Parallel()

	t.Run("Invalid", func(t *testing.T) {
		t.Parallel()

		_, err := NewLicense("")
		assert.Error(t, err)
	})
	t.Run("Invalid", func(t *testing.T) {
		t.Parallel()

		_, err := NewLicense("fasfd ssfd s")
		assert.Error(t, err)
	})

	t.Run("Auto", func(t *testing.T) {
		t.Parallel()

		license := "С227НА69"

		l, err := NewLicense(license)
		assert.NoError(t, err)

		assert.True(t, l.IsAuto())
		assert.Equal(t, Auto, l.LicenseType())
		assert.Equal(t, license, l.Original())
		assert.Equal(t, "С НА", l.Series())
		assert.Equal(t, "227", l.RegisterNumber())
		assert.Equal(t, "69", l.RegionCode())
	})
	t.Run("Trailer", func(t *testing.T) {
		t.Parallel()

		license := "АН733157"

		l, err := NewLicense(license)
		assert.NoError(t, err)

		assert.True(t, l.IsTrailer())
		assert.Equal(t, Trailer, l.LicenseType())
		assert.Equal(t, license, l.Original())
		assert.Equal(t, "АН", l.Series())
		assert.Equal(t, "7331", l.RegisterNumber())
		assert.Equal(t, "57", l.RegionCode())
	})
	t.Run("Motorcycle", func(t *testing.T) {
		t.Parallel()

		license := "8776АЕ64"

		l, err := NewLicense(license)
		assert.NoError(t, err)

		assert.True(t, l.IsMotorcycle())
		assert.Equal(t, Motorcycle, l.LicenseType())
		assert.Equal(t, license, l.Original())
		assert.Equal(t, "АЕ", l.Series())
		assert.Equal(t, "8776", l.RegisterNumber())
		assert.Equal(t, "64", l.RegionCode())
	})
	t.Run("Scooter", func(t *testing.T) {
		t.Parallel()

		license := "ММ55АА23"

		l, err := NewLicense(license)
		assert.NoError(t, err)

		assert.True(t, l.IsScooter())
		assert.Equal(t, Scooter, l.LicenseType())
		assert.Equal(t, license, l.Original())
		assert.Equal(t, "ММ АА", l.Series())
		assert.Equal(t, "55", l.RegisterNumber())
		assert.Equal(t, "23", l.RegionCode())
	})
}
