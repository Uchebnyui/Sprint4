package daysteps

import (
        "errors"
        "fmt"
        "strconv"
        "strings"
        "time"

        "github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
        StepLength = 0.65 // длина шага в метрах
)

const mInKm = 1000 // количество метров в километре.

func parsePackage(data string) (int, time.Duration, error) {
        // ваш код ниже
        dataParts := strings.Split(data, ",")
        if len(dataParts) != 2 {
                return 0, 0, errors.New("incorrect data input format")
        }

        steps, err := strconv.Atoi(dataParts[0])
        if err != nil {
                return 0, 0, fmt.Errorf("incorrect steps value: %w", err)
        }

        if steps <= 0 {
                return 0, 0, errors.New("incorrect steps value, must be greater then zero")
        }

        duration, err := time.ParseDuration(dataParts[1])
        if err != nil {
                return 0, 0, fmt.Errorf("incorrect duration format: %w", err)
        }
        return steps, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
        // ваш код ниже
        steps, duration, err := parsePackage(data)
        if err != nil {
                fmt.Println("error :", err)
                return ""
        }
        if steps <= 0 {
                return ""
        }
        distMeter := float64(steps) * StepLength
        distKm := distMeter / mInKm
        calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
        result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distKm, calories)
        return result
}