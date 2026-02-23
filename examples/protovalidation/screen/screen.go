package screen

import (
	"fmt"

	"github.com/zodimo/go-compose/compose"
	"github.com/zodimo/go-compose/compose/foundation/layout/box"
	"github.com/zodimo/go-compose/compose/foundation/layout/column"
	"github.com/zodimo/go-compose/compose/foundation/layout/spacer"
	"github.com/zodimo/go-compose/compose/foundation/text"
	"github.com/zodimo/go-compose/compose/material3"
	"github.com/zodimo/go-compose/compose/material3/button"
	"github.com/zodimo/go-compose/compose/material3/surface"
	"github.com/zodimo/go-compose/compose/material3/textfield"
	uiText "github.com/zodimo/go-compose/compose/ui/text"
	"github.com/zodimo/go-compose/compose/ui/unit"
	"github.com/zodimo/go-compose/modifiers/padding"
	"github.com/zodimo/go-compose/modifiers/size"
	"github.com/zodimo/go-compose/pkg/api"
	"github.com/zodimo/go-compose/pkg/flow"
	"github.com/zodimo/go-compose/state"
	"gitub.com/zodimo/go-compose-examples/examples/protovalidation/form"
)

func ScreenRoot(viewModel *ViewModel) compose.Composable {
	return func(c api.Composer) api.Composer {
		//Wrapping the Stateflow and the composable in a box is required for the state to be collected
		return box.Box(func(c api.Composer) api.Composer {
			//This is where the composer can react to the changes in the state
			modelState := flow.CollectStateFlowAsState(c, "LoginState", viewModel.AsStateFlow())
			actions := NewActions()
			return Screen(modelState.Get(), actions, func(action Action) {
				viewModel.OnAction(action)
			})(c)
		})(c)
	}
}

func Screen(viewState *ViewState, actions Actions, onAction func(Action)) compose.Composable {
	return func(c api.Composer) api.Composer {
		theme := material3.Theme(c)

		formStateValue := state.MustRemember(c, "form-state", func() form.FormState {
			return form.NewFormState()
		})

		formState := formStateValue.Get()

		person := viewState.User()

		ageStr := ""
		if person.Age > 0 {
			ageStr = fmt.Sprintf("%d", person.Age)
		}

		// Background surface
		surface.Surface(
			box.Box(
				column.Column(
					c.Sequence(
						// App Title
						box.Box(
							text.Text(
								"Registration Form",
								text.WithTextStyleOptions(
									uiText.WithFontSize(unit.Sp(32)),
									uiText.WithColor(theme.ColorScheme().Primary),
								),
							),
							box.WithAlignment(box.Center),
						),
						spacer.Height(8),
						// Subtitle
						box.Box(
							text.Text(
								"Please fill in the form below",
								text.WithTextStyleOptions(
									uiText.WithFontSize(unit.Sp(16)),
									uiText.WithColor(theme.ColorScheme().OnSurfaceVariant),
								),
							),
							box.WithAlignment(box.Center),
						),
						spacer.Height(32),
						c.If(viewState.SubmittedSuccessfully(),
							box.Box(
								text.Text(
									"Form Submitted Successfully!",
									text.WithTextStyleOptions(
										uiText.WithFontSize(unit.Sp(20)),
										uiText.WithColor(theme.ColorScheme().Primary),
									),
								),
								box.WithAlignment(box.Center),
							),
							c.Sequence(
								textfield.Filled(
									person.Name,
									func(value string) {
										onAction(actions.FieldOnChange("name", value))
									},
									textfield.WithLabel("Name"),
									textfield.WithError(formState.GetError("name") != ""),
									textfield.WithSupportingText(formState.GetError("name")),
									textfield.WithSingleLine(true),
									textfield.WithModifier(size.FillMaxWidth()),
								),
								spacer.Height(16),
								textfield.Filled(
									person.Email,
									func(newValue string) {
										onAction(actions.FieldOnChange("email", newValue))
									},
									textfield.WithLabel("Email"),
									textfield.WithError(formState.GetError("email") != ""),
									textfield.WithSupportingText(formState.GetError("email")),
									textfield.WithSingleLine(true),
									textfield.WithModifier(size.FillMaxWidth()),
								),
								spacer.Height(16),
								textfield.Filled(
									ageStr,
									func(newValue string) {
										onAction(actions.FieldOnChange("age", newValue))
									},
									textfield.WithLabel("Age"),
									textfield.WithError(formState.GetError("age") != ""),
									textfield.WithSupportingText(formState.GetError("age")),
									textfield.WithSingleLine(true),
									textfield.WithModifier(size.FillMaxWidth()),
								),
								spacer.Height(24),
								button.Filled(func() {
									onAction(actions.OnSubmit())
								}, "Submit", button.WithModifier(size.FillMaxWidth())),
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
