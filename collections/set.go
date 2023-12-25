package collections

type Set[T comparable] struct {
	elements map[T]bool
}

// NewSet creates a new set with the given elements.
//
// Accepts a variable number of parameters of any comparable type.
// Returns a new set containing the unique elements.
func NewSet[T comparable](elems ...T) Set[T] {
	s := make(map[T]bool, len(elems))
	for _, e := range elems {
		s[e] = true
	}
	return Set[T]{elements: s}
}

// ToOrderedSet converts the current Set to an OrderedSet.
//
// It returns an OrderedSet[T] containing the elements of the Set in the
// order they were added.
func (s *Set[T]) ToOrderedSet() OrderedSet[T] {
	return OrderedSet[T]{lookup: s.elements, elements: s.ToSlice()}
}

// NewSetOfSize creates a new Set of a specified size.
//
// It takes a size parameter, which specifies the size of the Set to be created.
// The size parameter is of type uint64.
//
// The function returns a Set of type T, where T is a comparable type.
func NewSetOfSize[T comparable](size uint64) Set[T] {
	return Set[T]{elements: make(map[T]bool, size)}
}

// Add adds elements to the set.
//
// It accepts a variadic number of elements of
// type T. It does not return any value.
func (s *Set[T]) Add(elems ...T) {
	for _, e := range elems {
		s.elements[e] = true
	}
}

// Remove deletes the specified elements from the set.
//
// elems are one or more elements to be removed.
func (s *Set[T]) Remove(elems ...T) {
	for _, e := range elems {
		delete(s.elements, e)
	}
}

// Contains checks if all elements are present in the set.
//
// elems is a variadic parameter of type T indicating the elements
// to check for presence in the set.
// Returns true if all elements are present, otherwise false.
func (s Set[T]) Contains(elems ...T) bool {
	for _, e := range elems {
		if !s.elements[e] {
			return false
		}
	}
	return true
}

// ToSlice converts the Set to a slice of type []T.
//
// No parameters.
// Returns a slice containing all elements of the set.
func (s Set[T]) ToSlice() []T {
	var elems []T
	for e := range s.elements {
		elems = append(elems, e)
	}
	return elems
}

// Len returns the number of elements in the Set.
//
// This method has no parameters.
// Return type is an int representing the number of elements.
func (s Set[T]) Len() int {
	return len(s.elements)
}

// IsEmpty checks if the set is empty.
//
// No parameters.
// Returns true if the set is empty, otherwise false.
func (s Set[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Copy creates a new Set as a copy of the current Set.
//
// No parameters.
// Returns a new Set containing copies of the elements.
func (s Set[T]) Copy() Set[T] {
	cp := make(map[T]bool, s.Len())
	for e := range s.elements {
		cp[e] = true
	}
	return Set[T]{elements: cp}
}

// Equals checks if two sets contain the same elements.
//
// 'other' is the set to compare with the receiver set.
// Returns true if sets are equal, otherwise false.
func (s Set[T]) Equals(other Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	for e := range s.elements {
		if !other.elements[e] {
			return false
		}
	}
	return true
}

// Union returns a new set which is the union of the current set and another.
//
// other is the other set to perform the union with.
// Returns a new set containing all elements from both sets.
func (s Set[T]) Union(other Set[T]) Set[T] {
	cp := s.Copy()
	for e := range other.elements {
		cp.elements[e] = true
	}
	return cp
}

// Intersection computes the set intersection of the current set with another.
//
// other: Set[T] to intersect with.
// Returns a new Set[T] containing the intersection.
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	intersection := NewSet[T]()
	for e := range s.elements {
		if other.elements[e] {
			intersection.elements[e] = true
		}
	}
	return intersection
}

// Difference returns a new set containing elements that
// are in the receiver set but not in the 'other' set.
//
// 'other' is the set to compare against.
// Returns a new Set[T] with the unique elements.
func (s Set[T]) Difference(other Set[T]) Set[T] {
	difference := NewSet[T]()
	for e := range s.elements {
		if !other.elements[e] {
			difference.elements[e] = true
		}
	}
	return difference
}

// SymmetricDifference returns a new set containing elements
// which are in either of the sets but not in their intersection.
//
// 'other' is the set to be compared with the receiver set.
// Returns a new set of the same type.
func (s Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	union := s.Union(other)
	intersection := s.Intersection(other)
	return union.Difference(intersection)
}

// IsSubset checks if the set is a subset of another.
//
// It iterates through the set and returns false if any
// element is not present in the other set.
// Returns true if all elements are present.
func (s Set[T]) IsSubset(other Set[T]) bool {
	for e := range s.elements {
		if !other.elements[e] {
			return false
		}
	}
	return true
}
