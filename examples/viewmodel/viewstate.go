package main

import (
	"regexp"
	"strings"
)

type LoginViewState struct {
	emailValue        string
	emailError        string
	emailValidated    bool
	passwordValue     string
	passwordError     string
	passwordValidated bool

	// State for validation errors (only shown after interaction)
	emailTouched    bool
	passwordTouched bool

	// State for form submission
	submitted    bool
	loginSuccess bool
}

func (s *LoginViewState) Copy(options ...ViewStateOpion) *LoginViewState {
	copy := LoginViewState{
		emailValue:      s.emailValue,
		passwordValue:   s.passwordValue,
		emailTouched:    s.emailTouched,
		passwordTouched: s.passwordTouched,
		submitted:       s.submitted,
		loginSuccess:    s.loginSuccess,
	}
	for _, option := range options {
		if option != nil {
			option(&copy)
		}
	}
	return &copy
}

func (s *LoginViewState) EmailValue() string {
	return s.emailValue
}

func (s *LoginViewState) PasswordValue() string {
	return s.passwordValue
}

func (s *LoginViewState) EmailTouched() bool {
	return s.emailTouched
}

func (s *LoginViewState) PasswordTouched() bool {
	return s.passwordTouched
}

func (s *LoginViewState) Submitted() bool {
	return s.submitted
}

func (s *LoginViewState) LoginSuccess() bool {
	return s.loginSuccess
}

func (s *LoginViewState) IsFormValid() bool {
	return s.emailValidated && s.emailError == "" && s.passwordValidated && s.passwordError == ""
}

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

func (s *LoginViewState) EmailHasError() bool {
	return s.emailError != ""
}

func (s *LoginViewState) EmailError() string {
	return s.emailError
}

func (s *LoginViewState) PasswordHasError() bool {
	return s.passwordError != ""
}

func (s *LoginViewState) PasswordError() string {
	return s.passwordError
}
