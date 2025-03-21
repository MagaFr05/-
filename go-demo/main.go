package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("__Калькулятор индекса массы тела__")
	for {

		userKg, userHeight := getUserInput()
		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			fmt.Println(err)
			continue
			//panic("Не заданы парматры для расчета")
		}
		outputResult(IMT)
		isRepeteCalculation := checkReppeatCalculation()
		if !isRepeteCalculation {
			break
		}

	}

}
func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела:%.0f", IMT)
	fmt.Println(result)
	switch {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("У вас дефицит массы тела ")
	case IMT < 25:
		fmt.Println("У вас нормальный вес")
	case IMT < 30:
		fmt.Println("У вас избыточная масса ")
	default:
		fmt.Println("У вас степень ожирения")
	}

}

func calculateIMT(userKg float64, userHeight float64) (float64, error) {
	if userHeight <= 0 || userKg <= 0 {
		return 0, errors.New("Не указан вес или высота")
	}
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}
func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост в сантиметрах:")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в килограммах:")
	fmt.Scan(&userKg)
	return userKg, userHeight
}
func checkReppeatCalculation() bool {
	var userChoise string
	fmt.Print("Вы хотите сделать еще расчет(Y/n):")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
