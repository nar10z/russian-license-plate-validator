package russian_license_plate_validator

import (
	"regexp"
	"strings"
)

func ValidateLicense(license string) bool {
	return ValidateLicenseByType(license, Auto) ||
		ValidateLicenseByType(license, Trailer) ||
		ValidateLicenseByType(license, Motorcycle) ||
		ValidateLicenseByType(license, Scooter)
}

func ValidateLicenseByType(license string, t Type) bool {
	re := getREbyType(t)
	if re == nil {
		return false
	}

	license = strings.ToUpper(license)

	return re.MatchString(license)
}

func getREbyType(t Type) *regexp.Regexp {
	switch t {
	case Auto:
		return reAuto
	case Trailer:
		return reTrailers
	case Motorcycle:
		return reMotorcycle
	case Scooter:
		return reScooter
	}

	return nil
}
