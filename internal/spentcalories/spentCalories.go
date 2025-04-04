package spentcalories

import (
	"time"
	"fmt"
	"errors"
	"strconv"
	"strings"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep = 0.65 // средняя длина шага.
	mInKm   = 1000 // количество метров в километре.
	minInH  = 60   // количество минут в часе.
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// ваш код ниже
dataParts := strings.Split(data, ",")
if len(dataParts) != 3{
	return 0,"",0, errors.New("Неправильный формат ввода, требуется три параметра")
}
steps, err := strconv.Atoi(dataParts[0])
if err != nil{
	return 0,"",0, fmt.Errorf("Неверное колличество шагов: %w", err)
} 
activity := dataParts[1]
duration, err := time.ParseDuration(dataParts[2])
if err != nil{
	return 0,"",0, fmt.Errorf("Неверный формат продолжительности: %w", err)
}
return steps, activity, duration, nil
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {
	// ваш код ниже
	distF := (float64(steps) * lenStep) / mInKm
	return distF
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	// ваш код ниже
	if duration <= 0{
		return 0
	}
	meanDist := distance(steps)
	hours:= duration.Hours()
	return meanDist / hours

}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {
	// ваш код ниже
steps, activity, duration, err := parseTraining(data)
if err != nil{
	return "Ошибка:" + err.Error()
}
switch activity {
case "Ходьба":
	
	distance := float64(steps) * lenStep / 1000 
		speed := meanSpeed(steps, duration)
		calories := WalkingSpentCalories(steps, weight, height, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activity, duration.Hours(), distance, speed, calories)

case "Бег":

		distance := float64(steps) * lenStep / 1000 
		speed := meanSpeed(steps, duration)
		calories := RunningSpentCalories(steps, weight, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activity, duration.Hours(), distance, speed, calories)
	default:
		return "Неизвестный тип тренировки"
}
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	// ваш код здесь
speed := meanSpeed(steps, duration)
calories := ((runningCaloriesMeanSpeedMultiplier * speed) - runningCaloriesMeanSpeedShift) * weight
if calories < 0{
	return 0
}
return calories
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	// ваш код здесь
speed := meanSpeed(steps, duration)
hours := duration.Hours()

calories := ((walkingCaloriesWeightMultiplier * weight) + (speed * speed / height) * walkingSpeedHeightMultiplier) * hours * minInH

if calories < 0 {
	return 0
}
return calories
}

