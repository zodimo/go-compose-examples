package booklist

import "gitub.com/zodimo/go-compose-examples/examples/book-app/book/domain"

type BookListAction interface {
	private()
}

var _ BookListAction = (*OnBookClick)(nil)

type OnBookClick struct {
	BookId string
}

func (a *OnBookClick) private() {}

type OnSearchQueryChange struct {
	Query string
}

func (a *OnSearchQueryChange) private() {}

type OnTabSelected struct {
	TabIndex int
}

func (a *OnTabSelected) private() {}

type Actions interface {
	OnSearchQueryChange(query string) BookListAction
	OnBookClick(book domain.Book) BookListAction
	OnTabSelected(index int) BookListAction
}

var _ Actions = (*actions)(nil)

type actions struct {
}

func NewActions() Actions {
	return &actions{}
}

func (a *actions) OnSearchQueryChange(query string) BookListAction {
	return &OnSearchQueryChange{Query: query}
}

func (a *actions) OnBookClick(book domain.Book) BookListAction {
	return &OnBookClick{BookId: book.Id}
}

func (a *actions) OnTabSelected(index int) BookListAction {
	return &OnTabSelected{TabIndex: index}
}
