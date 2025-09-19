package dtos

type ExerciseResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	RoutineID      uint   `json:"routineId"`
	ProposedRounds int64  `json:"proposedRounds"`
}
