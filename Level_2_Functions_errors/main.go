package main

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

// возвращает текущий день недели (Понедельник, Вторник, Среда, Четверг, Пятница, Суббота, Воскресенье)
func currentDayOfTheWeek() string {
	t := TimeNow()

	switch t.Weekday() {
	case time.Monday:
		return "Понедельник"
	case time.Tuesday:
		return "Вторник"
	case time.Wednesday:
		return "Среда"
	case time.Thursday:
		return "Четверг"
	case time.Friday:
		return "Пятница"
	case time.Saturday:
		return "Суббота"
	case time.Sunday:
		return "Воскресенье"
	default:
		return ""
	}
}

// возвращает, День сейчас или Ночь. День для программы начинается в 10 утра, а заканчивается в 22 вечера включительно
func dayOrNight() string {
	t := TimeNow()

	if t.Hour() >= 10 && t.Hour() <= 22 {
		return "День"
	}

	return "Ночь"
}

// вычисляет, сколько дней до следующей пятницы, включая сегодняшний день
func nextFriday() int {
	t := TimeNow()

	day := int(t.Weekday())
	day += 1
	day %= 7

	return 6 - day
}

// проверяет какой сейчас день недели
func CheckCurrentDayOfTheWeek(answer string) bool {
	return strings.ToLower(currentDayOfTheWeek()) == strings.ToLower(answer)
}

// проверяет день сейчас или ночь
func CheckNowDayOrNight(answer string) (bool, error) {
	if utf8.RuneCountInString(answer) != 4 {
		return false, errors.New("исправь свой ответ, а лучше ложись поспать")
	}

	return strings.ToLower(dayOrNight()) == strings.ToLower(answer), nil
}
