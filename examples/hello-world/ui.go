package main

import (
	"github.com/zodimo/go-compose/compose/foundation/text"
	"github.com/zodimo/go-compose/pkg/api"
)

func UI(c api.Composer) api.Composer {
	return text.Text("hello world")(c)
}
