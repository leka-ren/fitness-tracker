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
	trainingInform := strings.Split(data, ",")

	if len(trainingInform) != 3 {
		return 0, "", 0, errors.New("incorrect data length, should be 3")
	}

	stepsCount, errorStepsConvert := strconv.Atoi(trainingInform[0])
	activityType := trainingInform[1]
	walkDuration, errorWalkDurationParse := time.ParseDuration(trainingInform[2])

	if stepsCount <= 0 {
		return 0, "", 0, errors.New("incorrect steps count")
	}

	if walkDuration.Seconds() <= 0 {
		return 0, "", 0, errors.New("incorrect walk duration")
	}

	if errorStepsConvert != nil {
		return 0, "", 0, errorStepsConvert
	}

	if errorWalkDurationParse != nil {
		return 0, "", 0, errorWalkDurationParse
	}

	return stepsCount, activityType, walkDuration, nil
}

func distance(steps int, height float64) float64 {
	stepLength := stepLengthCoefficient * height
	walkDistanceInM := float64(steps) * stepLength
	walkDistanceInKm := walkDistanceInM / mInKm

	return walkDistanceInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	totalDistance := distance(steps, height)

	midTime := totalDistance / duration.Hours()

	return midTime
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activityType, totalTime, parseError := parseTraining(data)

	if parseError != nil {
		log.Println(parseError)
		return "", parseError
	}

	var caloriesSpentTotal float64
	var err error
	switch activityType {
	case "Ходьба":
		caloriesSpentTotal, err = WalkingSpentCalories(steps, weight, height, totalTime)
		if err != nil {
			return "", err
		}
	case "Бег":
		caloriesSpentTotal, err = RunningSpentCalories(steps, weight, height, totalTime)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	res := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activityType, totalTime.Hours(), distance(steps, height), meanSpeed(steps, height, totalTime), caloriesSpentTotal)

	return res, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("incorrect steps count")
	}
	if weight <= 0 {
		return 0, errors.New("incorrect weight")
	}
	if height <= 0 {
		return 0, errors.New("incorrect height")
	}
	if duration <= 0 {
		return 0, errors.New("incorrect duration")
	}

	midManSpeed := meanSpeed(steps, height, duration)
	caloriesSpentTotal := (weight * midManSpeed * duration.Minutes()) / minInH

	return caloriesSpentTotal, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("incorrect steps count")
	}
	if weight <= 0 {
		return 0, errors.New("incorrect weight")
	}
	if height <= 0 {
		return 0, errors.New("incorrect height")
	}
	if duration <= 0 {
		return 0, errors.New("incorrect duration")
	}

	midManSpeed := meanSpeed(steps, height, duration)
	caloriesSpentTotal := (weight * midManSpeed * duration.Minutes()) / minInH

	// корректировка под пеший тип активности
	caloriesSpentTotal *= walkingCaloriesCoefficient

	return caloriesSpentTotal, nil
}
