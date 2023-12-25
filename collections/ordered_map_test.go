package collections

import (
	"reflect"
	"testing"
)

func TestOrderedMap_KeysAndValues(t *testing.T) {
	myMap := NewOrderedMap[string]()

	// Test Keys and Values
	myMap.Set("one", 1)
	myMap.Set("two", 2)
	myMap.Set("three", 3)

	expectedKeys := []string{"one", "two", "three"}
	expectedValues := []interface{}{1, 2, 3}

	keys := myMap.Keys()
	values := myMap.Values()

	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("Expected keys %v, got %v", expectedKeys, keys)
	}

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Expected values %v, got %v", expectedValues, values)
	}
}

func TestOrderedMap_KeysAndValues2(t *testing.T) {
	myMap := NewOrderedMap[int]()

	// Test Keys and Values
	myMap.Set(1, "one")
	myMap.Set(2, "two")
	myMap.Set(3, "three")

	expectedKeys := []int{1, 2, 3}
	expectedValues := []interface{}{"one", "two", "three"}

	keys := myMap.Keys()
	values := myMap.Values()

	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("Expected keys %v, got %v", expectedKeys, keys)
	}

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Expected values %v, got %v", expectedValues, values)
	}
}

func TestOrderedMap_ToKeyValueArray(t *testing.T) {
	myMap := NewOrderedMap[string]()

	// Test ToKeyValueArray
	myMap.Set("one", 1)
	myMap.Set("two", 2)
	myMap.Set("three", 3)

	expectedArray := []KeyValue[string]{
		{Key: "one", Value: 1},
		{Key: "two", Value: 2},
		{Key: "three", Value: 3},
	}

	resultArray := myMap.ToKeyValueArray()

	if !reflect.DeepEqual(resultArray, expectedArray) {
		t.Errorf("Expected key-value array %v, got %v", expectedArray, resultArray)
	}
}

func TestOrderedMap_Clear(t *testing.T) {
	myMap := NewOrderedMap[string]()
	myMap.Set("one", 1)
	myMap.Set("two", 2)
	myMap.Clear()

	result := myMap.keys
	if len(result) != 0 {
		t.Errorf("Expected length 0, but got %v", len(result))
	}
}

func TestOrderedMap_Len(t *testing.T) {
	myMap := NewOrderedMap[string]()
	myMap.Set("one", 1)
	myMap.Set("two", 2)
	result := myMap.Len()
	if result != 2 {
		t.Errorf("Expected length 2, but got %v", result)
	}
}

func TestOrderedMap_IsEmpty(t *testing.T) {
	myMap := NewOrderedMap[string]()
	myMap.Set("one", 1)
	myMap.Set("two", 2)
	result := myMap.IsEmpty()
	if result {
		t.Errorf("Expected false, but got %v", result)
	}

	myMap.Clear()
	result = myMap.IsEmpty()
	if !result {
		t.Errorf("Expected true, but got %v", result)
	}
}

func TestOrderedMap_Delete(t *testing.T) {
	myMap := NewOrderedMap[string]()
	myMap.Set("one", 1)
	myMap.Set("two", 2)
	myMap.Delete("one")
	result := myMap.Len()
	if result != 1 {
		t.Errorf("Expected length 1, but got %v", result)
	}
}
