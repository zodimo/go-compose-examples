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
	"github.com/zodimo/go-compose/state"
)

// TodoFooter renders the footer with item count, filter buttons, and clear completed.
func TodoFooter(todoStateValue state.MutableValue) api.Composable {
	return func(c api.Composer) api.Composer {
		state := GetTodoState(todoStateValue)
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

				// Filter buttons - use c.If to get different keys for selected vs unselected
				c.If(
					state.Filter == FilterAll,
					button.Filled(func() {
						newState := GetTodoState(todoStateValue).SetFilter(FilterAll)
						todoStateValue.Set(newState)
					}, "All"),
					button.Text(func() {
						newState := GetTodoState(todoStateValue).SetFilter(FilterAll)
						todoStateValue.Set(newState)
					}, "All"),
				)(c)
				c.If(
					state.Filter == FilterActive,
					button.Filled(func() {
						newState := GetTodoState(todoStateValue).SetFilter(FilterActive)
						todoStateValue.Set(newState)
					}, "Active"),
					button.Text(func() {
						newState := GetTodoState(todoStateValue).SetFilter(FilterActive)
						todoStateValue.Set(newState)
					}, "Active"),
				)(c)
				c.If(
					state.Filter == FilterCompleted,
					button.Filled(func() {
						newState := GetTodoState(todoStateValue).SetFilter(FilterCompleted)
						todoStateValue.Set(newState)
					}, "Completed"),
					button.Text(func() {
						newState := GetTodoState(todoStateValue).SetFilter(FilterCompleted)
						todoStateValue.Set(newState)
					}, "Completed"),
				)(c)

				// Spacer
				row.Row(
					func(c api.Composer) api.Composer { return c },
					row.WithModifier(weight.Weight(1)),
				)(c)

				// Clear completed button
				if state.CompletedCount() > 0 {
					button.Text(func() {
						newState := GetTodoState(todoStateValue).ClearCompleted()
						todoStateValue.Set(newState)
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
