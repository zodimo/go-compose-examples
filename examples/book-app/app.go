package main

import (
	"image/color"

	"github.com/zodimo/go-compose/compose"
	"github.com/zodimo/go-compose/compose/foundation/layout/column"
	"github.com/zodimo/go-compose/compose/ui/graphics"
	"github.com/zodimo/go-compose/modifiers/background"
	"github.com/zodimo/go-compose/modifiers/size"
	"gitub.com/zodimo/go-compose-examples/examples/book-app/book/domain"
	"gitub.com/zodimo/go-compose-examples/examples/book-app/book/presentation/booklist"
)

func App() compose.Composable {
	return func(c compose.Composer) compose.Composer {
		viewModel := booklist.NewBookListViewModel()
		return column.Column(
			c.Sequence(
				booklist.BookListScreenRoot(viewModel, func(book domain.Book) {

				}),
			),

			column.WithModifier(
				background.Background(
					graphics.FromNRGBA(color.NRGBA{R: 0, G: 0, B: 200, A: 200})).
					Then(size.FillMaxHeight()),
			),
		)(c)
	}
}
