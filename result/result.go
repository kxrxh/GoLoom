package result

type Result[T any] struct {
	Value T
	Err   error
}

// NewResult creates a new Result encapsulating a value and an error.
//
// value is the data to be stored of any type.
// err is the error to be associated with the value.
// Returns a new Result containing the provided value and error.
func NewResult[T any](value T, err error) Result[T] {
	return Result[T]{
		Value: value,
		Err:   err,
	}
}

// ErrResult wraps an error in a Result type.
//
// err: The error to wrap.
// Return Result: A Result containing the error.
func ErrResult[T any](err error) Result[T] {
	return Result[T]{
		Err: err,
	}
}

// OkResult wraps a given value in a Result type.
//
// value: The value to be wrapped.
// Return: Result containing the value.
func OkResult[T any](value T) Result[T] {
	return Result[T]{
		Value: value,
	}
}

// IsSuccess checks if the Result is a success.
//
// It returns a boolean value indicating whether the Result has an error or not.
func (r Result[T]) IsSuccess() bool {
	return r.Err == nil
}

// IsFailure checks if the Result is a failure.
//
// It returns a boolean value indicating whether the Result is a failure.
func (r Result[T]) IsFailure() bool {
	return r.Err != nil
}

// Get returns the value stored in the Result struct.
//
// It does not take any parameters.
// It returns an interface{} type.
func (r Result[T]) Get() interface{} {
	return r.Value
}

// GetError returns the error associated with the result.
//
// No parameters.
// Returns an error.
func (r Result[T]) GetError() error {
	return r.Err
}

// GetOrDefault returns the value of the Result object if it is successful, otherwise it returns the defaultValue.
//
// defaultValue: the value to be returned if the Result object is not successful.
// interface{}: the type of the value to be returned.
func (r Result[T]) GetOrDefault(defaultValue T) T {
	if r.IsSuccess() {
		return r.Value
	}
	return defaultValue
}

// GetOrElse returns the value of the Result if it is successful. Otherwise, it returns the default value
// provided by the defaultValue function.
//
// defaultValue: A function that returns the default value.
// interface{}: The type of the value returned by the function.
func (r Result[T]) GetOrElse(defaultValue func() interface{}) interface{} {
	if r.IsSuccess() {
		return r.Value
	}
	return defaultValue()
}

// GetOrElseGet returns the value of the Result if it's successful.
//
// If the Result is a failure, it calls defaultValue and returns its result.
// Returns an interface{} which is either the Result value or the default.
func (r Result[T]) GetOrElseGet(defaultValue func() Result[T]) interface{} {
	if r.IsSuccess() {
		return r.Value
	}
	return defaultValue().Get()
}

// GetOrElseThrow returns the result's value if it's a success
// or throws an error using the provided defaultValue function.
//
// defaultValue is a function that returns an error.
// Return type is an interface{} which is either the result's
// value or the error thrown by defaultValue.
func (r Result[T]) GetOrElseThrow(defaultValue func() error) interface{} {
	if r.IsSuccess() {
		return r.Value
	}
	return defaultValue()
}
