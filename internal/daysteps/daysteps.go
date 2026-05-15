package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")

	if len(parts) < 2 {
		return  0, 0, fmt.Errorf("invalid format: %s", data)
	}

	num, err := strconv.Atoi(parts[0])

	if err != nil {
		return 0, 0, fmt.Errorf("invalid int: %w", err)
	}

	duration, err := time.ParseDuration(parts[1])

	if err != nil {
		return 0, 0, fmt.Errorf("invalid duration %w", err)
	}

	return num, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
}
