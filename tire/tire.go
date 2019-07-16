package tire

import "math"

type Info struct {
	Type          Type
	Width         float64
	AspectRatio   float64
	Construction  Construction
	Rim           float64
	Diameter      float64
	SideWall      float64
	Circumference float64
	Revolutions   Revolutions
}

type (
	Revolutions  float64
	Type         byte
	Construction byte
)

const (
	Passenger Type = 1 + iota
	LightTruck
)

const (
	Radial Construction = 1 + iota
)

func (r Revolutions) PerKilometre() float64 {
	return float64(r)
}

func (r Revolutions) PerMetre() float64 {
	return float64(r) / 1000.0
}

func (r Revolutions) PerMile() float64 {
	return float64(r) / 0.62000384402
}

func Parse(marking string) (Info, error) {
	i, err := (&parser{input: marking}).Parse()
	if err != nil {
		return Info{}, err
	}
	ratio := i.AspectRatio / 100.0
	i.Diameter = i.Rim + 2.0*i.Width*ratio
	i.SideWall = i.Width * ratio
	i.Circumference = i.Diameter * math.Pi
	i.Revolutions = Revolutions(1000000 / i.Circumference)
	return i, nil
}
