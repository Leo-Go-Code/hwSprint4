package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var StepLength = 0.65 // длина шага в метрах

func parsePackage(data string) (int, time.Duration, error) {
	pieces := strings.Split(data, ",")
	if len(pieces) != 2 {
		return 0, 0, errors.New("Incorrect data received")
	}

	steps, err1 := strconv.Atoi(pieces[0])
	if err1 != nil {
		return 0, 0, errors.New("Error converting string to number")
	} else if steps < 1 {
		return 0, 0, errors.New("The number of steps must be greater than 0")
	}

	duration, err2 := time.ParseDuration(pieces[1])
	if err2 != nil {
		return 0, 0, errors.New("Error converting string to time.Duration")
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
	steps, duration, errPP := parsePackage(data)
	if errPP != nil {
		return ""
	} else if steps < 1 {
		return ""
	}

	distance := StepLength * float64(steps) / 1_000

	kcal := spentcalories.WalkingSpentCalories(int(distance), weight, height, duration)

	return fmt.Sprintf("Количество шагов: %d. \nДистанция составила %.2f км. \nВы сожгли %.2f ккал.", steps, distance, kcal)
}
