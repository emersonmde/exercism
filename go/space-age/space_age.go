package space

import "fmt"

type Planet string

var planets = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(age float64, planet Planet) float64 {
	fmt.Printf("planet = %+v\n", planet)
	fmt.Printf("age = %+v\n", age)

	earthAge := age / 31557600.0
	return earthAge / planets[planet]
}
