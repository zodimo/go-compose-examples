package screen

// EventStyle, Data orientated
type Action interface {
	isAction()
}

var _ Actions = (*actions)(nil)

// Actions is the Action in the Flux pattern to cause a possible state change
type Actions interface {
	isActions()

	FieldOnChange(fieldPath string, value string) Action
	OnSubmit() Action
}

type actions struct {
}

func NewActions() Actions {
	return actions{}
}

func (a actions) isActions() {}

func (a actions) FieldOnChange(fieldPath string, value string) Action {
	return &fieldOnChangeAction{
		fieldPath: fieldPath,
		value:     value,
	}
}

type fieldOnChangeAction struct {
	fieldPath string
	value     string
}

func (a fieldOnChangeAction) isAction() {}

func (a actions) OnSubmit() Action {
	return &onSubmitAction{}
}

type onSubmitAction struct {
}

func (a onSubmitAction) isAction() {}
