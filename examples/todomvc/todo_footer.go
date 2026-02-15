package main

import (
	"fmt"

	"github.com/zodimo/go-compose/compose/foundation/layout/row"
	"github.com/zodimo/go-compose/compose/foundation/layout/spacer"
	"github.com/zodimo/go-compose/compose/foundation/text"
	"github.com/zodimo/go-compose/compose/material3/button"
	uiText "github.com/zodimo/go-compose/compose/ui/text"
	"github.com/zodimo/go-compose/compose/ui/unit"
	"github.com/zodimo/go-compose/modifiers/padding"
	"github.com/zodimo/go-compose/modifiers/size"
	"github.com/zodimo/go-compose/pkg/api"
	"github.com/zodimo/go-compose/state"
)

// TodoFooter renders the footer with item count, filter buttons, and clear completed.
func TodoFooter(todoStateValue state.MutableValueTyped[*TodoState]) api.Composable {
	return func(c api.Composer) api.Composer {
		state := todoStateValue.Get()
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
						uiText.WithFontSize(unit.Sp(12)),
					),
				)(c)

				// Spacer
				spacer.Weight(1)(c)

				// Filter buttons - use c.If to get different keys for selected vs unselected
				c.If(
					state.Filter == FilterAll,
					button.Filled(func() {
						newState := todoStateValue.Get().SetFilter(FilterAll)
						todoStateValue.Set(newState)
					}, "All"),
					button.Text(func() {
						newState := todoStateValue.Get().SetFilter(FilterAll)
						todoStateValue.Set(newState)
					}, "All"),
				)(c)
				c.If(
					state.Filter == FilterActive,
					button.Filled(func() {
						newState := todoStateValue.Get().SetFilter(FilterActive)
						todoStateValue.Set(newState)
					}, "Active"),
					button.Text(func() {
						newState := todoStateValue.Get().SetFilter(FilterActive)
						todoStateValue.Set(newState)
					}, "Active"),
				)(c)
				c.If(
					state.Filter == FilterCompleted,
					button.Filled(func() {
						newState := todoStateValue.Get().SetFilter(FilterCompleted)
						todoStateValue.Set(newState)
					}, "Completed"),
					button.Text(func() {
						newState := todoStateValue.Get().SetFilter(FilterCompleted)
						todoStateValue.Set(newState)
					}, "Completed"),
				)(c)

				// Spacer
				spacer.Weight(1)(c)

				// Clear completed button
				if state.CompletedCount() > 0 {
					button.Text(func() {
						newState := todoStateValue.Get().ClearCompleted()
						todoStateValue.Set(newState)
						_ = newState.SaveToFile()
					}, "Clear completed")(c)
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
