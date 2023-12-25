package collections


// RemoveElement removes the first occurrence of an element from a slice.
//
// Parameters:
// - slice: The input slice from which the element will be removed.
// - element: The element to be removed from the slice.
//
// Returns:
// The modified slice with the element removed. If the element is not found, the original slice is returned.
func RemoveElement[T comparable](slice []T, element T) []T {
    // Find the index of the element in the slice
    for i := 0; i < len(slice); i++ {
        if slice[i] == element {
            // Create a new slice excluding the element to remove
            return append(slice[:i], slice[i+1:]...)
        }
    }
    // If the element is not found, return the original slice
    return slice
}