package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/zodimo/go-compose/compose/ui/geometry"
	"github.com/zodimo/go-compose/compose/ui/graphics"
	"github.com/zodimo/go-compose/compose/viewmodel"
	"github.com/zodimo/go-compose/pkg/flow"
)

var AllColors = []graphics.Color{
	graphics.ColorWhite,
	graphics.ColorBlack,
	graphics.ColorRed,
	graphics.ColorGreen,
	graphics.ColorBlue,
}

type PathData struct {
	ID     string
	Points []geometry.Offset
	Color  graphics.Color
}

type DrawingViewModel struct {
	viewmodel.ViewModel
	mutableState *flow.MutableStateFlow[*DrawingState]
}

func NewDrawingViewModel() *DrawingViewModel {
	return &DrawingViewModel{
		mutableState: flow.NewMutableStateFlow(&DrawingState{
			SelectedColor: graphics.ColorBlack,
		}),
	}
}

func (vm *DrawingViewModel) OnInit() {}

func (vm *DrawingViewModel) AsStateFlow() flow.StateFlow[*DrawingState] {
	return vm.mutableState.AsStateFlow()
}

func (vm *DrawingViewModel) OnAction(action Action) {
	switch action := action.(type) {
	case *onSelectColorAction:
		vm.handleSelectColor(action)
	case *onStartPathAction:
		vm.handleStartPath(action)
	case *onDrawAction:
		vm.handleDraw(action)
	case *onDragCancelAction:
		vm.handleDragCancel()
	case *onEndPathAction:
		vm.handleEndPath()
	case *onClearCanvasClickAction:
		vm.handleClearCanvasClick()
	default:
		panic(fmt.Sprintf("InstanceViewModel: unknown action %T", action))
	}
}

func (vm *DrawingViewModel) handleSelectColor(action *onSelectColorAction) {
	vm.mutableState.Update(func(current *DrawingState) *DrawingState {
		return current.Copy(WithSelectedColor(action.color))
	})
}

func (vm *DrawingViewModel) handleStartPath(action *onStartPathAction) {
	vm.mutableState.Update(func(current *DrawingState) *DrawingState {
		return current.Copy(WithCurrentPath(&PathData{
			ID:     uuid.New().String(),
			Points: []geometry.Offset{action.point},
			Color:  current.SelectedColor,
		}))
	})

}

func (vm *DrawingViewModel) handleDraw(action *onDrawAction) {
	vm.mutableState.Update(func(current *DrawingState) *DrawingState {
		if current.CurrentPath == nil {
			return current
		}
		newPathData := PathData{
			ID:     current.CurrentPath.ID,
			Points: append(current.CurrentPath.Points, action.point),
			Color:  current.CurrentPath.Color,
		}
		return current.Copy(WithCurrentPath(&newPathData))
	})
}

func (vm *DrawingViewModel) handleDragCancel() {
	vm.mutableState.Update(func(current *DrawingState) *DrawingState {
		return current.Copy(WithCurrentPath(nil))
	})
}

func (vm *DrawingViewModel) handleEndPath() {
	vm.mutableState.Update(func(current *DrawingState) *DrawingState {
		if current.CurrentPath == nil {
			return current
		}
		newPaths := append(current.Paths, *current.CurrentPath)
		return current.Copy(WithPaths(newPaths), WithCurrentPath(nil))
	})
}

func (vm *DrawingViewModel) handleClearCanvasClick() {
	vm.mutableState.Update(func(current *DrawingState) *DrawingState {
		return current.Copy(WithPaths([]PathData{}), WithCurrentPath(nil))
	})
}
