package main

import (
	"image/color"

	"github.com/zodimo/go-compose/compose/foundation/layout/box"
	"github.com/zodimo/go-compose/compose/foundation/layout/row"
	"github.com/zodimo/go-compose/compose/foundation/text"
	"github.com/zodimo/go-compose/compose/material3/checkbox"
	"github.com/zodimo/go-compose/compose/material3/iconbutton"
	"github.com/zodimo/go-compose/compose/material3/textfield"

	"github.com/zodimo/go-compose/compose/ui/graphics"
	uiText "github.com/zodimo/go-compose/compose/ui/text"
	"github.com/zodimo/go-compose/modifiers/padding"
	"github.com/zodimo/go-compose/modifiers/size"
	"github.com/zodimo/go-compose/modifiers/weight"
	"github.com/zodimo/go-compose/pkg/api"

	"golang.org/x/exp/shiny/materialdesign/icons"
)

// TodoItem renders a single todo item with checkbox, text, edit/delete buttons.
// When isEditing is true, shows an editable text field instead of the text.
func TodoItem(
	todo Todo,
	isEditing bool,
	editText string,
	onEditTextChange func(string),
	onToggle func(),
	onEdit func(),
	onSaveEdit func(),
	onCancelEdit func(),
	onDelete func(),
) api.Composable {
	return func(c api.Composer) api.Composer {
		// Editing mode - show TextField
		return c.If(
			isEditing,
			row.Row(
				c.Sequence(
					// Editing text field
					textfield.TextField(
						editText,
						onEditTextChange,
						"Edit todo",
						textfield.WithSingleLine(true),
						textfield.WithOnSubmit(func() {
							onSaveEdit()
						}),
						textfield.WithModifier(
							weight.Weight(1).
								Then(padding.Horizontal(8, 8)),
						),
					),
					// Save button
					iconbutton.Standard(
						onSaveEdit,
						icons.ContentSave,
						"Save",
						iconbutton.WithModifier(padding.All(4)),
					),
					// Cancel button
					iconbutton.Standard(
						onCancelEdit,
						icons.NavigationClose,
						"Cancel",
						iconbutton.WithModifier(padding.All(4)),
					),
				),
				row.WithAlignment(row.Middle),
				row.WithModifier(
					size.FillMaxWidth().
						Then(padding.Vertical(4, 4)).
						Then(padding.Horizontal(8, 8)),
				),
			),
			row.Row(
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
							// Completed: gray + strikethrough
							text.Text(
								todo.Text,
								text.WithTextStyleOptions(
									uiText.WithColor(graphics.FromNRGBA(color.NRGBA{R: 150, G: 150, B: 150, A: 255})),
								),
								text.StyleWithStrikethrough(),
							),
							// Active: normal color
							text.Text(
								todo.Text,
								text.WithTextStyleOptions(
									uiText.WithColor(graphics.FromNRGBA(color.NRGBA{R: 50, G: 50, B: 50, A: 255})),
								),
							),
						),
						box.WithModifier(
							weight.Weight(1).
								Then(padding.Horizontal(8, 8)),
						),
						box.WithAlignment(box.W),
					),
					// Edit button
					iconbutton.Standard(
						onEdit,
						icons.EditorModeEdit,
						"Edit todo",
						iconbutton.WithModifier(padding.All(4)),
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
			),
		)(c)

	}

}
