package russian_license_plate_validator

type Type int

const (
	Auto Type = iota
	Trailer
	Motorcycle
	Scooter
)
