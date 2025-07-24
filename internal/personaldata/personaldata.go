package personaldata

import "fmt"

type Personal struct {
	//Создаем поля структуры
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	//Выводим данные структуры на экран
	fmt.Println("Имя: ", p.Name)
	fmt.Println("Вес: ", p.Weight)
	fmt.Println("Рост: ", p.Height)
}
