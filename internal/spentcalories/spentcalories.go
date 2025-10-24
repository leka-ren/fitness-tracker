package spentcalories

import (
	"errors"
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
	// "3456,Ходьба,3h00m"
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
	// Функция принимает количество шагов steps, рост пользователя height и продолжительность активности duration  и возвращает среднюю скорость.
	if duration <= 0 {
		return 0
	}

	totalDistance := distance(steps, height)

	midTime := totalDistance / duration.Hours()

	return midTime
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
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
