package main

import (
	"github.com/zodimo/go-compose/compose/foundation/layout/row"
	"github.com/zodimo/go-compose/compose/material3/button"
	"github.com/zodimo/go-compose/compose/material3/checkbox"
	"github.com/zodimo/go-compose/compose/material3/textfield"
	"github.com/zodimo/go-compose/modifiers/padding"
	"github.com/zodimo/go-compose/modifiers/size"
	"github.com/zodimo/go-compose/modifiers/weight"
	"github.com/zodimo/go-compose/pkg/api"
	"github.com/zodimo/go-compose/store"
)

// TodoInput renders the header with toggle-all checkbox and new todo input field.
func TodoInput(
	todoStateValue store.TypedMutableValueInterface[*TodoState],
	inputText string,
	onInputChange func(string),
	onSubmit func(),
) api.Composable {
	return func(c api.Composer) api.Composer {
		row.Row(
			c.Sequence(
				// Toggle all checkbox (only show if there are todos)
				c.When(len(todoStateValue.Get().Todos) > 0,
					checkbox.Checkbox(
						todoStateValue.Get().AllCompleted(),
						func(checked bool) {
							newstate := todoStateValue.Get().ToggleAll(checked)
							todoStateValue.Set(newstate)
						},
						checkbox.WithModifier(padding.All(8)),
					),
				),
				// New todo text field
				textfield.TextField(
					inputText,
					onInputChange,
					"What needs to be done?",
					textfield.WithSingleLine(true),
					textfield.WithOnSubmit(func() {
						if inputText != "" {
							onSubmit()
						}
					}),
					textfield.WithModifier(
						weight.Weight(1).
							Then(padding.Horizontal(8, 8)),
					),
				),
				// Add button to submit (TextField doesn't have onSubmit)
				button.Filled(
					func() {
						if inputText != "" {
							onSubmit()
						}
					},
					"Add",
					button.WithModifier(padding.All(4)),
				),
			),

			row.WithAlignment(row.Middle),
			row.WithModifier(
				size.FillMaxWidth().
					Then(padding.All(8)),
			),
		)(c)

		return c
	}
}
