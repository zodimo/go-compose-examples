package screen

import (
	"gitub.com/zodimo/go-compose-examples/examples/protovalidation/form"
	protov1 "gitub.com/zodimo/go-compose-examples/examples/protovalidation/proto/v1"
)

type LoginCopyOption func(*ViewState)

func (state *ViewState) Copy(options ...LoginCopyOption) *ViewState {
	return CopyLoginState(state, options...)
}

func CopyLoginState(state *ViewState, options ...LoginCopyOption) *ViewState {
	stateCopy := ViewState{
		user:                  state.User(),
		formState:             state.FormState(),
		submittedSuccessfully: state.SubmittedSuccessfully(),
	}
	for _, option := range options {
		if option != nil {
			option(&stateCopy)
		}
	}
	return &stateCopy
}

func WithUser(user *protov1.User) LoginCopyOption {
	return func(state *ViewState) {
		state.user = user
	}
}

func WithFormState(formState form.FormState) LoginCopyOption {
	return func(state *ViewState) {
		state.formState = formState
	}
}

func WithSubmittedSuccessfully(submittedSuccessfully bool) LoginCopyOption {
	return func(state *ViewState) {
		state.submittedSuccessfully = submittedSuccessfully
	}
}
