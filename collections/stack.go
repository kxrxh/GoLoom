package collections

type Stack[T comparable] struct {
	elements []T
	size     uint64
}

// NewStack creates a new stack with the given elements.
//
// The elements parameter is a variadic input of any type.
// The function returns a new stack of the same type as the elements.
func NewStack[T comparable](elements ...T) Stack[T] {
	s := Stack[T]{elements: make([]T, len(elements))}
	copy(s.elements, elements)
	return s
}

// Copy returns a new copy of the stack.
//
// It takes no parameters.
// It returns a Stack[T].
func (s Stack[T]) Copy() Stack[T] {
	return NewStack(s.elements...)
}

// Size returns the number of elements in the stack.
//
// It does not take any parameters.
// It returns an integer representing the size of the stack.
func (s Stack[T]) Size() int {
	return len(s.elements)
}

// IsEmpty checks if the stack is empty.
//
// It returns a boolean value indicating whether the stack is empty or not.
func (s Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Push adds an element to the top of the stack.
//
// element: the element to be added to the stack
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
	s.size++
}

// Pop removes and returns the top element from the stack.
//
// Returns the element that was removed from the stack.
func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	s.size--
	return element
}

// ToSlice returns a slice containing all elements in the stack.
//
// It does not modify the original stack.
// The returned slice will have the same order as the stack, with the top element
// at the end of the slice.
//
// Returns:
//
//	[]T: A slice containing all elements in the stack.
func (s Stack[T]) ToSlice() []T {
	return append([]T{}, s.elements...)
}

// Equals checks if the stack is equal to another stack.
//
// It takes a parameter `other` of type `Stack[T]`.
// It returns a boolean value.
func (s Stack[T]) Equals(other Stack[T]) bool {
	if s.size != other.size {
		return false
	}
	for i, e := range s.elements {
		if e != other.elements[i] {
			return false
		}
	}
	return true
}

// Contains checks if the stack contains all the specified elements.
//
// It takes a variable number of elements as arguments and returns a boolean value indicating whether the stack contains all the elements.
func (s Stack[T]) Contains(elems ...T) bool {
	for _, e := range elems {
		if !s.contains(e) {
			return false
		}
	}
	return true
}

// Peek returns the top element of the stack without removing it.
//
// It does not take any parameters.
// It returns the element of type T at the top of the stack.
func (s Stack[T]) Peek() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}
	return s.elements[len(s.elements)-1]
}

// Clear removes all elements from the stack.
//
// No parameters.
// No return types.
func (s *Stack[T]) Clear() {
	s.elements = nil
	s.size = 0
}

// contains checks if the stack contains the specified element.
func (s Stack[T]) contains(element T) bool {
	for _, e := range s.elements {
		if e == element {
			return true
		}
	}
	return false
}
