package screen

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"buf.build/go/protovalidate"
	"github.com/zodimo/go-compose/pkg/flow"
	"github.com/zodimo/go-compose/state"
	"gitub.com/zodimo/go-compose-examples/examples/protovalidation/form"
	selectv1 "gitub.com/zodimo/go-compose-examples/gen/select/v1"
	uiv1 "gitub.com/zodimo/go-compose-examples/gen/ui/v1"
	protov1 "gitub.com/zodimo/go-compose-examples/gen/user/v1"
	"google.golang.org/protobuf/proto"
)

type ViewModel struct {
	mutableState *flow.MutableStateFlow[*ViewState]

	rootContext context.Context
	userProto   state.MutableValueTyped[*protov1.User]
	formState   state.MutableValueTyped[form.FormState]

	validator protovalidate.Validator

	submittedSuccessfully state.MutableValueTyped[bool]

	genderSelect *uiv1.SelectInput
	roleSelect   *uiv1.SelectInput
}

func (vm *ViewModel) AsStateFlow() flow.StateFlow[*ViewState] {
	return vm.mutableState.AsStateFlow()
}

func NewViewModel(
	userProtoValue state.MutableValueTyped[*protov1.User],
	formStateValue state.MutableValueTyped[form.FormState],
	submittedSuccessfullyValue state.MutableValueTyped[bool],
	genderSelect *uiv1.SelectInput,
	roleSelect *uiv1.SelectInput,
) *ViewModel {

	validator, err := protovalidate.New()
	if err != nil {
		log.Fatal(err)
	}

	vm := &ViewModel{
		mutableState: flow.NewMutableStateFlow(
			NewViewState(
				userProtoValue.Get(),
				formStateValue.Get(),
				genderSelect,
				roleSelect,
				submittedSuccessfullyValue.Get(),
			),
		),
		userProto:             userProtoValue,
		formState:             formStateValue,
		validator:             validator,
		submittedSuccessfully: submittedSuccessfullyValue,
		genderSelect:          genderSelect,
		roleSelect:            roleSelect,
	}
	return vm
}

