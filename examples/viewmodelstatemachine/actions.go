package main

type UIAction interface {
	isUIAction()
}

type UIActionsInterface interface {
	OnSubmit() UIAction
	OnLogout() UIAction
	OnEmailChange(email string) UIAction
	OnPasswordChange(password string) UIAction
	isUIActions()
}

var _ UIActionsInterface = (*uiActions)(nil)

var UIActions UIActionsInterface = &uiActions{}

type uiActions struct{}

func (a *uiActions) isUIActions() {}

func (a *uiActions) OnSubmit() UIAction {
	return &onSubmitAction{}
}

func (a *uiActions) OnLogout() UIAction {
	return &onLogoutAction{}
}

func (a *uiActions) OnEmailChange(email string) UIAction {
	return &onEmailChangeAction{email: email}
}

func (a *uiActions) OnPasswordChange(password string) UIAction {
	return &onPasswordChangeAction{password: password}
}

type onSubmitAction struct{}

func (a *onSubmitAction) isUIAction() {}

type onLogoutAction struct{}

func (a *onLogoutAction) isUIAction() {}

type onEmailChangeAction struct {
	email string
}

func (a *onEmailChangeAction) isUIAction() {}

type onPasswordChangeAction struct {
	password string
}

func (a *onPasswordChangeAction) isUIAction() {}
