package service

type OrderValidator interface {
	Validate(number int) (int, error)
}
