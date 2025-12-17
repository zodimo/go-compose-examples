package main

import (
	"image/color"
	"regexp"
	"strings"

	"github.com/zodimo/go-compose/compose/foundation/layout/box"
	"github.com/zodimo/go-compose/compose/foundation/layout/column"
	"github.com/zodimo/go-compose/compose/foundation/layout/row"
	"github.com/zodimo/go-compose/compose/foundation/layout/spacer"
	"github.com/zodimo/go-compose/compose/foundation/material3/button"
	"github.com/zodimo/go-compose/compose/foundation/material3/surface"
	"github.com/zodimo/go-compose/compose/foundation/material3/textfield"
	"github.com/zodimo/go-compose/compose/foundation/text"
	"github.com/zodimo/go-compose/modifiers/padding"
	"github.com/zodimo/go-compose/modifiers/size"
	"github.com/zodimo/go-compose/pkg/api"
	"github.com/zodimo/go-compose/theme"
)

// ValidationResult represents the result of a field validation
type ValidationResult struct {
	Valid   bool
	Message string
}

// validateEmail performs mock email validation
func validateEmail(email string) ValidationResult {
	email = strings.TrimSpace(email)
	if email == "" {
		return ValidationResult{Valid: false, Message: "Email is required"}
	}

	// Simple email regex for validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return ValidationResult{Valid: false, Message: "Please enter a valid email address"}
	}

	return ValidationResult{Valid: true, Message: ""}
}

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

func UI(c api.Composer) api.Composer {
	colors := theme.ColorHelper.ColorSelector()

	// State for form fields
	emailValue := c.State("email", func() any { return "" })
	passwordValue := c.State("password", func() any { return "" })

	// State for validation errors (only shown after interaction)
	emailTouched := c.State("emailTouched", func() any { return false })
	passwordTouched := c.State("passwordTouched", func() any { return false })

	// State for form submission
	submitted := c.State("submitted", func() any { return false })
	loginSuccess := c.State("loginSuccess", func() any { return false })

	// Get current values
	email := emailValue.Get().(string)
	password := passwordValue.Get().(string)
	isEmailTouched := emailTouched.Get().(bool)
	isPasswordTouched := passwordTouched.Get().(bool)
	isSubmitted := submitted.Get().(bool)
	isLoginSuccess := loginSuccess.Get().(bool)

	// Validate fields
	emailValidation := validateEmail(email)
	passwordValidation := validatePassword(password)

	// Show errors only after touch or submission attempt
	showEmailError := (isEmailTouched || isSubmitted) && !emailValidation.Valid
	showPasswordError := (isPasswordTouched || isSubmitted) && !passwordValidation.Valid

	// Check if form is valid for submission
	isFormValid := emailValidation.Valid && passwordValidation.Valid

	// Mock login handler
	handleLogin := func() {
		submitted.Set(true)

		// Re-validate at submission time
		currentEmail := strings.TrimSpace(emailValue.Get().(string))
		currentPassword := passwordValue.Get().(string)

		emailVal := validateEmail(currentEmail)
		passwordVal := validatePassword(currentPassword)

		if emailVal.Valid && passwordVal.Valid {
			// Mock successful login
			// In a real app, this would make an API call
			loginSuccess.Set(true)
		}
	}

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
								text.StyleWithTextSize(32),
								text.StyleWithColor(colors.PrimaryRoles.Primary),
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
								text.StyleWithTextSize(16),
								text.StyleWithColor(colors.SurfaceRoles.OnVariant),
							),
						),
						box.WithAlignment(box.Center),
					),
					spacer.Height(32),
					// Show success message if logged in
					c.If(
						isLoginSuccess,
						c.Sequence(
							surface.Surface(
								box.Box(
									column.Column(
										c.Sequence(
											text.Text(
												"✓ Login Successful!",
												text.WithTextStyleOptions(
													text.StyleWithTextSize(20),
													text.StyleWithColor(theme.ColorHelper.SpecificColor(color.NRGBA{R: 46, G: 125, B: 50, A: 255})),
												),
											),
											spacer.Height(8),
											text.Text(
												"Welcome back!",
												text.WithTextStyleOptions(
													text.StyleWithColor(colors.SurfaceRoles.OnSurface),
												),
											),
										),
										column.WithAlignment(column.Middle),
									),
									box.WithModifier(padding.All(24)),
									box.WithAlignment(box.Center),
								),
								surface.WithColor(theme.ColorHelper.SpecificColor(color.NRGBA{R: 232, G: 245, B: 233, A: 255})),
								surface.WithModifier(size.FillMaxWidth()),
							),
							spacer.Height(16),
							// Logout button
							box.Box(
								button.Outlined(
									func() {
										// Reset form
										emailValue.Set("")
										passwordValue.Set("")
										emailTouched.Set(false)
										passwordTouched.Set(false)
										submitted.Set(false)
										loginSuccess.Set(false)
									},
									"Sign Out",
								),
								box.WithAlignment(box.Center),
							),
						),
						c.Sequence(
							// Email field
							textfield.TextField(
								email,
								func(newValue string) {
									emailValue.Set(newValue)
									if !isEmailTouched {
										emailTouched.Set(true)
									}
								},
								"Email",
								textfield.WithError(showEmailError),
								textfield.WithSupportingText(func() string {
									if showEmailError {
										return emailValidation.Message
									}
									return "Enter your email address"
								}()),
								textfield.WithSingleLine(true),
								textfield.WithModifier(size.FillMaxWidth()),
							),
							spacer.Height(16),
							// Password field
							textfield.TextField(
								password,
								func(newValue string) {
									passwordValue.Set(newValue)
									if !isPasswordTouched {
										passwordTouched.Set(true)
									}
								},
								"Password",
								textfield.WithError(showPasswordError),
								textfield.WithSupportingText(func() string {
									if showPasswordError {
										return passwordValidation.Message
									}
									return "Minimum 8 characters"
								}()),
								textfield.WithSingleLine(true),
								textfield.WithOnSubmit(handleLogin),
								textfield.WithModifier(size.FillMaxWidth()),
							),
							spacer.Height(24),
							// Login button
							box.Box(
								button.Filled(
									handleLogin,
									"Sign In",
									button.WithModifier(size.Width(200)),
								),
								box.WithAlignment(box.Center),
							),
							spacer.Height(16),
							// Form status indicator
							c.When(
								isSubmitted && !isFormValid,
								box.Box(
									text.Text(
										"Please fix the errors above",
										text.WithTextStyleOptions(
											text.StyleWithColor(theme.ColorHelper.SpecificColor(color.NRGBA{R: 176, G: 0, B: 32, A: 255})),
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
											text.StyleWithColor(colors.PrimaryRoles.Primary),
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
												text.StyleWithColor(colors.SurfaceRoles.OnVariant),
											),
										),
										text.Text(
											"Sign Up",
											text.WithTextStyleOptions(
												text.StyleWithColor(colors.PrimaryRoles.Primary),
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
		surface.WithColor(colors.SurfaceRoles.Surface),
		surface.WithModifier(size.FillMax()),
	)(c)

	return c
}
