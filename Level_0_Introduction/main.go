package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var date string
	var fname, lname, patr string
	var s1, s2, s3 float64

	fmt.Scan(&date, &fname, &lname, &patr, &s1, &s2, &s3)
	sum := s1 + s2 + s3

	// dd.MM.yyyy
	const Form = "02.01.2006"
	time1, _ := time.Parse(Form, date)
	time1 = time1.Add(time.Hour * 24 * 15)

	fmt.Printf(
		"Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.\n"+
			"Дата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день.\n"+
			"Общая сумма выплат составит %.0f руб. %.0f коп.\n\n"+
			"С уважением,\n"+
			"Гл. бух. Иванов А.Е.\n",
		lname,
		fname,
		patr,
		time1.Format(Form),
		math.Floor(sum),
		math.Floor(sum*100.0)-math.Floor(sum)*100,
	)
}
