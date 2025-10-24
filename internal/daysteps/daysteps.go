package daysteps

import (
	"errors"
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
	// TODO: реализовать функцию

	walkInform := strings.Split(data, ",")

	if len(walkInform) != 2 {
		return 0, 0, errors.New("incorrect data length, should be 2")
	}

	stepsCount, errorStepsConvert := strconv.Atoi(walkInform[0])
	walkDuration, errorWalkDurationParse := time.ParseDuration(walkInform[1])

	if stepsCount <= 0 {
		return 0, 0, errors.New("incorrect steps count")
	}

	if walkDuration.Seconds() <= 0 {
		return 0, 0, errors.New("incorrect walk duration")
	}

	if errorStepsConvert != nil {
		return 0, 0, errorStepsConvert
	}

	if errorWalkDurationParse != nil {
		return 0, 0, errorWalkDurationParse
	}

	return stepsCount, walkDuration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
}
