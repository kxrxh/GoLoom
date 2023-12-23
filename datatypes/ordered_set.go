package datatypes

import "sort"

type OrderedSet[T comparable] struct {
	lookup   map[T]bool
	elements []T
}

// NewOrderedSet creates a new ordered set with the given elements.
//
// The elements parameter is a variadic parameter that accepts any number
// of elements of type T, where T is a comparable type.
//
// The function returns an OrderedSet[T] containing the elements in the
// same order as they were provided.
func NewOrderedSet[T comparable](elements ...T) OrderedSet[T] {
	s := make(map[T]bool, len(elements))
	elOrder := make([]T, 0, len(elements)/2)
	for _, e := range elements {
		if s[e] {
			continue
		}
		elOrder = append(elOrder, e)
		s[e] = true
	}
	return OrderedSet[T]{lookup: s, elements: elOrder}
}

// Add adds the given elements to the ordered set.
//
// It takes a variadic parameter `elems` which represents the elements to be added.
// There is no return value.
func (s *OrderedSet[T]) Add(elems ...T) {
	for _, e := range elems {
		if s.lookup[e] {
			continue
		}
		s.lookup[e] = true
		s.elements = append(s.elements, e)
	}
}

// Remove removes the specified elements from the OrderedSet.
//
// The Remove function takes a variadic parameter 'elems' of type T, representing the elements to be removed from the OrderedSet.
// It iterates over each element in 'elems' using a range loop and deletes it from the lookup map of the OrderedSet.
func (s *OrderedSet[T]) Remove(elems ...T) {
	for _, e := range elems {
		delete(s.lookup, e)
		s.elements = RemoveElement(s.elements, e)
	}
}

// Contains checks if the elements are present in the OrderedSet.
//
// Parameters:
// - elems: The elements to be checked.
//
// Returns:
// - bool: True if all elements are present, false otherwise.
func (s *OrderedSet[T]) Contains(elems ...T) bool {
	for _, e := range elems {
		if !s.lookup[e] {
			return false
		}
	}
	return true
}

// ToSlice returns a slice containing all elements of the ordered set.
//
// No parameters.
// Returns a slice of type T.
func (s OrderedSet[T]) ToSlice() []T {
	return append([]T{}, s.elements...)
}

// Get returns the element at the specified index in the OrderedSet.
//
// Parameters:
// - i: the index of the element to retrieve.
//
// Returns:
// - T: the element at the specified index.
func (s OrderedSet[T]) Get(i int) T {
	if i < 0 || i >= len(s.elements) {
		var zero T
		return zero
	}
	return s.elements[i]
}

// Len returns the length of the OrderedSet.
//
// It returns an integer representing the number of elements in the OrderedSet.
func (s OrderedSet[T]) Len() int {
	return len(s.elements)
}

// IsEmpty returns true if the OrderedSet is empty, otherwise it returns false.
//
// No parameters.
// Return type: bool.
func (s OrderedSet[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Equals checks if the OrderedSet is equal to another OrderedSet.
//
// It takes in the other OrderedSet as a parameter and returns a boolean value indicating if the two sets are equal.
func (s OrderedSet[T]) Equals(other OrderedSet[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	for i, e := range s.elements {
		if e != other.elements[i] {
			return false
		}
	}
	return true
}

// SortWithComparator sorts the elements of the OrderedSet using the provided comparator function.
// The comparator function should return true if the element at index i is less than the element at index j.
func (s *OrderedSet[T]) SortWithComparator(comparator func(i, j int) bool) {
	sort.Slice(s.elements, comparator)
}

// Union returns a new OrderedSet that is the union of the current OrderedSet and the other OrderedSet.
//
// Parameters:
//   - other: the OrderedSet to be combined with the current OrderedSet.
//
// Return type:
//   - OrderedSet[T]: a new OrderedSet that contains all the unique elements from both the current OrderedSet and the other OrderedSet.
func (s OrderedSet[T]) Union(other OrderedSet[T]) OrderedSet[T] {
	unionSet := NewOrderedSet(s.elements...)
	unionSet.Add(other.elements...)
	return unionSet
}

// Intersection returns a new OrderedSet that contains the intersection of the calling OrderedSet and the other OrderedSet.
//
// Parameters:
// - other: the other OrderedSet to compute the intersection with.
//
// Returns:
// - OrderedSet[T]: a new OrderedSet that contains the intersection of the calling OrderedSet and the other OrderedSet.
func (s OrderedSet[T]) Intersection(other OrderedSet[T]) OrderedSet[T] {
	intersectionSet := NewOrderedSet[T]()
	for _, e := range s.elements {
		if other.Contains(e) {
			intersectionSet.Add(e)
		}
	}
	return intersectionSet
}

// Difference returns a new OrderedSet containing the elements that are present in the receiver
// (s) but not in the other OrderedSet (other).
//
// The function takes one parameter:
// - other: an OrderedSet of type T.
//
// The function returns an OrderedSet of type T.
func (s OrderedSet[T]) Difference(other OrderedSet[T]) OrderedSet[T] {
	differenceSet := NewOrderedSet(s.elements...)
	differenceSet.Remove(other.elements...)
	return differenceSet
}

// Clear clears the elements and lookup map of the OrderedSet.
//
// No parameters.
// No return values.
func (s *OrderedSet[T]) Clear() {
	s.elements = make([]T, 0)
	s.lookup = make(map[T]bool)
}
