package main

import (
	"flag"
	"fmt"
	"hw-1/internal/slice"
	"log"
)

type Operation int

const (
	OperationExit Operation = iota
	OperationAppend
	OperationRemove
	OperationAddOne
)

func main() {
	sizeMax, err := parseFlags()
	if err != nil {
		log.Fatal("parse flags: ", err)
	}

	if err := run(sizeMax); err != nil {
		log.Fatal("run: ", err)
	}
}

func run(sizeMax int) error {
	var s []int

	for {
		printArray(s)
		printMenu()

		var choice Operation
		if _, err := fmt.Scan(&choice); err != nil {
			continue
		}

		switch choice {
		case OperationExit:
			fmt.Println("Программа завершена.")
			return nil

		case OperationAppend:
			fmt.Print("Введите элемент для добавления: ")

			var elem int

			for {
				if _, err := fmt.Scan(&elem); err != nil {
					fmt.Println("Введите целое число")
					continue
				}

				break
			}

			s = slice.AppendElement(s, elem)

		case OperationRemove:
			s = slice.RemoveElement(s)

		case OperationAddOne:
			s = slice.AddOneToAll(s)

		default:
			fmt.Println("Неверный выбор. Попробуйте еще раз.")
		}

		if len(s) >= sizeMax {
			break
		}
	}

	return nil
}

func parseFlags() (sizeMax int, err error) {
	flag.IntVar(&sizeMax, "NMAX", 0, "max slice size (required)")

	flag.Parse()

	if sizeMax == 0 {
		return 0, fmt.Errorf("missing required argument NMAX")
	}

	if sizeMax < 0 {
		return 0, fmt.Errorf("NMAX and NMIN must not be less than 0")
	}

	return
}

func printMenu() {
	fmt.Print("Меню:" +
		"\n\t1 - Добавить элемент" +
		"\n\t2 - Удалить элемент" +
		"\n\t3 - Добавить 1 к каждому элементу" +
		"\n\t0 - Выйти" +
		"\nВыберите операцию (0-3): ")
}

// Вспомогательная функция для вывода массива на экран
func printArray(array []int) {
	fmt.Println("Текущий массив:", array)
}
