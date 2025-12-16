package main

// Filter represents the current filter state for displaying todos.
type Filter int

const (
	FilterAll Filter = iota
	FilterActive
	FilterCompleted
)

// Todo represents a single todo item.
type Todo struct {
	ID        int
	Text      string
	Completed bool
}

// TodoState holds the entire application state.
type TodoState struct {
	Todos     []Todo
	Filter    Filter
	EditingID int // -1 when not editing
	NextID    int
}

// NewTodoState creates a new empty TodoState.
func NewTodoState() *TodoState {
	return &TodoState{
		Todos:     []Todo{},
		Filter:    FilterAll,
		EditingID: -1,
		NextID:    1,
	}
}

// AddTodo adds a new todo with the given text.
func (s *TodoState) AddTodo(text string) *TodoState {
	if text == "" {
		return s
	}
	todos := append(s.Todos, Todo{
		ID:        s.NextID,
		Text:      text,
		Completed: false,
	})

	return &TodoState{
		Todos:     todos,
		Filter:    s.Filter,
		EditingID: s.EditingID,
		NextID:    s.NextID + 1,
	}
}

// ToggleTodo toggles the completion status of a todo.
func (s *TodoState) ToggleTodo(id int) *TodoState {
	todos := make([]Todo, len(s.Todos))
	copy(todos, s.Todos)
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = !todos[i].Completed
			return &TodoState{
				Todos:     todos,
				Filter:    s.Filter,
				EditingID: s.EditingID,
				NextID:    s.NextID,
			}
		}
	}
	return s
}

// DeleteTodo removes a todo by ID.
func (s *TodoState) DeleteTodo(id int) *TodoState {
	newTodos := make([]Todo, 0, len(s.Todos))
	for _, t := range s.Todos {
		if t.ID != id {
			newTodos = append(newTodos, t)
		}
	}
	return &TodoState{
		Todos:     newTodos,
		Filter:    s.Filter,
		EditingID: s.EditingID,
		NextID:    s.NextID,
	}
}

// UpdateTodo updates the text of a todo.
func (s *TodoState) UpdateTodo(id int, text string) *TodoState {
	todos := make([]Todo, len(s.Todos))
	copy(todos, s.Todos)
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Text = text
			return &TodoState{
				Todos:     todos,
				Filter:    s.Filter,
				EditingID: s.EditingID,
				NextID:    s.NextID,
			}
		}
	}
	return s
}

// ClearCompleted removes all completed todos.
func (s *TodoState) ClearCompleted() *TodoState {
	newTodos := make([]Todo, 0, len(s.Todos))
	for _, t := range s.Todos {
		if !t.Completed {
			newTodos = append(newTodos, t)
		}
	}
	return &TodoState{
		Todos:     newTodos,
		Filter:    s.Filter,
		EditingID: s.EditingID,
		NextID:    s.NextID,
	}
}

// ToggleAll sets all todos to the given completion state.
func (s *TodoState) ToggleAll(completed bool) *TodoState {
	todos := make([]Todo, len(s.Todos))
	copy(todos, s.Todos)
	for i := range todos {
		todos[i].Completed = completed
	}
	return &TodoState{
		Todos:     todos,
		Filter:    s.Filter,
		EditingID: s.EditingID,
		NextID:    s.NextID,
	}
}

// AllCompleted returns true if all todos are completed.
func (s *TodoState) AllCompleted() bool {
	if len(s.Todos) == 0 {
		return false
	}
	for _, t := range s.Todos {
		if !t.Completed {
			return false
		}
	}
	return true
}

// FilteredTodos returns todos matching the current filter.
func (s *TodoState) FilteredTodos() []Todo {
	switch s.Filter {
	case FilterActive:
		result := make([]Todo, 0)
		for _, t := range s.Todos {
			if !t.Completed {
				result = append(result, t)
			}
		}
		return result
	case FilterCompleted:
		result := make([]Todo, 0)
		for _, t := range s.Todos {
			if t.Completed {
				result = append(result, t)
			}
		}
		return result
	default:
		return s.Todos
	}
}

// ActiveCount returns the number of uncompleted todos.
func (s *TodoState) ActiveCount() int {
	count := 0
	for _, t := range s.Todos {
		if !t.Completed {
			count++
		}
	}
	return count
}

// CompletedCount returns the number of completed todos.
func (s *TodoState) CompletedCount() int {
	count := 0
	for _, t := range s.Todos {
		if t.Completed {
			count++
		}
	}
	return count
}

// SetFilter sets the current filter.
func (s *TodoState) SetFilter(f Filter) *TodoState {
	return &TodoState{
		Todos:     s.Todos,
		Filter:    f,
		EditingID: s.EditingID,
		NextID:    s.NextID,
	}
}

// SetEditing starts editing mode for the given todo ID.
func (s *TodoState) SetEditing(id int) *TodoState {
	return &TodoState{
		Todos:     s.Todos,
		Filter:    s.Filter,
		EditingID: id,
		NextID:    s.NextID,
	}
}

// CancelEditing exits editing mode without saving.
func (s *TodoState) CancelEditing() *TodoState {
	return &TodoState{
		Todos:     s.Todos,
		Filter:    s.Filter,
		EditingID: -1,
		NextID:    s.NextID,
	}
}

// IsEditing returns true if the given todo ID is being edited.
func (s *TodoState) IsEditing(id int) bool {
	return s.EditingID == id
}
