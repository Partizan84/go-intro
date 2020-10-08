package transport

import "fmt"

type windowsCarOpen struct {
	windowCar1 bool
	windowCar2 bool
	windowCar3 bool
	windowCar4 bool
}

type car struct {
	brandCar string
	yearCar int
	trunkVolumeCar int
	trunkOccupancyCar int
	engineCarOn bool
	windowsCarOpen
}

type windowTruckOpen struct{
	windowTruck1 bool
	windowTruck2 bool
}

type truck struct {
	brandTruck string
	yearTruck int
	loadingCapacityTruck int
	trunkOccupancyTruck int
	engineCarTruck bool
	windowTruckOpen
}

func transport() {
	car1 := car{
		brandCar: "Subaru XV",
		yearCar: 2019,
		trunkVolumeCar: 310,
		trunkOccupancyCar: 150,
		engineCarOn: true,
		windowsCarOpen: windowsCarOpen{
			windowCar1: true,
		},
	}

	truck1 := truck{
		brandTruck: "Ural NEXT 4320-5111-73",
		yearTruck: 2020,
		loadingCapacityTruck: 6500,
		trunkOccupancyTruck: 1500,
		engineCarTruck: false,
		windowTruckOpen: windowTruckOpen{
			windowTruck1: false,
			windowTruck2: false,
		},
	}
	
	//Легковой транспорт
	fmt.Println("Легковой автомобиль:")

	fmt.Println(car1)
	
	if car1.engineCarOn == false {
		fmt.Println(car1.brandCar, " двигатель включен.")
	} else {
		fmt.Println(car1.brandCar," двигатель включен.")
	}
	
	//Грузовой транспорт
	fmt.Println("Грузовой транспорт:")

	fmt.Println(truck1)

	if truck1.engineCarTruck == false {
		fmt.Println(truck1.brandTruck," двигатель выключен.")
	} else {
		fmt.Println(truck1.brandTruck," двигатель включен.")
	}

}