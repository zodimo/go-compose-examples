package main

import (
	"github.com/zodimo/go-compose/compose/ui/geometry"
	"github.com/zodimo/go-compose/compose/ui/graphics"
)

type Actions interface {
	OnSelectColor(c graphics.Color) Action
	OnStartPath(p geometry.Offset) Action
	OnDraw(p geometry.Offset) Action
	OnDragCancel() Action
	OnEndPath() Action
	OnClearCanvasClick() Action
}

type Action interface {
	isAction() // Marker interface
}

type actions struct{}

func NewActions() Actions {
	return &actions{}
}

func (a *actions) OnSelectColor(c graphics.Color) Action {
	return &onSelectColorAction{color: c}
}

func (a *actions) OnStartPath(p geometry.Offset) Action {
	return &onStartPathAction{point: p}
}

func (a *actions) OnDraw(p geometry.Offset) Action {
	return &onDrawAction{point: p}
}

func (a *actions) OnDragCancel() Action {
	return &onDragCancelAction{}
}

func (a *actions) OnEndPath() Action {
	return &onEndPathAction{}
}

func (a *actions) OnClearCanvasClick() Action {
	return &onClearCanvasClickAction{}
}

type onSelectColorAction struct {
	color graphics.Color
}

type onStartPathAction struct {
	point geometry.Offset
}

type onDrawAction struct {
	point geometry.Offset
}

type onDragCancelAction struct{}

type onEndPathAction struct{}

type onClearCanvasClickAction struct{}

func (a *onSelectColorAction) isAction()      {}
func (a *onStartPathAction) isAction()        {}
func (a *onDrawAction) isAction()             {}
func (a *onDragCancelAction) isAction()       {}
func (a *onEndPathAction) isAction()          {}
func (a *onClearCanvasClickAction) isAction() {}
