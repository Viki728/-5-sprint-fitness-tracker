package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	//Создаем поля
	Steps                 int
	Duration              time.Duration
	personaldata.Personal //встроенная структура
	// Personal из пакета personaldata, у которой есть метод Print().
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	//Разделяем строку на слайс строк
	sl := strings.Split(datastring, ",")
	if len(sl) != 2 {
		return errors.New("slice length is not equal to 2")
	}
	// Преобразуем первый элемент в int, второй элемент второй к типу time.Duration
	st, err := strconv.Atoi(sl[0])
	if err != nil {
		return err
	}
	dur, err := time.ParseDuration(sl[1])
	if err != nil {
		return err
	}

	if dur <= 0 {
		return errors.New("negative duration")
	}
	//Присваиваем полям структуры DaySteps полученные значения
	ds.Steps = st
	ds.Duration = dur
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	//Вычисляем дистанцию
	dist := spentenergy.Distance(ds.Steps, ds.Height)

	//Вычисляем количество сожженных калорий
	cal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, cal), nil
}
