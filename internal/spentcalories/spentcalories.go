package spentcalories

import (
	"errors"
	"fmt"
	"log"
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
	// TODO: done
	dataSlice := strings.Split(data, ",")

	if len(dataSlice) != 3 {
		return 0, "0", 0, errors.New("неправильно количество данных")
	}

	amountOfStep, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return 0, "0", 0, err
	}

	activityType := dataSlice[1]
	activityDuration, err := time.ParseDuration(dataSlice[2])
	if err != nil {
		return 0, "0", 0, err
	}

	if amountOfStep <= 0 || activityDuration <= 0 {
		return 0, "0", 0, errors.New("некорректный ввод данных")
	}

	return amountOfStep, activityType, activityDuration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: done
	return ((height * stepLengthCoefficient) * float64(steps)) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: done
	if duration <= 0 {
		return 0
	}

	total := distance(steps, height)

	return total / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	var info string

	steps, activityType, activityDur, err := parseTraining(data)
	if err != nil {
		log.Println(err)
	}

	switch activityType {
	case "Бег":
		runCal, err := RunningSpentCalories(steps, weight, height, activityDur)
		if err != nil {
			return "", err
		}

		meanSP := meanSpeed(steps, height, activityDur)
		dist := distance(steps, height)

		info = fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`, activityType, activityDur.Hours(), dist, meanSP, runCal)
		break
	case "Ходьба":
		walkCal, err := WalkingSpentCalories(steps, weight, height, activityDur)
		if err != nil {
			return "", err
		}

		meanSP := meanSpeed(steps, height, activityDur)
		dist := distance(steps, height)

		info = fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`, activityType, activityDur.Hours(), dist, meanSP, walkCal)
		break
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	return info, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: done
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("неправильные данные")
	}

	meanSP := meanSpeed(steps, height, duration)
	durationInMin := duration.Minutes()

	return (weight * meanSP * durationInMin) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: done
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("неправильные данные")
	}

	meanSP := meanSpeed(steps, height, duration)
	durationInMin := duration.Minutes()

	return ((weight * meanSP * durationInMin) / minInH) * walkingCaloriesCoefficient, nil
}
