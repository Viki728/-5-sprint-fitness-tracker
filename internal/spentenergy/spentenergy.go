package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	//Проверка входных параметров
	if steps <= 0 {
		return 0, errors.New("incorrect number of steps")
	}
	if duration <= 0 {
		return 0, errors.New("incorrect number of duration")
	}
	//Рассчитываем и возвращаем количество калорий
	calor := ((MeanSpeed(steps, height, duration) * duration.Minutes() * weight) / minInH) * walkingCaloriesCoefficient
	return calor, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	//Проверка входных параметров
	if steps <= 0 {
		return 0, errors.New("incorrect number of steps")
	}
	if duration <= 0 {
		return 0, errors.New("incorrect number of duration")
	}
	//Рассчитываем и возвращаем количество калорий
	calor := (MeanSpeed(steps, height, duration) * duration.Minutes() * weight) / minInH
	return calor, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Проверка отрицательных значений (шагов и продолжительности)
	if steps <= 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}
	//Вычисляем и возвращаем среднюю скорость
	v := Distance(steps, height) / (duration.Hours())
	return v
}

func Distance(steps int, height float64) float64 {
	//Вычисляем дистанцию в км
	lenStep := ((height * stepLengthCoefficient) * float64(steps)) / float64(mInKm)
	return lenStep
}
