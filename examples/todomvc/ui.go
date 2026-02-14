package main

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/zodimo/go-compose/compose/foundation/layout/box"
	"github.com/zodimo/go-compose/compose/foundation/layout/column"
	"github.com/zodimo/go-compose/compose/foundation/lazy"
	"github.com/zodimo/go-compose/compose/foundation/text"
	"github.com/zodimo/go-compose/compose/material3/divider"
	"github.com/zodimo/go-compose/compose/material3/surface"
	"github.com/zodimo/go-compose/compose/ui/graphics"
	uiText "github.com/zodimo/go-compose/compose/ui/text"
	"github.com/zodimo/go-compose/compose/ui/unit"
	"github.com/zodimo/go-compose/modifiers/padding"
	"github.com/zodimo/go-compose/modifiers/size"
	"github.com/zodimo/go-compose/modifiers/weight"
	"github.com/zodimo/go-compose/pkg/api"
	"github.com/zodimo/go-compose/state"
)

// saveState saves the state to file (ignores errors for simplicity)
func saveState(state *TodoState) {
	_ = state.SaveToFile()
}

func UI(c api.Composer) api.Composer {
	// State management - load from file on first run
	todoStateValue := state.MustState[*TodoState](c, "todoState", func() *TodoState {
		return LoadFromFile()
	})

	// Input text state
	inputTextValue := c.State("inputText", func() any {
		return ""
	})
	inputText := inputTextValue.Get().(string)

	// Edit text state - keyed by todo ID
	editTextValue := c.State("editText", func() any {
		return ""
	})
	editText := editTextValue.Get().(string)

	// Todo list
	filteredTodos := todoStateValue.Get().FilteredTodos()
	currentState := todoStateValue.Get()

	// Background surface
	surface.Surface(
		column.Column(
			c.Sequence(
				// Header title
				box.Box(
					text.Text(
						"todos",
						text.WithTextStyleOptions(
							uiText.WithFontSize(unit.Sp(48)),
							uiText.WithColor(graphics.FromNRGBA(color.NRGBA{R: 175, G: 47, B: 47, A: 100})),
						),
					),
					box.WithModifier(padding.All(16)),
					box.WithAlignment(box.Center),
				),

				// Input section
				TodoInput(
					todoStateValue,
					inputText,
					func(newText string) {
						inputTextValue.Set(newText)
					},
					func() {
						// Get current value at click time, not composition time
						currentText := inputTextValue.Get().(string)
						currentText = strings.TrimSpace(currentText)
						if currentText == "" {
							return
						}
						newState := todoStateValue.Get().AddTodo(currentText)
						todoStateValue.Set(newState)
						inputTextValue.Set("")
						saveState(newState)
					},
				),

				divider.Divider(),
				box.Box(
					c.When(
						len(filteredTodos) > 0,
						lazy.LazyColumn(
							func(scope lazy.LazyListScope) {
								scope.Items(
									len(filteredTodos),
									func(index int) any {
										t := filteredTodos[index]
										isEditing := currentState.IsEditing(t.ID)
										return fmt.Sprintf("%d-%v", t.ID, isEditing)
									},
									func(index int) api.Composable {
										t := filteredTodos[index]
										isEditing := currentState.IsEditing(t.ID)

										return column.Column(
											c.Sequence(
												TodoItem(
													t,
													isEditing,
													editText,
													// onEditTextChange
													func(newText string) {
														editTextValue.Set(newText)
													},
													// onToggle
													func() {
														newState := todoStateValue.Get().ToggleTodo(t.ID)
														todoStateValue.Set(newState)
														saveState(newState)
													},
													// onEdit
													func() {
														newState := todoStateValue.Get().SetEditing(t.ID)
														todoStateValue.Set(newState)
														editTextValue.Set(t.Text)
													},
													// onSaveEdit
													func() {
														currentEditText := editTextValue.Get().(string)
														currentEditText = strings.TrimSpace(currentEditText)
														var newState *TodoState
														if currentEditText == "" {
															// Delete if empty
															newState = todoStateValue.Get().DeleteTodo(t.ID)
														} else {
															// Update text
															newState = todoStateValue.Get().UpdateTodo(t.ID, currentEditText)
														}
														newState = newState.CancelEditing()
														todoStateValue.Set(newState)
														editTextValue.Set("")
														saveState(newState)
													},
													// onCancelEdit
													func() {
														newState := todoStateValue.Get().CancelEditing()
														todoStateValue.Set(newState)
														editTextValue.Set("")
													},
													// onDelete
													func() {
														newState := todoStateValue.Get().DeleteTodo(t.ID)
														todoStateValue.Set(newState)
														saveState(newState)
													},
												),
												divider.Divider(),
											),
										)
									},
								)
							},
							lazy.WithModifier(size.FillMaxWidth()),
						),
					),
					box.WithModifier(weight.Weight(1)),
				),
				// Footer
				TodoFooter(todoStateValue),
			),
			column.WithModifier(size.FillMax()),
		),
		surface.WithColor(graphics.FromNRGBA(color.NRGBA{R: 245, G: 245, B: 245, A: 255})),
		surface.WithModifier(size.FillMax()),
	)(c)

	return c
}
