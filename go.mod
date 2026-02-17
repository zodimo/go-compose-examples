module gitub.com/zodimo/go-compose-examples

go 1.25.0

require (
	gioui.org v0.9.0
	github.com/zodimo/go-compose v0.1.94
	github.com/zodimo/go-maybe v0.1.7
	golang.org/x/exp/shiny v0.0.0-20260212183809-81e46e3db34a
)

require (
	gioui.org/shader v1.0.8 // indirect
	git.sr.ht/~schnwalter/gio-mw v0.0.0-20250713180710-9d8d98474447 // indirect
	github.com/go-text/typesetting v0.3.3 // indirect
	github.com/zodimo/go-lazy v0.1.1 // indirect
	github.com/zodimo/go-ternary v0.2.0 // indirect
	github.com/zodimo/go-zero-hash v0.1.0 // indirect
	golang.org/x/image v0.36.0 // indirect
	golang.org/x/sys v0.41.0 // indirect
	golang.org/x/text v0.34.0 // indirect
)

// replace github.com/zodimo/go-compose => ../go-compose

// v0.3.3 has a bug with single line text rendering
replace github.com/go-text/typesetting => github.com/zodimo/typesetting v0.3.4-0.20260209162200-1565df70b998
