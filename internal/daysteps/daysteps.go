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
		return 0, 0, errors.New("некорректная длинна значений, должно быть 3")
	}

	stepsCount, err := strconv.Atoi(walkInform[0])

	if err != nil {
		return 0, 0, err
	}

	if stepsCount <= 0 {
		return 0, 0, errors.New("некорректное значение шагов")
	}

	walkDuration, err := time.ParseDuration(walkInform[1])

	if err != nil {
		return 0, 0, err
	}

	if walkDuration.Seconds() <= 0 {
		return 0, 0, errors.New("некорректное значение времени")
	}

	return stepsCount, walkDuration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, walkTime, err := parsePackage(data)

	if err != nil {
		log.Println(err)
		return ""
	}

	distanseInM := stepLength * float64(steps)
	distanseInKm := distanseInM / mInKm
	caloriesTotal, err := spentcalories.WalkingSpentCalories(steps, weight, height, walkTime)

	if err != nil {
		log.Println(err)
		return ""
	}

	res := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanseInKm, caloriesTotal)

	return res
}
