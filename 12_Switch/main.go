package main

import "fmt"

func dayOfTheWeek(number int) string {
	switch number {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Invalid day number"
	}
}

func dayOfTheWeek2(number int) string {
	var weekDay string

	switch {
	case number == 1:
		weekDay = "Domingo"
		// Falltrough makes Go jump INSIDE the next condition (without checking it)
		fallthrough
	case number == 2:
		weekDay = "Segunda-feira"
	case number == 3:
		weekDay = "Terça-feira"
	case number == 4:
		weekDay = "Quarta-feira"
	case number == 5:
		weekDay = "Quinta-feira"
	case number == 6:
		weekDay = "Sexta-feira"
	case number == 7:
		weekDay = "Sábado"
	}

	return weekDay
}

func main() {
	day := dayOfTheWeek(2)

	fmt.Println(day)

	day2 := dayOfTheWeek2(1)

	fmt.Println(day2)
}
