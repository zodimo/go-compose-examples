package main

import (
	"github.com/zodimo/go-compose/compose/foundation/layout/column"
	"github.com/zodimo/go-compose/compose/material3/text"
	"github.com/zodimo/go-compose/compose/viewmodel"
	"github.com/zodimo/go-compose/modifiers/size"
	"github.com/zodimo/go-compose/pkg/api"
	"github.com/zodimo/go-compose/pkg/flow"
)

func UI() api.Composable {
	return func(c api.Composer) api.Composer {

		model := viewmodel.RememberViewModel(c, func() *DrawingViewModel {
			return NewDrawingViewModel()
		})

		state := flow.CollectStateFlowAsState(c, "DrawingState", model.AsStateFlow()).Get()
		actions := NewActions()

		return column.Column(
			c.Sequence(
				text.HeadlineMedium("Canvas Demo"),
				DrawingCanvas(
					state.Paths,
					state.CurrentPath,
					func(action Action) { model.OnAction(action) },
					actions,
					size.FillMax(),
				),
			),
			column.WithSpacing(column.SpaceSides),
			column.WithAlignment(column.Middle),
			column.WithModifier(size.FillMax()),
		)(c)
	}
}
