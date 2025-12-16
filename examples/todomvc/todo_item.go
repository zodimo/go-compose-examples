package main

import (
	"image/color"

	"github.com/zodimo/go-compose/compose/foundation/layout/box"
	"github.com/zodimo/go-compose/compose/foundation/layout/row"
	"github.com/zodimo/go-compose/compose/foundation/material3/checkbox"
	"github.com/zodimo/go-compose/compose/foundation/material3/iconbutton"
	"github.com/zodimo/go-compose/compose/foundation/text"
	"github.com/zodimo/go-compose/modifiers/padding"
	"github.com/zodimo/go-compose/modifiers/size"
	"github.com/zodimo/go-compose/modifiers/weight"
	"github.com/zodimo/go-compose/pkg/api"

	"golang.org/x/exp/shiny/materialdesign/icons"
)

// TodoItem renders a single todo item with checkbox, text, and delete button.
func TodoItem(
	todo Todo,
	onToggle func(),
	onDelete func(),
) api.Composable {
	return func(c api.Composer) api.Composer {
		return row.Row(
			c.Sequence(
				// Completion checkbox
				checkbox.Checkbox(
					todo.Completed,
					func(checked bool) {
						onToggle()
					},
					checkbox.WithModifier(padding.All(4)),
				),
				// Todo text
				box.Box(
					c.If(
						todo.Completed,
						text.Text(
							todo.Text,
							text.WithTextStyleOptions(
								text.StyleWithColor(color.NRGBA{R: 150, G: 150, B: 150, A: 255}),
							),
						),
						text.Text(
							todo.Text,
							text.WithTextStyleOptions(
								text.StyleWithColor(color.NRGBA{R: 50, G: 50, B: 50, A: 255}),
							),
						),
					),
					box.WithModifier(
						weight.Weight(1).
							Then(padding.Horizontal(8, 8)),
					),
					box.WithAlignment(box.W),
				),
				// Delete button
				iconbutton.Standard(
					onDelete,
					icons.ActionDelete,
					"Delete todo",
					iconbutton.WithModifier(padding.All(4)),
				),
			),
			row.WithAlignment(row.Middle),
			row.WithModifier(
				size.FillMaxWidth().
					Then(padding.Vertical(4, 4)).
					Then(padding.Horizontal(8, 8)),
			),
		)(c)
	}
}
