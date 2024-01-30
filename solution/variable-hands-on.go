package solution

func VariableHandsOn() float64 {
	var sum float64
	productData := map[string][]float64{
		"Product A": {100000, 200, 0},
		"Product B": {67000, 12, 0.2},
		"Product C": {56000, 80, 0},
		"Product D": {1000, 1350, 0},
		"Product E": {20000, 1, 0},
		"Product F": {38455, 7, 0.15},
		"Product G": {76000, 5644, 0},
		"Product H": {530120, 30, 0.1},
		"Product I": {143000, 54, 0},
		"Product J": {16000, 109, 0},
	}

	for k := range productData {
		sum += (productData[k][0] * productData[k][1]) - (productData[k][0] * (productData[k][2]) * productData[k][1])
	}

	return sum
}
