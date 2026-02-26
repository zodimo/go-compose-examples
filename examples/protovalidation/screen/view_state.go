package screen

import (
	"gitub.com/zodimo/go-compose-examples/examples/protovalidation/form"
	selectv1 "gitub.com/zodimo/go-compose-examples/gen/select/v1"
	uiv1 "gitub.com/zodimo/go-compose-examples/gen/ui/v1"
	protov1 "gitub.com/zodimo/go-compose-examples/gen/user/v1"
)

type ViewState struct {
	user                  *protov1.User
	formState             form.FormState
	genderSelect          *uiv1.SelectInput
	roleSelect            *uiv1.SelectInput
	submittedSuccessfully bool
}

func NewViewState(
	user *protov1.User,
	formState form.FormState,

	genderSelect *uiv1.SelectInput,
	roleSelect *uiv1.SelectInput,
	submittedSuccessfully bool,

) *ViewState {
	return &ViewState{
		user:                  user,
		formState:             formState,
		genderSelect:          genderSelect,
		roleSelect:            roleSelect,
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

func (vs *ViewState) GenderSelect() *uiv1.SelectInput {
	return vs.genderSelect
}

func (vs *ViewState) RoleSelect() *uiv1.SelectInput {
	return vs.roleSelect
}

func (vs *ViewState) GenderLabelForSelectedOption(option selectv1.Gender) string {
	for _, v := range vs.GenderSelect().Options {
		if v.Value == option.String() {
			if v.Value == selectv1.Gender_GENDER_UNSPECIFIED.String() {
				if touched, ok := vs.formState.Touched["gender"]; ok && touched {
					return v.Label
				}
			} else {
				return v.Label
			}
		}
	}
	return ""
}

func (vs *ViewState) RoleLabelForSelectedOption(option selectv1.UserRole) string {
	for _, v := range vs.RoleSelect().Options {
		if v.Value == option.String() {
			if v.Value == selectv1.UserRole_USER_ROLE_UNSPECIFIED.String() {
				if touched, ok := vs.formState.Touched["role"]; ok && touched {
					return v.Label
				}
			} else {
				return v.Label
			}
		}
	}
	return ""
}
