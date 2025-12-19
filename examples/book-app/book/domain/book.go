package domain

import (
	"github.com/zodimo/go-maybe"
)

type Book struct {
	Id              string
	Title           string
	ImageUrl        string
	Authors         []string
	Description     string
	Languages       []string
	FirtPublishYear maybe.Maybe[string]
	AverageRating   maybe.Maybe[float64]
	RatingCount     maybe.Maybe[int]
	PageCount       maybe.Maybe[int]
	EditionCount    int
}
