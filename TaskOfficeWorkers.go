package main
import "fmt"

type Employee struct {
	Name     string
	Age      int  
	Position string 
	Salary   int 
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
	count := 0
	
	for {
		cmd := 0
		fmt.Print(commands)
		fmt.Scanf("%d\n", &cmd)

		switch cmd {
		case 1:
			if count >= size {
				fmt.Println("ERROR! 512 сотрудников уже добавлены")
				continue
			}
			empl := new(Employee)
			fmt.Print("Имя: ")
			fmt.Scanf("%s\n", &empl.Name)
			fmt.Print("Возраст: ")
			fmt.Scanf("%d\n", &empl.Age)
			fmt.Print("Позиция: ")
			fmt.Scanf("%s\n", &empl.Position)
			fmt.Print("Зарплата: ")
			fmt.Scanf("%d\n", &empl.Salary)
			for i := 0; i < size; i++ {
				if empls[i] == nil {
					empls[i] = empl
					count++
					fmt.Printf("Сотрудник %s добавлен!\n", empl.Name)
					break
				}
			}
		case 2:
			if count == 0 {
				fmt.Println("ERROR! Нет сотрудников - удаление невозможно")
				continue
			}
			fmt.Print("Введите номер сотрудника для удаления: ")
			var index int
			fmt.Scanf("%d\n", &index)
			if index < 1 || index > size {
				fmt.Println("ERROR! Невозможный номер")
				continue
			}
			if empls[index-1] == nil {
				fmt.Println("ERROR! Сотрудник с таким номером не нейден")
				continue
			}
			name := empls[index-1].Name
			empls[index-1] = nil
			count--
			fmt.Printf("Сотрудник %s удален!\n", name)
		case 3:
			if count == 0 {
				fmt.Println("ERROR! Сотрудники не найдены")
				continue
			}
			fmt.Println("\nСписок сотрудников:")
			fmt.Println("№   Имя\t Возраст\t Позиция\t Зарплата")
			fmt.Println("")
			for i := 0; i < size; i++ {
				if empls[i] != nil {
					fmt.Printf("%-3d %-10s %-15d %-15s %-10d\n", 
						i+1, empls[i].Name, empls[i].Age, empls[i].Position, empls[i].Salary)
				}
			}
			fmt.Printf("\nКоличество сотрудников: %d\n", count)
		case 4:
			fmt.Println("Выход")
			return
		default:
			fmt.Println("Неизвестная команда")
		}
		fmt.Println()
	}
}
