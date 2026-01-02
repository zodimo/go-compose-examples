package booklist

import (
	"image/color"

	"github.com/zodimo/go-compose/compose/ui/graphics"
	"github.com/zodimo/go-compose/modifiers/background"
	"github.com/zodimo/go-compose/pkg/api"
	"github.com/zodimo/go-compose/theme"
	"gitub.com/zodimo/go-compose-examples/examples/book-app/book/domain"
	"gitub.com/zodimo/go-compose-examples/examples/book-app/book/presentation/booklist/components"
)

func BookListScreenRoot(viewModel *BookListViewModel, onBookClick func(book domain.Book)) api.Composable {
	state := viewModel.mutableState.AsStateFlow()
	return bookListScreen(state.Value(), func(action BookListAction) {
		viewModel.OnAction(action)
	})
}

func bookListScreen(state BoolListState, onAction func(action BookListAction)) api.Composable {
	actions := NewActions()
	return func(c api.Composer) api.Composer {
		return components.SearchBar(
			state.SearchQuery,

			//on search query change
			func(query string) {
				onAction(actions.OnSearchQueryChange(query))
			},
			//ime search
			func() {

			},
			background.Background(
				theme.ColorHelper.SpecificColor(
					graphics.FromNRGBA(color.NRGBA{R: 255, G: 255, B: 255, A: 255})),
			),
		)(c)
	}
}
