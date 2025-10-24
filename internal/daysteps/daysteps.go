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

func parsePackage(data string) (int, time.Duration, error) {
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
	steps, walkTime, errorDataParse := parsePackage(data)

	if errorDataParse != nil {
		log.Println(errorDataParse)
		return ""

	}

	distanseInM := stepLength * float64(steps)
	distanseInKm := distanseInM / mInKm
	caloriesTotal, errCaloriesCalc := spentcalories.WalkingSpentCalories(steps, weight, height, walkTime)

	if errCaloriesCalc != nil {
		log.Println(errCaloriesCalc)
		return ""
	}

	res := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanseInKm, caloriesTotal)

	return res
}
