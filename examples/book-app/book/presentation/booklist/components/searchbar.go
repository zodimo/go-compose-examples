package components

import (
	"github.com/zodimo/go-compose/compose/foundation/material3/textfield"
	"github.com/zodimo/go-compose/compose/ui"
	"github.com/zodimo/go-compose/modifiers/size"
	"github.com/zodimo/go-compose/pkg/api"
)

func SearchBar(
	searchQuery string,
	onSearchQueryChange func(query string),
	onImeSearch func(),
	modifier ui.Modifier,

) api.Composable {
	return func(c api.Composer) api.Composer {
		return textfield.Outlined(
			searchQuery,
			onSearchQueryChange,
			"",
			textfield.WithModifier(
				modifier.Then(
					size.FillMaxWidth(),
				),
			))(c)
	}
}
