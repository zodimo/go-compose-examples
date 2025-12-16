package main

import (
	"image/color"

	"github.com/zodimo/go-compose/compose/foundation/layout/box"
	"github.com/zodimo/go-compose/compose/foundation/layout/column"
	"github.com/zodimo/go-compose/compose/foundation/lazy"
	"github.com/zodimo/go-compose/compose/foundation/material3/divider"
	"github.com/zodimo/go-compose/compose/foundation/material3/surface"
	"github.com/zodimo/go-compose/compose/foundation/text"
	"github.com/zodimo/go-compose/modifiers/padding"
	"github.com/zodimo/go-compose/modifiers/size"
	"github.com/zodimo/go-compose/pkg/api"
)

// func GetState[T any](mutableValue api.MutableValue) T {
// 	return mutableValue.Get().(T)
// }

func GetTodoState(todoStateValue api.MutableValue) *TodoState {
	return todoStateValue.Get().(*TodoState)
}

func UI(c api.Composer) api.Composer {
	// State management
	todoStateValue := c.State("todoState", func() any {
		return NewTodoState()
	})
	// state := todoStateValue.Get().(*TodoState)

	// Input text state
	inputTextValue := c.State("inputText", func() any {
		return ""
	})
	inputText := inputTextValue.Get().(string)

	// Background surface
	surface.Surface(
		func(c api.Composer) api.Composer {
			column.Column(
				func(c api.Composer) api.Composer {
					// Header title
					box.Box(
						text.Text(
							"todos",
							text.WithTextStyleOptions(
								text.StyleWithTextSize(48),
								text.StyleWithColor(color.NRGBA{R: 175, G: 47, B: 47, A: 100}),
							),
						),
						box.WithModifier(padding.All(16)),
						box.WithAlignment(box.Center),
					)(c)

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
							newState := GetTodoState(todoStateValue).AddTodo(currentText)
							todoStateValue.Set(newState)
							inputTextValue.Set("")
						},
					)(c)

					divider.Divider()(c)

					// Todo list
					filteredTodos := GetTodoState(todoStateValue).FilteredTodos()
					if len(filteredTodos) > 0 {
						lazy.LazyColumn(
							func(scope lazy.LazyListScope) {
								for _, todo := range filteredTodos {
									// Capture todo in closure
									t := todo
									scope.Item(t.ID, func(c api.Composer) api.Composer {
										TodoItem(
											t,
											func() {
												newState := GetTodoState(todoStateValue).ToggleTodo(t.ID)
												todoStateValue.Set(newState)
											},
											func() {
												newState := GetTodoState(todoStateValue).DeleteTodo(t.ID)
												todoStateValue.Set(newState)
											},
										)(c)
										divider.Divider()(c)
										return c
									})
								}
							},
							lazy.WithModifier(size.FillMaxWidth()),
						)(c)
					}

					// Footer
					TodoFooter(todoStateValue)(c)

					return c
				},
				column.WithModifier(size.FillMax()),
			)(c)

			return c
		},
		surface.WithColor(color.NRGBA{R: 245, G: 245, B: 245, A: 255}),
		surface.WithModifier(size.FillMax()),
	)(c)

	return c
}
