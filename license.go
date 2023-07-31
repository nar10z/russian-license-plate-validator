package russian_license_plate_validator

import "regexp"

const minLicenseSymbols = 7

var (
	reAuto       = regexp.MustCompile(`([АВЕКМНОРСТУХ])(\d{3})([АВЕКМНОРСТУХ]{2})(\d{2,3})`)
	reTrailers   = regexp.MustCompile(`([АВЕКМНОРСТУХ]{2})(\d{4})(\d{2,3})`)
	reMotorcycle = regexp.MustCompile(`(\d{4})([АВЕКМНОРСТУХ]{2})(\d{2,3})`)
	reScooter    = regexp.MustCompile(`([АВЕКМНОРСТУХ]{2})(\d{2})([АВЕКМНОРСТУХ]{2})(\d{2,3})`)
)

func NewLicense(license string) (License, error) {
	if len(license) < minLicenseSymbols {
		return License{}, ErrInvalidLicense
	}

	l := License{
		original: license,
	}

	switch {
	case reAuto.MatchString(license):
		l.t = Auto

		sub := reAuto.FindStringSubmatch(license)
		if len(sub) < 5 {
			return l, ErrInvalidLicense
		}

		l.series = sub[1] + " " + sub[3]
		l.registerNumber = sub[2]
		l.regionCode = sub[4]

	case reTrailers.MatchString(license):
		l.t = Trailer

		sub := reTrailers.FindStringSubmatch(license)
		if len(sub) < 4 {
			return l, ErrInvalidLicense
		}

		l.series = sub[1]
		l.registerNumber = sub[2]
		l.regionCode = sub[3]

	case reMotorcycle.MatchString(license):
		l.t = Motorcycle

		sub := reMotorcycle.FindStringSubmatch(license)
		if len(sub) < 4 {
			return l, ErrInvalidLicense
		}

		l.series = sub[2]
		l.registerNumber = sub[1]
		l.regionCode = sub[3]

	case reScooter.MatchString(license):
		l.t = Scooter

		sub := reScooter.FindStringSubmatch(license)
		if len(sub) < 5 {
			return l, ErrInvalidLicense
		}

		l.series = sub[1] + " " + sub[3]
		l.registerNumber = sub[2]
		l.regionCode = sub[4]

	default:
		return License{}, ErrInvalidLicense
	}

	return l, nil
}

type License struct {
	original       string
	series         string
	registerNumber string
	regionCode     string
	t              Type
}

func (l License) IsAuto() bool {
	return l.t == Auto
}

func (l License) IsTrailer() bool {
	return l.t == Trailer
}

func (l License) IsMotorcycle() bool {
	return l.t == Motorcycle
}

func (l License) IsScooter() bool {
	return l.t == Scooter
}

func (l License) Original() string {
	return l.original
}

func (l License) Series() string {
	return l.series
}

func (l License) RegisterNumber() string {
	return l.registerNumber
}

func (l License) RegionCode() string {
	return l.regionCode
}

func (l License) LicenseType() Type {
	return l.t
}
