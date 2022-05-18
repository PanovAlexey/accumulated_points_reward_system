package service

type OrderValidator interface {
	Validate(number int64) error
}
