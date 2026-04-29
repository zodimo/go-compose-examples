package main

import (
	"fmt"
	"strings"

	"github.com/zodimo/go-compose/compose/viewmodel"
	"github.com/zodimo/go-compose/pkg/flow"
)

type LoginViewModel struct {
	viewmodel.ViewModel

	viewStateMutableStateFlow *flow.MutableStateFlow[*LoginViewState]
}

func NewViewModel() *LoginViewModel {
	return &LoginViewModel{
		viewStateMutableStateFlow: flow.NewMutableStateFlow[*LoginViewState](
			&LoginViewState{},
		),
	}
}

func (vm *LoginViewModel) AsStateFlow() flow.StateFlow[*LoginViewState] {
	return vm.viewStateMutableStateFlow.AsStateFlow()
}

func (vm *LoginViewModel) OnAction(action UIAction) {
	switch action := action.(type) {
	case *onSubmitAction:

		vm.viewStateMutableStateFlow.Update(func(current *LoginViewState) *LoginViewState {

			// Re-validate at submission time
			currentEmail := strings.TrimSpace(current.EmailValue())
			currentPassword := current.PasswordValue()

			emailVal := validateEmail(currentEmail)
			passwordVal := validatePassword(currentPassword)

			if emailVal.Valid && passwordVal.Valid {
				return current.Copy(
					WithEmailTouched(true),
					WithEmailError(emailVal.Message),
					WithEmailValidated(true),
					WithPasswordTouched(true),
					WithPasswordError(passwordVal.Message),
					WithPasswordValidated(true),
					WithSubmitted(true),
					WithLoginSuccess(true),
				)
			} else {
				return current.Copy(
					WithEmailTouched(true),
					WithEmailError(emailVal.Message),
					WithEmailValidated(true),
					WithPasswordTouched(true),
					WithPasswordError(passwordVal.Message),
					WithPasswordValidated(true),
					WithSubmitted(true),
				)
			}

		})
	case *onLogoutAction:
		vm.viewStateMutableStateFlow.Update(func(current *LoginViewState) *LoginViewState {
			return &LoginViewState{}
		})
	case *onEmailChangeAction:
		vm.viewStateMutableStateFlow.Update(func(current *LoginViewState) *LoginViewState {

			// Validate fields
			emailValidation := validateEmail(current.EmailValue())

			return current.Copy(
				WithEmail(action.email),
				WithEmailTouched(true),
				WithEmailError(emailValidation.Message),
				WithEmailValidated(true),
			)
		})
	case *onPasswordChangeAction:
		vm.viewStateMutableStateFlow.Update(func(current *LoginViewState) *LoginViewState {
			passwordValidation := validatePassword(current.PasswordValue())

			return current.Copy(
				WithPassword(action.password),
				WithPasswordTouched(true),
				WithPasswordError(passwordValidation.Message),
				WithPasswordValidated(true),
			)
		})
	default:
		panic(fmt.Sprintf("Unknown action %T", action))

	}
}
