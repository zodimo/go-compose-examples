package main

import (
	"github.com/zodimo/go-compose/compose"
	"github.com/zodimo/go-compose/compose/foundation/text"
	"github.com/zodimo/go-compose/pkg/api"
)

func UI() compose.Composable {
	return func(c api.Composer) api.Composer {
		return text.Text("hello world")(c)
	}
}
