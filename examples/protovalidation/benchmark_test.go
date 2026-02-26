package main

import (
	"log"
	"testing"

	"buf.build/go/protovalidate"
	"gitub.com/zodimo/go-compose-examples/examples/protovalidation/form"
	protov1 "gitub.com/zodimo/go-compose-examples/gen/user/v1"
)

func BenchmarkProtovalidate_Precompiled(b *testing.B) {
	validator, err := protovalidate.New()
	if err != nil {
		log.Fatal(err)
	}

	user := &protov1.User{
		Name:  "Jane Doe",
		Email: "jane.doe@example.com",
		Age:   25,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := validator.Validate(user)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProtovalidate_NewValidatorEachTime(b *testing.B) {
	user := &protov1.User{
		Name:  "Jane Doe",
		Email: "jane.doe@example.com",
		Age:   25,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// New() compiles the rules, which is slow.
		validator, err := protovalidate.New()
		if err != nil {
			b.Fatal(err)
		}

		err = validator.Validate(user)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProtovalidate_WithErrors(b *testing.B) {
	validator, err := protovalidate.New()
	if err != nil {
		log.Fatal(err)
	}

	// Invalid user (e.g. name too short, invalid email, underage)
	user := &protov1.User{
		Name:  "J",
		Email: "invalid-email",
		Age:   10,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = validator.Validate(user)
	}
}

func BenchmarkFormState_Clone(b *testing.B) {
	fs := form.NewFormState()
	fs = fs.TouchField("name").SetError("name", "error")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fs.Clone()
	}
}

func BenchmarkFormState_TouchField(b *testing.B) {
	fs := form.NewFormState()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fs.TouchField("name")
	}
}

func BenchmarkFormState_SetError(b *testing.B) {
	fs := form.NewFormState()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fs.SetError("name", "Name is required")
	}
}

func BenchmarkFormState_ClearError(b *testing.B) {
	fs := form.NewFormState().SetError("name", "error")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fs.ClearError("name")
	}
}

func BenchmarkFormState_ValidateField(b *testing.B) {
	validator, err := protovalidate.New()
	if err != nil {
		b.Fatal(err)
	}
	fs := form.NewFormState()
	user := &protov1.User{
		Name:  "Jane",
		Email: "invalid", // Keep one invalid to match real case
		Age:   25,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fs.ValidateField(validator, user, "email")
	}
}

func BenchmarkFormState_Validate(b *testing.B) {
	validator, err := protovalidate.New()
	if err != nil {
		b.Fatal(err)
	}
	fs := form.NewFormState()
	// Invalid user to trigger errors
	user := &protov1.User{
		Name:  "J",
		Email: "invalid-email",
		Age:   10,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fs.Validate(validator, user)
	}
}
