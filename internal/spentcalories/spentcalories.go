package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	training := strings.Split(data, ",")

	if len(training) != 3 {
		return 0, "", 0, fmt.Errorf("Неправильный формат данных: %s", data)
	}

	steps, err := strconv.Atoi(training[0])

	if err != nil {
		return 0, "", 0, fmt.Errorf("Неправильный формат шагов: %w", err)
	}

	duration, err := time.ParseDuration(training[2])

	if err != nil {
		return 0, "", 0, fmt.Errorf("Неверный формат времени: %w", err)
	}

	return steps, training[1], duration, nil
}

func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient

	return (float64(steps) * stepLength)/mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	distanceData := distance(steps, height)
	hours := duration.Hours()
	
	return distanceData / hours
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("Количество шагов не может быть меньше или равно нулю")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("Вес должен быть больше 0")
	}

	if height <= 0 {
		return 0, fmt.Errorf("Рост должен быть больше 0")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("Длительность должна быть больше 0")
	}

	averageSpeed := meanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	return (weight * averageSpeed * durationInMinutes) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}
