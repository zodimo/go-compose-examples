package main

import (
	"github.com/zodimo/go-compose/compose/foundation/canvas"
	"github.com/zodimo/go-compose/compose/ui"
	"github.com/zodimo/go-compose/compose/ui/geometry"
	"github.com/zodimo/go-compose/compose/ui/graphics"
	graphicsPath "github.com/zodimo/go-compose/compose/ui/graphics/path"
	uiInputPointer "github.com/zodimo/go-compose/compose/ui/input/pointer"
	"github.com/zodimo/go-compose/modifiers/pointer"
	"github.com/zodimo/go-compose/pkg/api"
)

func DrawingCanvas(
	paths []PathData,
	currentPath *PathData,
	onAction func(Action),
	actions Actions,
	modifier ui.Modifier,
) api.Composable {
	return func(c api.Composer) api.Composer {
		return canvas.Canvas(
			func(s graphics.DrawScope) {
				// Draw Path

				for _, pathData := range paths {
					path := PathDataToPath(pathData)
					s.DrawPath(path, pathData.Color, graphics.WithPathStyle(graphics.NewStroke(2)))
				}

				// Draw Current Path
				if currentPath != nil {
					path := PathDataToPath(*currentPath)

					s.DrawPath(path, currentPath.Color, graphics.WithPathStyle(graphics.NewStroke(2)))
				}

			},
			canvas.WithModifier(modifier.Then(
				pointer.PointerInput("canvas_pointer_input", func(scope uiInputPointer.PointerInputScope) {
					scope.DetectDragGestures(
						// onStart
						func(startPosition geometry.Offset) {
							onAction(actions.OnStartPath(startPosition))
						},
						// onDragEnd
						func() {
							onAction(actions.OnEndPath())
						},
						// onDragCancel
						func() {
							onAction(actions.OnDragCancel())
						},
						// onDrag
						func(change uiInputPointer.PointerInputChange, amount geometry.Offset) {
							onAction(actions.OnDraw(amount))
						},
					)
				}),
			)),
		)(c)
	}
}

func PathDataToPath(pathData PathData) graphics.Path {
	if len(pathData.Points) < 2 {
		return graphicsPath.New()
	}

	path := graphicsPath.New()
	startPoint := pathData.Points[0]
	path.MoveTo(startPoint.X(), startPoint.Y())
	for _, p := range pathData.Points[1:] {
		path.RelativeLineTo(p.X(), p.Y())
	}
	return path
}
