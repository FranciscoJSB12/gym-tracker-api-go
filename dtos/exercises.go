package dtos

type ExerciseResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	RoutineID      uint   `json:"routineID"`
	ProposedRounds uint   `json:"proposedRounds"`
	Repetitions    uint   `json:"repetitions"`
}
