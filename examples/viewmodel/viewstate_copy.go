package main

type ViewStateOpion func(s *LoginViewState)

func WithEmail(email string) ViewStateOpion {
	return func(s *LoginViewState) {
		s.emailValue = email
	}
}

func WithPassword(password string) ViewStateOpion {
	return func(s *LoginViewState) {
		s.passwordValue = password
	}
}

func WithEmailTouched(emailTouched bool) ViewStateOpion {
	return func(s *LoginViewState) {
		s.emailTouched = emailTouched
	}
}

func WithPasswordTouched(passwordTouched bool) ViewStateOpion {
	return func(s *LoginViewState) {
		s.passwordTouched = passwordTouched
	}
}

func WithEmailError(emailError string) ViewStateOpion {
	return func(s *LoginViewState) {
		s.emailError = emailError
	}
}

func WithPasswordError(passwordError string) ViewStateOpion {
	return func(s *LoginViewState) {
		s.passwordError = passwordError
	}
}

func WithEmailValidated(emailValidated bool) ViewStateOpion {
	return func(s *LoginViewState) {
		s.emailValidated = emailValidated
	}
}

func WithPasswordValidated(passwordValidated bool) ViewStateOpion {
	return func(s *LoginViewState) {
		s.passwordValidated = passwordValidated
	}
}

func WithSubmitted(submitted bool) ViewStateOpion {
	return func(s *LoginViewState) {
		s.submitted = submitted
	}
}

func WithLoginSuccess(loginSuccess bool) ViewStateOpion {
	return func(s *LoginViewState) {
		s.loginSuccess = loginSuccess
	}
}
