package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	earthYears := seconds / 31557600.0;
	switch planet {
		case "Mercury":
			return earthYears / 0.2408467;
		case "Venus":
			return earthYears / 0.61519726;
		case "Earth":
			return earthYears;
		case "Mars":
			return earthYears / 1.8808158;
		case "Jupiter":
			return earthYears / 11.862615;
		case "Saturn":
			return earthYears / 29.447498;
		case "Uranus":
			return earthYears / 84.016846;
		case "Neptune":
			return earthYears / 164.79132;
		default:
			return earthYears;
	}
}
