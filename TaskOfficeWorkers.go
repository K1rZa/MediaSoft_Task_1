package main

import "fmt"

type Employee struct {
	Name     string // имя
	Age      int    // возраст
	Position string // позиция
	Salary   int    // зарплата
}

var commands = `
1 - Добавить нового сотрудника
2 - Удалить сотрудника
3 - Вывести список сотрудников
4 - Выйти из программы
`

func main() {
	const size = 512
	empls := [size]*Employee{}
	for {
		cmd := 0
		fmt.Print(commands)
		fmt.Scanf("%d", &cmd)

		switch cmd {
		case 1:
			// Добавляем нового сотрудника
			empl := new(Employee)
			fmt.Println("\nИмя:")
			fmt.Scanf("%s", &empl.Name)
			fmt.Println("Возраст:")
			fmt.Scanf("%d", &empl.Age)
			fmt.Println("Позиция:")
			fmt.Scanf("%s", &empl.Position)
			fmt.Println("Зарплата:")
			fmt.Scanf("%d", &empl.Salary)
			for i := 0; i < size; i++ {
				if empls[i] == nil {
					empls[i] = empl
					break
				}
			}
		case 2:
			fmt.Println("Удаляем сотрудника")
		case 3:
			fmt.Println("Вывод сотрудников")
		case 4:
			break
		}
	}
}
