package results

type Measurement struct {
	CurrentValue float64
	MinValue     float64
	AverageValue float64
	MaxValue     float64
}

func NewMeasurement(currentValue float64, minValue float64,
	averageValue float64, maxValue float64) *Measurement {
	return &Measurement{
		CurrentValue: currentValue,
		MinValue:     minValue,
		AverageValue: averageValue,
		MaxValue:     maxValue,
	}
}
