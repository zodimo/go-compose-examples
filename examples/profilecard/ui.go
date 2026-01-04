package main

import (
	"embed"
	_ "image/png"

	fImage "github.com/zodimo/go-compose/compose/foundation/image"
	"github.com/zodimo/go-compose/compose/foundation/layout/box"
	"github.com/zodimo/go-compose/compose/foundation/layout/column"
	"github.com/zodimo/go-compose/compose/foundation/layout/row"
	"github.com/zodimo/go-compose/compose/foundation/layout/spacer"
	ftext "github.com/zodimo/go-compose/compose/foundation/text"
	"github.com/zodimo/go-compose/compose/material3"
	"github.com/zodimo/go-compose/compose/material3/button"
	"github.com/zodimo/go-compose/compose/material3/card"
	"github.com/zodimo/go-compose/compose/material3/iconbutton"
	"github.com/zodimo/go-compose/compose/material3/surface"
	"github.com/zodimo/go-compose/compose/ui/graphics"
	"github.com/zodimo/go-compose/compose/ui/graphics/shape"
	uilayout "github.com/zodimo/go-compose/compose/ui/layout"
	"github.com/zodimo/go-compose/compose/ui/unit"
	"github.com/zodimo/go-compose/modifiers/border"
	"github.com/zodimo/go-compose/modifiers/clip"
	"github.com/zodimo/go-compose/modifiers/offset"
	"github.com/zodimo/go-compose/modifiers/size"
	"github.com/zodimo/go-compose/pkg/api"

	uiText "github.com/zodimo/go-compose/compose/ui/text"
	mdicons "golang.org/x/exp/shiny/materialdesign/icons"
)

//go:embed gopher.png
var assets embed.FS

func UI(c api.Composer) api.Composer {
	theme := material3.Theme(c)

	// Load profile image
	imageResource := graphics.NewResourceFromImageFS(assets, "gopher.png")

	// Icon data for social buttons
	iconShare := mdicons.SocialShare
	iconPublic := mdicons.SocialPublic
	iconPhoto := mdicons.ImagePhoto
	iconVideo := mdicons.AVPlayCircleOutline

	// Profile card using Card component
	card.Elevated(
		card.CardContents(
			// Header as ContentCover (full width, no padding)
			card.ContentCover(
				surface.Surface(
					spacer.Height(100),
					surface.WithColor(theme.ColorScheme().Primary),
				),
			),

			// Main content with profile image, text, buttons
			card.Content(
				column.Column(
					c.Sequence(
						// Profile image (overlapping header with offset)
						box.Box(
							fImage.Image(
								imageResource,
								fImage.WithContentScale(uilayout.ContentScaleCrop),
								fImage.WithAlignment(size.Center),
								fImage.WithModifier(
									size.Size(100, 100).
										Then(clip.Clip(shape.CircleShape)).
										Then(border.Border(4, theme.ColorScheme().Primary, shape.CircleShape)),
								),
							),
							box.WithAlignment(box.Center),
							box.WithModifier(size.FillMaxWidth().Then(offset.OffsetY(-50))),
						),

						// Username (offset up to account for overlapping image)
						box.Box(
							ftext.Text(
								"CodingLab",
								ftext.WithTextStyleOptions(
									uiText.WithFontSize(unit.Sp(24)),
									uiText.WithColor(material3.Theme(c).ColorScheme().OnSurface),
								),
							),
							box.WithAlignment(box.Center),
							box.WithModifier(size.FillMaxWidth().Then(offset.OffsetY(-40))),
						),

						// Subtitle
						box.Box(
							ftext.Text(
								"YouTuber & Blogger",
								ftext.WithTextStyleOptions(
									uiText.WithFontSize(unit.Sp(14)),
									uiText.WithColor(material3.Theme(c).ColorScheme().OnSurfaceVariant),
								),
							),
							box.WithAlignment(box.Center),
							box.WithModifier(size.FillMaxWidth().Then(offset.OffsetY(-35))),
						),

						spacer.Height(20),

						// Social media icons row
						box.Box(
							row.Row(
								c.Sequence(
									iconbutton.FilledTonal(func() {}, iconShare, "Facebook"),
									spacer.Width(12),
									iconbutton.FilledTonal(func() {}, iconPublic, "Twitter"),
									spacer.Width(12),
									iconbutton.FilledTonal(func() {}, iconPhoto, "Instagram"),
									spacer.Width(12),
									iconbutton.FilledTonal(func() {}, iconVideo, "YouTube"),
								),
								row.WithAlignment(row.Middle),
							),
							box.WithAlignment(box.Center),
							box.WithModifier(size.FillMaxWidth().Then(offset.OffsetY(-25))),
						),

						spacer.Height(4),

						// Subscribe and Message buttons
						box.Box(
							row.Row(
								c.Sequence(
									button.Filled(func() {}, "Subscribe",
										button.WithModifier(size.Width(120)),
									),
									spacer.Width(16),
									button.Outlined(func() {}, "Message",
										button.WithModifier(size.Width(120)),
									),
								),
								row.WithAlignment(row.Middle),
							),
							box.WithAlignment(box.Center),
							box.WithModifier(size.FillMaxWidth().Then(offset.OffsetY(-15))),
						),

						spacer.Height(8),

						// Stats row
						box.Box(
							row.Row(
								c.Sequence(
									ftext.Text("♡ 60.4k", ftext.WithTextStyleOptions(
										uiText.WithFontSize(unit.Sp(14)),
										uiText.WithColor(material3.Theme(c).ColorScheme().OnSurfaceVariant),
									)),
									spacer.Width(24),
									ftext.Text("◎ 20k", ftext.WithTextStyleOptions(
										uiText.WithFontSize(unit.Sp(14)),
										uiText.WithColor(material3.Theme(c).ColorScheme().OnSurfaceVariant),
									)),
									spacer.Width(24),
									ftext.Text("⇄ 12.4k", ftext.WithTextStyleOptions(
										uiText.WithFontSize(unit.Sp(14)),
										uiText.WithColor(material3.Theme(c).ColorScheme().OnSurfaceVariant),
									)),
								),
								row.WithAlignment(row.Middle),
							),
							box.WithAlignment(box.Center),
							box.WithModifier(size.FillMaxWidth()),
						),
					),
					column.WithModifier(size.FillMaxWidth()),
				),
			),
		),
		card.WithModifier(size.FillMax()),
	)(c)

	return c
}
