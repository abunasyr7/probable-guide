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
		return 0, "", 0, fmt.Errorf("invalid data format: %s", data)
	}

	steps, err := strconv.Atoi(training[0])

	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid steps format: %w", err)
	}

	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("steps must be positive: %d", steps)
	}

	duration, err := time.ParseDuration(training[2])

	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid duration format: %w", err)
	}

	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("duration must be positive: %v", duration)
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


func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("steps must be positive")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("weight must be positive")
	}

	if height <= 0 {
		return 0, fmt.Errorf("height must be positive")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("duration must be positive")
	}

	averageSpeed := meanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	return (weight * averageSpeed * durationInMinutes) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
		if steps <= 0 {
		return 0, fmt.Errorf("steps must be positive")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("weight must be positive")
	}

	if height <= 0 {
		return 0, fmt.Errorf("height must be positive")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("duration must be positive")
	}

	averageSpeed := meanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	return ((weight * averageSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient, nil

}


func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, training, duration, err := parseTraining(data) 

	if err != nil {
		return "", fmt.Errorf("error: %w", err)
	}

	if steps <= 0 {
		return "", fmt.Errorf("steps must be positive")
	}

	if duration <= 0 {
		return "", fmt.Errorf("duration must be positive")
	}

	var calories float64

	distance := distance(steps, height)
	averageSpeed := meanSpeed(steps, height, duration)

	switch training {
	case "Ходьба": 
		calories, err = WalkingSpentCalories(steps, weight, height, duration)

		if err != nil {
			return "", fmt.Errorf("something went wrong")
		}
	

	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration) 

		if err != nil {
			return "", fmt.Errorf("something went wrong")
		}

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", training, duration.Hours(), distance, averageSpeed, calories), nil
}