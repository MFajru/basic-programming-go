package solution

import "fmt"

func inputDuration() (int, error) {
	var duration int
	fmt.Print("Duration (in hours): ")
	_, err := fmt.Scanln(&duration)

	return duration, err
}

func calculateParkingPrice(duration int, vehicle string) int {
	parkingPrice := map[string][]int{
		"Motorcycle": {3000, 2000, 20000},
		"Car":        {7000, 5000, 50000},
	}
	var price int

	if duration <= 1 {
		price = parkingPrice[vehicle][0]
	} else if duration > 1 && duration <= 24 {
		duration -= 1
		price = parkingPrice[vehicle][0] + (parkingPrice[vehicle][1] * duration)
	} else {
		duration -= 1
		price = parkingPrice[vehicle][0] + (parkingPrice[vehicle][1] * duration) + parkingPrice[vehicle][2]
	}

	return price
}

func parking(vehicle string) {
	duration, err := inputDuration()

	if err != nil || duration < 0 {
		fmt.Println("Duration input invalid")
		return
	}

	price := calculateParkingPrice(duration, vehicle)

	fmt.Printf("Parking price: %d\n", price)

}

func ParkingMenu() (string, string) {
	fmt.Println("Parking Menu")
	strMotor := "Motorcycle"
	fmt.Printf("1. %s\n", strMotor)
	strCar := "Car"
	fmt.Printf("2. %s\n", strCar)
	fmt.Println("0. Exit")

	return strMotor, strCar
}

func LotBilling() {
	var menu int
	strMotor, strCar := ParkingMenu()

	fmt.Print("Choose your vehicle: ")
	_, err := fmt.Scanln(&menu)

	if err != nil {
		fmt.Println("Ticket invalid")
		return
	}

	switch menu {
	case 1:
		parking(strMotor)
	case 2:
		parking(strCar)
	case 0:
		return
	default:
		fmt.Printf("Vehicle number %d doesn't exist\n", menu)
	}

}
