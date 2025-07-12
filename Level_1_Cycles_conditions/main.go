package main

import (
	"fmt"
)

func main() {
	f, s, t, fo, fi := "-", "-", "-", "-", "-"
	load := 0

	ww, ol, bm := "некорректный ввод", "очередь переполнена", "место уже занято"

	var com string
	for {
		fmt.Scan(&com)

		switch com {
		case "очередь", "конец":
			fmt.Printf(
				"1. %s\n"+
					"2. %s\n"+
					"3. %s\n"+
					"4. %s\n"+
					"5. %s\n",
				f, s, t, fo, fi,
			)
		case "количество":
			fmt.Printf("Осталось свободных мест: %d\n"+
				"Всего человек в очереди: %d\n",
				5-load, load,
			)
		default:
			var pos string
			fmt.Scanln(&pos)

			switch pos {
			case "1":
				if load == 5 {
					fmt.Printf("Запись на место номер %s невозможна: %s\n", pos, ol)
					continue
				}

				if f == "-" {
					f = com
					load++
				} else {
					fmt.Printf("Запись на место номер %s невозможна: %s\n", pos, bm)
					continue
				}
			case "2":
				if load == 5 {
					fmt.Printf("Запись на место номер %s невозможна: %s\n", pos, ol)
					continue
				}

				if s == "-" {
					s = com
					load++
				} else {
					fmt.Printf("Запись на место номер %s невозможна: %s\n", pos, bm)
					continue
				}
			case "3":
				if load == 5 {
					fmt.Printf("Запись на место номер %s невозможна: %s\n", pos, ol)
					continue
				}

				if t == "-" {
					t = com
					load++
				} else {
					fmt.Printf("Запись на место номер %s невозможна: %s\n", pos, bm)
					continue
				}
			case "4":
				if load == 5 {
					fmt.Printf("Запись на место номер %s невозможна: %s\n", pos, ol)
					continue
				}

				if fo == "-" {
					fo = com
					load++
				} else {
					fmt.Printf("Запись на место номер %s невозможна: %s\n", pos, bm)
					continue
				}
			case "5":
				if load == 5 {
					fmt.Printf("Запись на место номер %s невозможна: %s\n", pos, ol)
					continue
				}

				if fi == "-" {
					fi = com
					load++
				} else {
					fmt.Printf("Запись на место номер %s невозможна: %s\n", pos, bm)
					continue
				}
			default:
				fmt.Printf("Запись на место номер %s невозможна: %s\n", pos, ww)
				continue
			}
		}

		if com == "конец" {
			break
		}
	}
}
