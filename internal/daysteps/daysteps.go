package daysteps

import (
	"errors"
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
	steps, _, errorDataParse := parsePackage(data)

	if errorDataParse != nil {
		fmt.Println(errorDataParse)
		return ""
		// в тз написано: "Проверить, чтобы количество шагов было больше 0. В противном случае вернуть пустую строку.",
		// но если я сделаю тут проверку на колчество шагов, она никогда не отработает в случае если шагов будет 0,
		// потому что такая проверка есть в функции parsePackage, и она возвращает ошибку в случае если шагов 0 или меньше
		// соответственно errorDataParse != nil будет true и функция отобразит ошибку и вернет пустую строку.
	}
	distanseInM := stepLength * float64(steps)
	distanseInKm := distanseInM / mInKm
	caloriesTotal := 0.00

	res := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanseInKm, caloriesTotal)

	return res
}
