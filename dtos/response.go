package dtos

type Response struct {
	Status  int
	Message string
	Ok      bool
	Data    interface{}
}
