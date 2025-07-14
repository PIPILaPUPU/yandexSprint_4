package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

var (
	errWrongDurationData = errors.New("неправильные данные о продолжительности")
	errWrongSteps        = errors.New("неправильные данные о шагах")
	errNotEnoughData     = errors.New("недостаточно данных")
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: done
	dataSlice := strings.Split(data, ",")
	if len(dataSlice) != 2 {
		return 0, 0, errNotEnoughData
	}

	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, errWrongSteps
	}

	activityDur, err := time.ParseDuration(dataSlice[1])
	if err != nil {
		return 0, 0, err
	}
	if activityDur <= 0 {
		return 0, 0, errWrongDurationData
	}

	return steps, activityDur, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, activityDur, err := parsePackage(data)
	if errors.Is(err, errWrongSteps) {
		log.Println(err)
		return ""
	}

	if errors.Is(err, errWrongDurationData) {
		log.Println(err)
		return ""
	}

	distance := float64(steps) * stepLength / mInKm
	dayCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, activityDur)
	if err != nil {
		log.Println(err)
		return ""
	}

	info := fmt.Sprintf(`Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.
`, steps, distance, dayCalories)

	return info
}
