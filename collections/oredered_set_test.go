package collections

import (
	"reflect"
	"testing"
)

func TestNewOrderedSet(t *testing.T) {
	// Test case 1
	set1 := NewOrderedSet(1, 2, 3)
	expected1 := OrderedSet[int]{lookup: map[int]bool{1: true, 2: true, 3: true}, elements: []int{1, 2, 3}}
	if !reflect.DeepEqual(set1, expected1) {
		t.Errorf("NewOrderedSet: expected %v, got %v", expected1, set1)
	}

	// Test case 2
	set2 := NewOrderedSet("a", "b", "c")
	expected2 := OrderedSet[string]{lookup: map[string]bool{"a": true, "b": true, "c": true}, elements: []string{"a", "b", "c"}}
	if !reflect.DeepEqual(set2, expected2) {
		t.Errorf("NewOrderedSet: expected %v, got %v", expected2, set2)
	}
}

func TestAdd(t *testing.T) {
	set := NewOrderedSet(1, 2, 3)
	set.Add(4, 5)

	expected := OrderedSet[int]{lookup: map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}, elements: []int{1, 2, 3, 4, 5}}

	if !reflect.DeepEqual(set, expected) {
		t.Errorf("Add: expected %v, got %v", expected, set)
	}
}

func TestRemove(t *testing.T) {
	set := NewOrderedSet(1, 2, 3, 4, 5)
	set.Remove(2, 4)

	expected := OrderedSet[int]{lookup: map[int]bool{1: true, 3: true, 5: true}, elements: []int{1, 3, 5}}

	if !reflect.DeepEqual(set, expected) {
		t.Errorf("Remove: expected %v, got %v", expected, set)
	}
}

func TestContains(t *testing.T) {
	set := NewOrderedSet("apple", "banana", "orange")

	// Test case 1: Elements present
	if !set.Contains("banana", "orange") {
		t.Errorf("Contains: expected true, got false")
	}

	// Test case 2: Element not present
	if set.Contains("grape") {
		t.Errorf("Contains: expected false, got true")
	}
}

func TestToSlice(t *testing.T) {
	set := NewOrderedSet(1, 2, 3)
	result := set.ToSlice()

	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ToSlice: expected %v, got %v", expected, result)
	}
}

func TestGet(t *testing.T) {
	set := NewOrderedSet("apple", "banana", "orange")

	// Test case 1: Valid index
	result1 := set.Get(1)
	expected1 := "banana"
	if result1 != expected1 {
		t.Errorf("Get: expected %v, got %v", expected1, result1)
	}

	// Test case 2: Invalid index
	result2 := set.Get(3)
	expected2 := ""
	if result2 != expected2 {
		t.Errorf("Get: expected %v, got %v", expected2, result2)
	}
}

func TestLen(t *testing.T) {
	set := NewOrderedSet("apple", "banana", "orange")
	result := set.Len()

	expected := 3

	if result != expected {
		t.Errorf("Len: expected %v, got %v", expected, result)
	}
}

func TestIsEmpty(t *testing.T) {
	// Test case 1: Non-empty set
	set1 := NewOrderedSet("apple", "banana", "orange")
	if set1.IsEmpty() {
		t.Errorf("IsEmpty: expected false, got true")
	}

	// Test case 2: Empty set
	set2 := NewOrderedSet[int]()
	if !set2.IsEmpty() {
		t.Errorf("IsEmpty: expected true, got false")
	}
}

func TestEquals(t *testing.T) {
	set1 := NewOrderedSet(1, 2, 3)
	set2 := NewOrderedSet(1, 2, 3)
	set3 := NewOrderedSet(1, 2, 4)

	// Test case 1: Equal sets
	if !set1.Equals(set2) {
		t.Errorf("Equals: expected true, got false")
	}

	// Test case 2: Unequal sets
	if set1.Equals(set3) {
		t.Errorf("Equals: expected false, got true")
	}
}

func TestSortWithComparator(t *testing.T) {
	set := NewOrderedSet(3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5)
	set.SortWithComparator(func(i, j int) bool {
		return set.elements[i] < set.elements[j]
	})

	expected := OrderedSet[int]{lookup: map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 9: true}, elements: []int{1, 2, 3, 4, 5, 6, 9}}

	if !reflect.DeepEqual(set, expected) {
		t.Errorf("SortWithComparator: expected %v, got %v", expected, set)
	}
}

func TestUnion(t *testing.T) {
	set1 := NewOrderedSet(1, 2, 3)
	set2 := NewOrderedSet(3, 4, 5)

	result := set1.Union(set2)

	expected := OrderedSet[int]{lookup: map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}, elements: []int{1, 2, 3, 4, 5}}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Union: expected %v, got %v", expected, result)
	}
}

func TestIntersection(t *testing.T) {
	set1 := NewOrderedSet(1, 2, 3, 4, 5)
	set2 := NewOrderedSet(3, 4, 5, 6, 7)

	result := set1.Intersection(set2)

	expected := OrderedSet[int]{lookup: map[int]bool{3: true, 4: true, 5: true}, elements: []int{3, 4, 5}}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Intersection: expected %v, got %v", expected, result)
	}
}

func TestDifference(t *testing.T) {
	set1 := NewOrderedSet(1, 2, 3, 4, 5)
	set2 := NewOrderedSet(3, 4, 5, 6, 7)

	result := set1.Difference(set2)

	expected := OrderedSet[int]{lookup: map[int]bool{1: true, 2: true}, elements: []int{1, 2}}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Difference: expected %v, got %v", expected, result)
	}
}

func TestClear(t *testing.T) {
	set := NewOrderedSet(1, 2, 3)
	set.Clear()

	expected := OrderedSet[int]{lookup: map[int]bool{}, elements: []int{}}

	if !reflect.DeepEqual(set, expected) {
		t.Errorf("Clear: expected %v, got %v", expected, set)
	}
}
