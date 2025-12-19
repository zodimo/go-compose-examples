package booklist

import "github.com/zodimo/go-compose/pkg/flow"

type BookListViewModel struct {
	mutableState *flow.MutableStateFlow[BoolListState]
}

func NewBookListViewModel() *BookListViewModel {

	mutableState := flow.NewMutableStateFlow(NewBookListState())

	return &BookListViewModel{
		mutableState: mutableState,
	}

}

func (vm *BookListViewModel) OnAction(bookListAction BookListAction) {
	switch action := bookListAction.(type) {
	case *OnBookClick:

	case *OnSearchQueryChange:
		vm.mutableState.Update(func(state BoolListState) BoolListState {
			return BoolListState{
				SearchQuery:       action.Query,
				SearchBookResults: state.SearchBookResults,
				FavoriteBooks:     state.FavoriteBooks,
				IsLoading:         state.IsLoading,
				SelectedTabIndex:  state.SelectedTabIndex,
				ErrorMessage:      state.ErrorMessage,
			}
		})

	case *OnTabSelected:
		vm.mutableState.Update(func(state BoolListState) BoolListState {
			return BoolListState{
				SearchQuery:       state.SearchQuery,
				SearchBookResults: state.SearchBookResults,
				FavoriteBooks:     state.FavoriteBooks,
				IsLoading:         state.IsLoading,
				SelectedTabIndex:  action.TabIndex,
				ErrorMessage:      state.ErrorMessage,
			}
		})
	default:
		panic("unknown action")

	}

}
