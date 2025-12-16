package main

import (
	"fmt"

	"github.com/zodimo/go-compose/compose/foundation/layout/row"
	"github.com/zodimo/go-compose/compose/foundation/material3/button"
	"github.com/zodimo/go-compose/compose/foundation/text"
	"github.com/zodimo/go-compose/modifiers/padding"
	"github.com/zodimo/go-compose/modifiers/size"
	"github.com/zodimo/go-compose/modifiers/weight"
	"github.com/zodimo/go-compose/pkg/api"
)

// TodoFooter renders the footer with item count, filter buttons, and clear completed.
func TodoFooter(
	state *TodoState,
	onFilterChange func(Filter),
	onClearCompleted func(),
) api.Composable {
	return func(c api.Composer) api.Composer {
		if len(state.Todos) == 0 {
			return c
		}

		row.Row(
			func(c api.Composer) api.Composer {
				// Items left count
				activeCount := state.ActiveCount()
				itemText := "items"
				if activeCount == 1 {
					itemText = "item"
				}
				text.Text(
					fmt.Sprintf("%d %s left", activeCount, itemText),
					text.WithTextStyleOptions(
						text.StyleWithTextSize(12),
					),
				)(c)

				// Spacer
				row.Row(
					func(c api.Composer) api.Composer { return c },
					row.WithModifier(weight.Weight(1)),
				)(c)

				// Filter buttons
				filterButton := func(label string, filter Filter) api.Composable {
					if state.Filter == filter {
						return button.Filled(func() { onFilterChange(filter) }, label)
					}
					return button.Text(func() { onFilterChange(filter) }, label)
				}

				filterButton("All", FilterAll)(c)
				filterButton("Active", FilterActive)(c)
				filterButton("Completed", FilterCompleted)(c)

				// Spacer
				row.Row(
					func(c api.Composer) api.Composer { return c },
					row.WithModifier(weight.Weight(1)),
				)(c)

				// Clear completed button
				if state.CompletedCount() > 0 {
					button.Text(onClearCompleted, "Clear completed")(c)
				}

				return c
			},
			row.WithAlignment(row.Middle),
			row.WithModifier(
				size.FillMaxWidth().
					Then(padding.All(12)),
			),
		)(c)

		return c
	}
}
