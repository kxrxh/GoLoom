package collections

type OrderedMap[T comparable] struct {
	keys   []T
	values map[T]interface{}
}

// NewOrderedMap creates a new instance of the OrderedMap struct.
//
// This function takes no parameters.
// It returns a pointer to an OrderedMap object.
func NewOrderedMap[T comparable]() *OrderedMap[T] {
	return &OrderedMap[T]{
		keys:   make([]T, 0),
		values: make(map[T]interface{}),
	}
}

// Set adds or updates a key-value pair in the OrderedMap.
//
// It checks if the key already exists in the map. If not, it appends the key to the
// list of keys in the OrderedMap. Then, it sets the value for the key.
func (m *OrderedMap[T]) Set(key T, value interface{}) {
	// Check if the key already exists
	if _, exists := m.values[key]; !exists {
		m.keys = append(m.keys, key)
	}
	// Set the value for the key
	m.values[key] = value
}

// Get returns the value associated with the given key and a boolean indicating
// whether the key exists in the OrderedMap.
//
// Parameters:
// - key: The key to get the value for.
//
// Returns:
// - value: The value associated with the key.
// - exists: A boolean indicating whether the key exists in the OrderedMap.
func (m *OrderedMap[T]) Get(key T) (interface{}, bool) {
	value, exists := m.values[key]
	return value, exists
}

// Keys returns the keys of the OrderedMap.
//
// It does not modify the OrderedMap.
// Returns a slice representing the keys in the OrderedMap.
func (m *OrderedMap[T]) Keys() []T {
	return m.keys
}

// Values returns a slice of all the values in the OrderedMap.
//
// No parameters are required.
// It returns a slice of interface{} that contains all the values in the OrderedMap.
func (m *OrderedMap[T]) Values() []interface{} {
	result := make([]interface{}, len(m.keys))
	for i, key := range m.keys {
		result[i] = m.values[key]
	}
	return result
}

// ToKeyValueArray converts the OrderedMap to a slice of KeyValue.
//
// No parameters.
// Returns a slice of KeyValue.
func (m *OrderedMap[T]) ToKeyValueArray() []KeyValue[T] {
	result := make([]KeyValue[T], len(m.keys))
	for i, key := range m.keys {
		result[i] = KeyValue[T]{Key: key, Value: m.values[key]}
	}
	return result
}

// KeyValue represents a key-value pair.
type KeyValue[T comparable] struct {
	Key   T
	Value interface{}
}

// Len returns the length of the OrderedMap.
//
// It does not modify the OrderedMap and returns an integer value.
func (m *OrderedMap[T]) Len() int {
	return len(m.keys)
}

// IsEmpty returns true if the OrderedMap is empty, otherwise returns false.
//
// No parameters.
// Returns a boolean value.
func (m *OrderedMap[T]) IsEmpty() bool {
	return len(m.keys) == 0
}

// Clear removes all elements from the ordered map.
//
//
func (m *OrderedMap[T]) Clear() {
	m.keys = make([]T, 0)
	m.values = make(map[T]interface{})
}

// Delete deletes the key-value pair with the specified key from the OrderedMap.
//
// Parameters:
//   - key: the key to be deleted from the OrderedMap.
//
// Return type(s):
//   None.
func (m *OrderedMap[T]) Delete(key T) {
	delete(m.values, key)
	for i, k := range m.keys {
		if k == key {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			break
		}
	}
}