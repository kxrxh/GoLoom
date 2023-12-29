package collections

type Counter struct {
	elements map[interface{}]uint64
}

// NewCounter initializes a new Counter struct.
//
// No parameters.
// Returns a Counter with an initialized map.
func NewCounter() Counter {
	return Counter{elements: make(map[interface{}]uint64)}
}

// CounterFromSlice creates a Counter from a slice of elements.
//
// elements is a slice of any comparable type.
// Returns a new Counter object.
func CounterFromSlice[T comparable](elements []T) Counter {
	c := NewCounter()
	for _, e := range elements {
		c.elements[e]++
	}
	return c
}

// FromString creates and returns a Counter from a given string.
//
// It takes a string parameter `s` and iterates over each element in the string,
// incrementing the corresponding count in the Counter `c`.
// The function returns the created Counter `c`.
func FromString(s string) Counter {
	c := NewCounter()
	for _, e := range s {
		c.elements[e]++
	}
	return c
}

// Add increments the count for each element in the provided variadic
// slice.
//
// elems is a variadic slice of interface{} elements whose counts are
// to be increased.
func (c *Counter) Add(elems ...any) {
	for _, e := range elems {
		c.elements[e]++
	}
}

// Remove deletes the specified elements from the counter.
//
// Accepts a variadic number of interface{} elements to be removed.
// Does not return any value.
func (c *Counter) Remove(elems ...any) {
	for _, e := range elems {
		delete(c.elements, e)
	}
}

// Contains determines if all the provided elements are
// present in the Counter.
//
// 'elems' is a variadic parameter of elements to check
// for presence in the Counter.
// Returns true if all elements are present, otherwise
// false.
func (c *Counter) Contains(elems ...any) bool {
	for _, e := range elems {
		if _, ok := c.elements[e]; !ok {
			return false
		}
	}
	return true
}

// ToSlice converts the Counter's elements into a slice.
//
// It does not take any parameters.
// Returns a slice of interface{} containing all the elements.
func (c Counter) ToSlice() []any{
	var elems []interface{}
	for e := range c.elements {
		elems = append(elems, e)
	}
	return elems
}

// Len returns the number of unique elements in the Counter.
//
// This method has no parameters.
// Returns an uint64 representing the count of elements.
func (c Counter) Len() uint64 {
	return uint64(len(c.elements))
}

// Clear resets the Counter's elements to an empty state.
//
// No parameters.
// No return values.
func (c *Counter) Clear() {
	c.elements = make(map[interface{}]uint64)
}

// Get retrieves the count for the specified element.
//
// elem is the element for which to retrieve the count.
// Returns the count as a uint64.
func (c Counter) Get(elem interface{}) uint64 {
	return c.elements[elem]
}
