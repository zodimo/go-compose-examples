package screen

import (
	"gitub.com/zodimo/go-compose-examples/examples/protovalidation/form"
	protov1 "gitub.com/zodimo/go-compose-examples/examples/protovalidation/proto/v1"
)

type ViewState struct {
	user                  *protov1.User
	formState             form.FormState
	submittedSuccessfully bool
}

func NewViewState(
	user *protov1.User,
	formState form.FormState,

	submittedSuccessfully bool,
) *ViewState {
	return &ViewState{
		user:                  user,
		formState:             formState,
		submittedSuccessfully: submittedSuccessfully,
	}
}

func (vs *ViewState) User() *protov1.User {
	return vs.user
}

func (vs *ViewState) FormState() form.FormState {
	return vs.formState
}

func (vs *ViewState) SubmittedSuccessfully() bool {
	return vs.submittedSuccessfully
}
