package presentation

import "github.com/zodimo/go-compose/pkg/api"

type UiText interface {
	private()
	AsString() api.Composable
}

type DynamicText interface {
	UiText
}
type StringResourceId interface {
	UiText
}
