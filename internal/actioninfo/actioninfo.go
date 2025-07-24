package actioninfo

import "fmt"

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	//Перебираем значения сдайса
	for _, v := range dataset {
		//Парсим значения, при ощибке логируем ее
		err := dp.Parse(v)
		if err != nil {
			fmt.Println("ошибка парсинга:", err)
			continue
		}
		//Формируем и выводим строку с инфо-ей об активности, при ошибке логируем ее
		act, err := dp.ActionInfo()
		if err != nil {
			fmt.Println("ошибка получения информации об активности:", err)
			continue
		}
		fmt.Println(act)
	}
}
