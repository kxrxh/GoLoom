package collections

import (
	"testing"
)

func TestSet_Add(t *testing.T) {
	s := NewSet[int]()
	s.Add(1, 2, 3)

	expected := NewSet(1, 2, 3)
	if !s.Equals(expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}

	// Test adding duplicates
	s.Add(1, 2)
	expected = NewSet(1, 2, 3)
	if !s.Equals(expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestSet_Remove(t *testing.T) {
	s := NewSet(1, 2, 3)
	s.Remove(2)

	expected := NewSet(1, 3)
	if !s.Equals(expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}

	// Test removing non-existent element
	s.Remove(5)
	if !s.Equals(expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestSet_Contains(t *testing.T) {
	s := NewSet[int](1, 2, 3)

	if !s.Contains(1, 2) {
		t.Errorf("Expected true, but got false")
	}

	if s.Contains(4) {
		t.Errorf("Expected false, but got true")
	}
}

func TestSet_ToSlice(t *testing.T) {
	s := NewSet[int](1, 2, 3)

	expected := []int{1, 2, 3}
	result := s.ToSlice()

	if len(result) != len(expected) {
		t.Errorf("Expected length %v, but got %v", len(expected), len(result))
	}

	// Test ToSlice on an empty set
	emptySet := NewSet[int]()
	emptyResult := emptySet.ToSlice()

	if len(emptyResult) != 0 {
		t.Errorf("Expected empty slice, but got %v", emptyResult)
	}
}

func TestSet_Len(t *testing.T) {
	s := NewSet[int](1, 2, 3, 3, 2)

	result := s.Len()
	expected := 3

	if result != expected {
		t.Errorf("Expected length %v, but got %v", expected, result)
	}

	// Test Len on an empty set
	emptySet := NewSet[int]()
	emptyResult := emptySet.Len()

	if emptyResult != 0 {
		t.Errorf("Expected length 0, but got %v", emptyResult)
	}
}

func TestSet_IsEmpty(t *testing.T) {
	s := NewSet[int](1, 2, 3)

	if s.IsEmpty() {
		t.Errorf("Expected non-empty set, but got empty")
	}

	// Test IsEmpty on an empty set
	emptySet := NewSet[int]()

	if !emptySet.IsEmpty() {
		t.Errorf("Expected empty set, but got non-empty")
	}
}

func TestSet_Copy(t *testing.T) {
	s := NewSet[int](1, 2, 3)
	copySet := s.Copy()

	if !s.Equals(copySet) {
		t.Errorf("Expected sets to be equal, but they are not")
	}

	// Modify the original set and ensure the copy remains unchanged
	s.Add(4)

	if s.Equals(copySet) {
		t.Errorf("Expected sets to be different after modification, but they are still equal")
	}
}

func TestSet_Equals(t *testing.T) {
	s1 := NewSet[int](1, 2, 3)
	s2 := NewSet[int](1, 2, 3)
	s3 := NewSet[int](1, 2, 4)

	if !s1.Equals(s2) {
		t.Errorf("Expected sets to be equal, but they are not")
	}

	if s1.Equals(s3) {
		t.Errorf("Expected sets to be different, but they are equal")
	}
}

func TestSet_Union(t *testing.T) {
	s1 := NewSet[int](1, 2, 3)
	s2 := NewSet[int](3, 4, 5)

	unionSet := s1.Union(s2)

	expected := NewSet[int](1, 2, 3, 4, 5)

	if !unionSet.Equals(expected) {
		t.Errorf("Expected union set %v, but got %v", expected, unionSet)
	}
}

func TestSet_Intersection(t *testing.T) {
	s1 := NewSet[int](1, 2, 3, 4)
	s2 := NewSet[int](3, 4, 5)

	intersectionSet := s1.Intersection(s2)

	expected := NewSet[int](3, 4)

	if !intersectionSet.Equals(expected) {
		t.Errorf("Expected intersection set %v, but got %v", expected, intersectionSet)
	}
}

func TestSet_Difference(t *testing.T) {
	s1 := NewSet[int](1, 2, 3, 4)
	s2 := NewSet[int](3, 4, 5)

	differenceSet := s1.Difference(s2)

	expected := NewSet[int](1, 2)

	if !differenceSet.Equals(expected) {
		t.Errorf("Expected difference set %v, but got %v", expected, differenceSet)
	}
}

func TestSet_SymmetricDifference(t *testing.T) {
	s1 := NewSet[int](1, 2, 3, 4)
	s2 := NewSet[int](3, 4, 5)

	symmetricDifferenceSet := s1.SymmetricDifference(s2)

	expected := NewSet[int](1, 2, 5)

	if !symmetricDifferenceSet.Equals(expected) {
		t.Errorf("Expected symmetric difference set %v, but got %v", expected, symmetricDifferenceSet)
	}
}

func TestSet_IsSubset(t *testing.T) {
	s1 := NewSet[int](1, 2)
	s2 := NewSet[int](1, 2, 3, 4)

	if !s1.IsSubset(s2) {
		t.Errorf("Expected s1 to be a subset of s2, but it is not")
	}

	// Test with an empty set (should always be a subset)
	emptySet := NewSet[int]()
	if !emptySet.IsSubset(s1) {
		t.Errorf("Expected empty set to be a subset, but it is not")
	}
}
