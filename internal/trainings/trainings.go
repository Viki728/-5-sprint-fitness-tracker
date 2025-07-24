package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	//Разделяем строку на слайс строк и проверяем, чтобы длина слайса была равна 3
	slice := strings.Split(datastring, ",")
	if len(slice) != 3 {
		return errors.New("slice is not equal to 3")
	}

	//Преобразуем первый и третий элементы слайсы и обрабатываем ошибки
	st, err := strconv.Atoi(slice[0])
	if err != nil {
		return err
	}
	t.Steps = st

	dur, err := time.ParseDuration(slice[2])
	if err != nil {
		return err
	}
	t.Duration = dur
	//Сохраняем значение типа тренировки в поле TrainingType
	t.TrainingType = slice[1]

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// Вычисляем дистанцию, среднюю скорость
	dist := spentenergy.Distance(t.Steps, t.Height)
	v := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	//Проверяем вид тренировки и рассчитываем калории для каждого вида
	if t.TrainingType == "Ходьба" {
		calor, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), dist, v, calor), nil
	}

	if t.TrainingType == "Бег" {
		calor, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), dist, v, calor), nil
	}
	return "", errors.New("неизвестный тип тренировки")
}
