package booklist

import (
	"github.com/zodimo/go-maybe"
	"gitub.com/zodimo/go-compose-examples/examples/book-app/book/domain"
	"gitub.com/zodimo/go-compose-examples/examples/book-app/core/presentation"
)

type BoolListState struct {
	SearchQuery       string
	SearchBookResults []domain.Book
	FavoriteBooks     []domain.Book
	IsLoading         bool
	SelectedTabIndex  int
	ErrorMessage      maybe.Maybe[presentation.UiText]
}

func NewBookListState() BoolListState {
	return BoolListState{
		SearchQuery:       "",
		SearchBookResults: []domain.Book{},
		FavoriteBooks:     []domain.Book{},
		IsLoading:         false,
		SelectedTabIndex:  0,
		ErrorMessage:      maybe.None[presentation.UiText](),
	}
}
