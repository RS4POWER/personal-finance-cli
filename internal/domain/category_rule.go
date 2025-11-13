package domain

type CategoryRule struct {
	ID       int64
	Pattern  string // regex pattern
	Category string
}
