package screen

import (
	"context"

	"github.com/zodimo/go-compose/lifecycle"
)

var _ lifecycle.HasViewModelScope = (*ViewModel)(nil)

func (vm *ViewModel) SetViewModelScope(ctx context.Context) {
	vm.rootContext = ctx
}

func (vm *ViewModel) Context() context.Context {
	if vm.rootContext == nil {
		panic("LoginViewModel: Context() called before SetViewModelScope()")
	}
	return vm.rootContext
}

func (vm *ViewModel) Launch(block func(ctx context.Context)) {
	go func() {
		block(vm.Context())
	}()
}