func (vm *ViewModel) OnAction(action Action) {
	// fmt.Printf("LoginViewModel: action: %T\n", action)
	switch action := action.(type) {
	case *fieldOnChangeAction:
		switch action.fieldPath {
		case "name":
			vm.submittedSuccessfully.Update(func(b bool) bool {
				return false
			})
			vm.userProto.Update(func(p *protov1.User) *protov1.User {
				userClone := proto.CloneOf(p)
				userClone.Name = action.value
				return userClone
			})
			vm.formState.Update(func(fs form.FormState) form.FormState {
				return fs.TouchField("name").ValidateField(vm.validator, vm.userProto.Get(), "name")
			})
			vm.mutableState.Update(func(state *ViewState) *ViewState {
				return state.Copy(
					WithUser(vm.userProto.Get()),
					WithFormState(vm.formState.Get()),
				)
			})

		case "email":
			vm.submittedSuccessfully.Update(func(b bool) bool {
				return false
			})
			vm.userProto.Update(func(p *protov1.User) *protov1.User {
				//Critical: need to clone the proto before modifying it
				userClone := proto.CloneOf(p)
				userClone.Email = action.value
				return userClone
			})
			vm.formState.Update(func(fs form.FormState) form.FormState {
				return fs.TouchField("email").ValidateField(vm.validator, vm.userProto.Get(), "email")
			})
			vm.mutableState.Update(func(state *ViewState) *ViewState {
				return state.Copy(
					WithUser(vm.userProto.Get()),
					WithFormState(vm.formState.Get()),
					WithSubmittedSuccessfully(vm.submittedSuccessfully.Get()),
				)
			})
		case "age":
			vm.submittedSuccessfully.Update(func(b bool) bool {
				return false
			})
			if action.value == "" {
				vm.userProto.Update(func(p *protov1.User) *protov1.User {
					//Critical: need to clone the proto before modifying it
					userClone := proto.CloneOf(p)
					//unspecified
					userClone.Age = 0
					return userClone
				})
				vm.formState.Update(func(fs form.FormState) form.FormState {
					return fs.TouchField("age").ValidateField(vm.validator, vm.userProto.Get(), "age")
				})
				vm.mutableState.Update(func(state *ViewState) *ViewState {
					return state.Copy(
						WithUser(vm.userProto.Get()),
						WithFormState(vm.formState.Get()),
						WithSubmittedSuccessfully(vm.submittedSuccessfully.Get()),
					)
				})
				return
			}
			ageInt, err := strconv.ParseInt(action.value, 10, 32)
			if err != nil {
				vm.formState.Update(func(fs form.FormState) form.FormState {
					fState := fs.TouchField("age")
					return fState.SetError("age", "Must be a valid integer")
				})
			} else {
				vm.userProto.Update(func(p *protov1.User) *protov1.User {
					//Critical: need to clone the proto before modifying it
					userClone := proto.CloneOf(p)
					userClone.Age = int32(ageInt)
					return userClone
				})
				vm.formState.Update(func(fs form.FormState) form.FormState {
					return fs.TouchField("age").ValidateField(vm.validator, vm.userProto.Get(), "age")
				})
				vm.mutableState.Update(func(state *ViewState) *ViewState {
					return state.Copy(
						WithUser(vm.userProto.Get()),
						WithFormState(vm.formState.Get()),
						WithSubmittedSuccessfully(vm.submittedSuccessfully.Get()),
					)
				})
			}

		case "gender":
			vm.submittedSuccessfully.Update(func(b bool) bool {
				return false
			})
			gender := selectv1.Gender_GENDER_UNSPECIFIED
			switch action.value {
			case selectv1.Gender_GENDER_FEMALE.String():
				gender = selectv1.Gender_GENDER_FEMALE
			case selectv1.Gender_GENDER_MALE.String():
				gender = selectv1.Gender_GENDER_MALE
			case selectv1.Gender_GENDER_OTHER.String():
				gender = selectv1.Gender_GENDER_OTHER
			}

			vm.userProto.Update(func(p *protov1.User) *protov1.User {
				//Critical: need to clone the proto before modifying it
				userClone := proto.CloneOf(p)
				userClone.Gender = gender
				return userClone
			})

			// fmt.Println(protojson.Format(vm.userProto.Get()))
			vm.formState.Update(func(fs form.FormState) form.FormState {
				return fs.TouchField("gender").ValidateField(vm.validator, vm.userProto.Get(), "gender")
			})
			vm.mutableState.Update(func(state *ViewState) *ViewState {
				return state.Copy(
					WithUser(vm.userProto.Get()),
					WithFormState(vm.formState.Get()),
					WithSubmittedSuccessfully(vm.submittedSuccessfully.Get()),
				)
			})

		case "role":
			vm.submittedSuccessfully.Update(func(b bool) bool {
				return false
			})
			role := selectv1.UserRole_USER_ROLE_UNSPECIFIED
			switch action.value {
			case selectv1.UserRole_USER_ROLE_NONE.String():
				role = selectv1.UserRole_USER_ROLE_NONE
			case selectv1.UserRole_USER_ROLE_MEMBER.String():
				role = selectv1.UserRole_USER_ROLE_MEMBER
			case selectv1.UserRole_USER_ROLE_ADMIN.String():
				role = selectv1.UserRole_USER_ROLE_ADMIN
			}

			vm.userProto.Update(func(p *protov1.User) *protov1.User {
				//Critical: need to clone the proto before modifying it
				userClone := proto.CloneOf(p)
				userClone.Role = role
				return userClone
			})

			// fmt.Println(protojson.Format(vm.userProto.Get()))
			vm.formState.Update(func(fs form.FormState) form.FormState {
				return fs.TouchField("role").ValidateField(vm.validator, vm.userProto.Get(), "role")
			})
			vm.mutableState.Update(func(state *ViewState) *ViewState {
				return state.Copy(
					WithUser(vm.userProto.Get()),
					WithFormState(vm.formState.Get()),
					WithSubmittedSuccessfully(vm.submittedSuccessfully.Get()),
				)
			})

		default:
			panic(fmt.Sprintf("LoginViewModel: unknown field path %s", action.fieldPath))
		}
	case *onSubmitAction:
		vm.formState.Update(func(fs form.FormState) form.FormState {
			return fs.Validate(vm.validator, vm.userProto.Get())
		})
		if !vm.formState.Get().HasErrors() {
			vm.submittedSuccessfully.Set(true)
		}
		vm.mutableState.Update(func(state *ViewState) *ViewState {
			return state.Copy(
				WithUser(vm.userProto.Get()),
				WithFormState(vm.formState.Get()),
				WithSubmittedSuccessfully(vm.submittedSuccessfully.Get()),
			)
		})

	default:
		panic(fmt.Sprintf("LoginViewModel: unknown action %T", action))
	}

}
