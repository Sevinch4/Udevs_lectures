package workplace

type Workplace struct {
	PlaceName string
	Salary    int
	Level     string
}

func (w Workplace) Worker(placeName string, salary int, level string) Workplace {
	return Workplace{
		PlaceName: placeName,
		Salary:    salary,
		Level:     level,
	}
}
