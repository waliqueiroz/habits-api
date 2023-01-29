package rest

import (
	"time"

	"github.com/waliqueiroz/habits-api/internal/domain"
)

type HabitRequestDTO struct {
	Title    string `json:"title"`
	Weekdays []int  `json:"weekdays"`
}

func mapHabitRequestToDomain(habit HabitRequestDTO) domain.Habit {
	return domain.NewHabit(habit.Title, habit.Weekdays)
}

type HabitResponseDTO struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func mapHabitFromDomain(habit domain.Habit) HabitResponseDTO {
	return HabitResponseDTO{
		ID:        habit.ID,
		Title:     habit.Title,
		CreatedAt: habit.CreatedAt,
	}
}

func mapHabitsFromDomain(habits []domain.Habit) []HabitResponseDTO {
	list := make([]HabitResponseDTO, len(habits))

	for i, habit := range habits {
		list[i] = mapHabitFromDomain(habit)
	}

	return list
}

type DayResumeDTO struct {
	PossibleHabits  []HabitResponseDTO `json:"possible_habits"`
	CompletedHabits []string           `json:"completed_habits"`
}

func mapDayResumeFromDomain(dayResume domain.DayResume) DayResumeDTO {
	completedHabits := make([]string, len(dayResume.CompletedHabits))

	for i, completedHabit := range dayResume.CompletedHabits {
		completedHabits[i] = completedHabit.ID
	}

	return DayResumeDTO{
		PossibleHabits:  mapHabitsFromDomain(dayResume.PossibleHabits),
		CompletedHabits: completedHabits,
	}
}

type DailySummary struct {
	ID        string    `json:"id"`
	Date      time.Time `json:"date"`
	Completed int       `json:"completed"`
	Amount    int       `json:"amount"`
}

func mapDailySummaryFromDomain(dailySummary domain.DailySummary) DailySummary {
	return DailySummary{
		ID:        dailySummary.ID,
		Date:      dailySummary.Date,
		Completed: dailySummary.Completed,
		Amount:    dailySummary.Amount,
	}
}

func mapDailySummariesFromDomain(dailySummaries []domain.DailySummary) []DailySummary {
	list := make([]DailySummary, len(dailySummaries))

	for i, habit := range dailySummaries {
		list[i] = mapDailySummaryFromDomain(habit)
	}

	return list
}
