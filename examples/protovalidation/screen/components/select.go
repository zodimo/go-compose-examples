package components

import (
	"github.com/zodimo/go-compose/pkg/api"
	uiv1 "gitub.com/zodimo/go-compose-examples/gen/ui/v1"
)

func SelectInput(selectInput *uiv1.SelectInput, onSelect func(string)) api.Composable {
	return func(c api.Composer) api.Composer {
		return c
	}
}
