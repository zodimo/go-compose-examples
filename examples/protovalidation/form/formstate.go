package form

import (
	"errors"

	"buf.build/go/protovalidate"
	"google.golang.org/protobuf/proto"
)

type FormState struct {
	Touched   map[string]bool
	Errors    map[string]string
	Submitted bool
}

func NewFormState() FormState {
	return FormState{
		Touched:   make(map[string]bool),
		Errors:    make(map[string]string),
		Submitted: false,
	}
}

func (f FormState) Clone() FormState {
	copyFormState := FormState{
		Touched:   make(map[string]bool),
		Errors:    make(map[string]string),
		Submitted: f.Submitted,
	}
	for k, v := range f.Touched {
		copyFormState.Touched[k] = v
	}
	for k, v := range f.Errors {
		copyFormState.Errors[k] = v
	}
	return copyFormState
}

func (f FormState) TouchField(field string) FormState {
	copyFormState := f.Clone()
	copyFormState.Touched[field] = true
	return copyFormState
}

func (f FormState) IsTouched(field string) bool {
	return f.Touched[field]
}

func (f FormState) SetError(field, msg string) FormState {
	copyFormState := f.Clone()
	copyFormState.Errors[field] = msg
	return copyFormState
}

func (f FormState) ClearError(field string) FormState {
	copyFormState := f.Clone()
	delete(copyFormState.Errors, field)
	return copyFormState
}

func (f FormState) HasErrors() bool {
	return len(f.Errors) > 0
}

func (f FormState) HasAnyTouched() bool {
	for _, v := range f.Touched {
		if v {
			return true
		}
	}
	return false
}

func (f FormState) Validate(v protovalidate.Validator, msg proto.Message) FormState {
	copyFormState := f.Clone()
	copyFormState.Submitted = true
	copyFormState.Errors = make(map[string]string) // reset errors before full validation

	err := v.Validate(msg)
	if err != nil {
		var valErr *protovalidate.ValidationError
		if errors.As(err, &valErr) {
			for _, violation := range valErr.Violations {
				fieldPath := protovalidate.FieldPathString(violation.Proto.GetField())
				copyFormState.Errors[fieldPath] = violation.Proto.GetMessage()
			}
		}
	}

	return copyFormState
}

func (f FormState) ValidateField(v protovalidate.Validator, msg proto.Message, targetField string) FormState {
	copyFormState := f.Clone()

	// Clear only the specific field error before checking
	delete(copyFormState.Errors, targetField)

	// We still validate the whole message to catch cross-field dependencies
	err := v.Validate(msg)

	if err != nil {
		var valErr *protovalidate.ValidationError
		if errors.As(err, &valErr) {
			for _, violation := range valErr.Violations {
				// Only update the map if the violation matches our active field
				fieldPath := protovalidate.FieldPathString(violation.Proto.GetField())

				if fieldPath == targetField {
					copyFormState.Errors[targetField] = violation.Proto.GetMessage()
				}
			}
		}
	}

	return copyFormState
}

func (f FormState) GetError(field string) string {
	if !f.Touched[field] && !f.Submitted {
		return ""
	}
	err, ok := f.Errors[field]
	if !ok {
		return ""
	}
	return err
}

func (f FormState) GetTouched(field string) bool {
	touched, ok := f.Touched[field]
	if !ok {
		return false
	}
	return touched
}
