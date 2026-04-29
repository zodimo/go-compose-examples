package main

import (
	"image/color"

	"github.com/zodimo/go-compose/compose"
	"github.com/zodimo/go-compose/compose/foundation/layout/box"
	"github.com/zodimo/go-compose/compose/foundation/layout/column"
	"github.com/zodimo/go-compose/compose/foundation/layout/row"
	"github.com/zodimo/go-compose/compose/foundation/layout/spacer"
	"github.com/zodimo/go-compose/compose/foundation/text"
	"github.com/zodimo/go-compose/compose/material3"
	"github.com/zodimo/go-compose/compose/material3/button"
	"github.com/zodimo/go-compose/compose/material3/surface"
	"github.com/zodimo/go-compose/compose/material3/textfield"
	"github.com/zodimo/go-compose/compose/ui/graphics"
	uiText "github.com/zodimo/go-compose/compose/ui/text"
	"github.com/zodimo/go-compose/compose/ui/unit"
	"github.com/zodimo/go-compose/compose/viewmodel"
	"github.com/zodimo/go-compose/modifiers/padding"
	"github.com/zodimo/go-compose/modifiers/size"
	"github.com/zodimo/go-compose/pkg/api"
	"github.com/zodimo/go-compose/pkg/flow"
)

// validatePassword performs mock password validation
func validatePassword(password string) ValidationResult {
	if password == "" {
		return ValidationResult{Valid: false, Message: "Password is required"}
	}

	if len(password) < 8 {
		return ValidationResult{Valid: false, Message: "Password must be at least 8 characters"}
	}

	return ValidationResult{Valid: true, Message: ""}
}

func UI() compose.Composable {
	return func(c api.Composer) api.Composer {
		theme := material3.Theme(c)

		viewModel := viewmodel.RememberViewModel(c, func() *LoginViewModel {
			return NewViewModel()
		})

		viewState := flow.CollectStateFlowAsState(c, "ViewState", viewModel.AsStateFlow()).Get()

		// Background surface
		surface.Surface(
			box.Box(
				column.Column(
					c.Sequence(
						// App Title
						box.Box(
							text.Text(
								"GoCompose Login",
								text.WithTextStyleOptions(
									uiText.WithFontSize(unit.Sp(32)),
									uiText.WithColor(material3.Theme(c).ColorScheme().Primary),
								),
							),
							box.WithAlignment(box.Center),
						),
						spacer.Height(8),
						// Subtitle
						box.Box(
							text.Text(
								"Sign in to continue",
								text.WithTextStyleOptions(
									uiText.WithFontSize(unit.Sp(16)),
									uiText.WithColor(material3.Theme(c).ColorScheme().OnSurfaceVariant),
								),
							),
							box.WithAlignment(box.Center),
						),
						spacer.Height(32),
						// Show success message if logged in
						c.If(
							viewState.LoginSuccess(),
							c.Sequence(
								surface.Surface(
									box.Box(
										column.Column(
											c.Sequence(
												text.Text(
													"✓ Login Successful!",
													text.WithTextStyleOptions(
														uiText.WithFontSize(unit.Sp(20)),
														uiText.WithColor(graphics.FromNRGBA(color.NRGBA{R: 46, G: 125, B: 50, A: 255})),
													),
												),
												spacer.Height(8),
												text.Text(
													"Welcome back!",
													text.WithTextStyleOptions(
														uiText.WithColor(material3.Theme(c).ColorScheme().OnSurface),
													),
												),
											),
											column.WithAlignment(column.Middle),
										),
										box.WithModifier(padding.All(24)),
										box.WithAlignment(box.Center),
									),
									surface.WithColor(graphics.FromNRGBA(color.NRGBA{R: 232, G: 245, B: 233, A: 255})),
									surface.WithModifier(size.FillMaxWidth()),
								),
								spacer.Height(16),
								// Logout button
								box.Box(
									button.Outlined(
										func() {
											viewModel.OnAction(UIActions.OnLogout())
										},
										"Sign Out",
									),
									box.WithAlignment(box.Center),
								),
							),
							c.Sequence(
								// Email field
								textfield.TextField(
									viewState.EmailValue(),
									func(newValue string) {
										viewModel.OnAction(UIActions.OnEmailChange(newValue))
									},
									"Email",
									textfield.WithError(viewState.EmailHasError()),
									textfield.WithSupportingText(func() string {
										if viewState.EmailHasError() {
											return viewState.EmailError()
										}
										return "Enter your email address"
									}()),
									textfield.WithSingleLine(true),
									textfield.WithModifier(size.FillMaxWidth()),
								),
								spacer.Height(16),
								// Password field
								textfield.TextField(
									viewState.PasswordValue(),
									func(newValue string) {
										viewModel.OnAction(UIActions.OnPasswordChange(newValue))
									},
									"Password",
									textfield.WithError(viewState.PasswordHasError()),
									textfield.WithSupportingText(func() string {
										if viewState.PasswordHasError() {
											return viewState.PasswordError()
										}
										return "Minimum 8 characters"
									}()),
									textfield.WithSingleLine(true),
									textfield.WithOnSubmit(func() {
										// viewModel.OnAction(UIActions.OnSubmit())
									}),
									textfield.WithModifier(size.FillMaxWidth()),
								),
								spacer.Height(24),
								// Login button
								box.Box(
									button.Filled(
										func() {
											viewModel.OnAction(UIActions.OnSubmit())
										},
										"Sign In",
										button.WithModifier(size.Width(200)),
									),
									box.WithAlignment(box.Center),
								),
								spacer.Height(16),
								// Form status indicator
								c.When(
									viewState.Submitted() && !viewState.IsFormValid(),
									box.Box(
										text.Text(
											"Please fix the errors above",
											text.WithTextStyleOptions(
												uiText.WithColor(graphics.FromNRGBA(color.NRGBA{R: 176, G: 0, B: 32, A: 255})),
											),
										),
										box.WithAlignment(box.Center),
									),
								),
								spacer.Height(24),
								// Forgot password link (styled as text button)
								box.Box(
									row.Row(
										text.Text(
											"Forgot password?",
											text.WithTextStyleOptions(
												uiText.WithColor(material3.Theme(c).ColorScheme().Primary),
											),
										),
										row.WithAlignment(row.Middle),
									),
									box.WithAlignment(box.Center),
								),
								spacer.Height(16),
								// Sign up prompt
								box.Box(
									row.Row(
										c.Sequence(
											text.Text(
												"Don't have an account? ",
												text.WithTextStyleOptions(
													uiText.WithColor(material3.Theme(c).ColorScheme().OnSurfaceVariant), //colors.SurfaceRoles.OnVariant),
												),
											),
											text.Text(
												"Sign Up",
												text.WithTextStyleOptions(
													uiText.WithColor(material3.Theme(c).ColorScheme().Primary),
												),
											),
										),
										row.WithAlignment(row.Middle),
									),
									box.WithAlignment(box.Center),
								),
							),
						),
					),
					column.WithModifier(
						size.FillMaxWidth().
							Then(padding.Horizontal(32, 32)),
					),
				),
				box.WithModifier(size.FillMax()),
				box.WithAlignment(box.Center),
			),
			surface.WithColor(theme.ColorScheme().Surface),
			surface.WithModifier(size.FillMax()),
		)(c)

		return c
	}
}
