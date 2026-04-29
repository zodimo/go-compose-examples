package main

import "github.com/zodimo/go-compose/compose/ui/graphics"

type DrawingState struct {
	SelectedColor graphics.Color
	CurrentPath   *PathData
	Paths         []PathData
}

func NewDrawingState(
	selectedColor graphics.Color,
	currentPath *PathData,
	paths []PathData,
) *DrawingState {
	return &DrawingState{
		SelectedColor: selectedColor,
		CurrentPath:   currentPath,
		Paths:         paths,
	}
}

type DrawingStateCopyOption func(s *DrawingState)

func WithSelectedColor(c graphics.Color) DrawingStateCopyOption {
	return func(s *DrawingState) {
		s.SelectedColor = c
	}
}

func WithCurrentPath(p *PathData) DrawingStateCopyOption {
	return func(s *DrawingState) {
		s.CurrentPath = p
	}
}

func WithPaths(p []PathData) DrawingStateCopyOption {
	return func(s *DrawingState) {
		s.Paths = p
	}
}

func (s *DrawingState) Copy(options ...DrawingStateCopyOption) *DrawingState {
	newState := DrawingState{
		SelectedColor: s.SelectedColor,
		CurrentPath:   s.CurrentPath,
		Paths:         s.Paths,
	}
	for _, opt := range options {
		if opt != nil {
			opt(&newState)
		}
	}

	return &newState
}
