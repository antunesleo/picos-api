package domain

type SpotRepository interface {
	ListAll() []Spot
}
