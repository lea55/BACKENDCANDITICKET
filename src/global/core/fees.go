package core

const FeePercent float32 = 10

func CalculateFee(value float32) float32 {
	percent := FeePercent * value
	total := percent / 100
	return total
}
