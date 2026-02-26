package main

import (
	"github.com/zodimo/go-compose/compose"
	"github.com/zodimo/go-compose/pkg/api"
	"github.com/zodimo/go-compose/state"
	"gitub.com/zodimo/go-compose-examples/examples/protovalidation/form"
	"gitub.com/zodimo/go-compose-examples/examples/protovalidation/proto/ui"
	"gitub.com/zodimo/go-compose-examples/examples/protovalidation/screen"
	selectv1 "gitub.com/zodimo/go-compose-examples/gen/select/v1"
	protov1 "gitub.com/zodimo/go-compose-examples/gen/user/v1"
)

func UI() compose.Composable {
	return func(c api.Composer) api.Composer {

		formStateValue := state.MustRemember(c, "form-state", func() form.FormState {
			return form.NewFormState()
		})

		userModel := state.MustRemember(c, "user-model",
			func() *protov1.User {
				return &protov1.User{}
			},
		)

		submittedSuccessfullyValue := state.MustRemember(c, "submitted-successfully", func() bool {
			return false
		})

		viewModel := state.MustRemember(c, "view-model", func() *screen.ViewModel {
			return screen.NewViewModel(
				userModel,
				formStateValue,
				submittedSuccessfullyValue,
				ui.GetSelectInputFromEnum(selectv1.Gender_GENDER_UNSPECIFIED.Descriptor()),
				ui.GetSelectInputFromEnum(selectv1.UserRole_USER_ROLE_UNSPECIFIED.Descriptor()),
			)
		})

		return screen.ScreenRoot(viewModel.Get())(c)
	}
}
