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

//Transport Ввод и вывод информации о транспортных средствах.
func Transport() {
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
	fmt.Println(car1.brandCar," - марка легкового автомобиля.")
	fmt.Println(car1.yearCar," - год выпуска легкового автомобиля.")
	fmt.Println(car1.trunkVolumeCar," - общий объём багажного отделения легкового автомоблия.")
	fmt.Println(car1.trunkOccupancyCar," - занятный объём багажного отделения легкового автомоблия.")
	
	if car1.engineCarOn == false {
		fmt.Println(car1.brandCar, " двигатель включен.")
	} else {
		fmt.Println(car1.brandCar," двигатель включен.")
	}

	if (car1.windowCar1 == false) && (car1.windowCar2 == false) && (car1.windowCar3 == false) && (car1.windowCar4 == false) {
		fmt.Println(car1.brandCar," - Все окна закрыты.")
	} else {
		fmt.Println(car1.brandCar,"Одно или несколько окон открыты.")
	}
	
	//Грузовой транспорт
	fmt.Println("Грузовой транспорт:")

	fmt.Println(truck1)
	fmt.Println(truck1.brandTruck," - марка грузового автомобиля.")
	fmt.Println(truck1.yearTruck," - год выпуска грузового автомобиля.")
	fmt.Println(truck1.loadingCapacityTruck," - общий объём кузова грузового автомобиля в килограммах.")
	fmt.Println(truck1.trunkOccupancyTruck," - занятный объём кузова грузового автомобиля.")

	if truck1.engineCarTruck == false {
		fmt.Println(truck1.brandTruck," двигатель выключен.")
	} else {
		fmt.Println(truck1.brandTruck," двигатель включен.")
	}

	if (truck1.windowTruck1 == false) && (truck1.windowTruck2 == false) {
		fmt.Println(truck1.brandTruck," - Все окна закрыты.")
	} else {
		fmt.Println(truck1.brandTruck,"Одно или несколько окон открыты.")
	}

}