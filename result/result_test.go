package result

import (
	"fmt"
	"testing"
)

func ReturnResult(value int) Result[int] {
	if value < 0 {
		return ErrResult[int](fmt.Errorf("invalid value: %d", value))
	}
	return OkResult(value)
}

func TestHandleResultTest(t *testing.T) {
	result := ReturnResult(1)
	if !result.IsSuccess() {
		t.Errorf("expected success, got %v", result)
	}
	if result.Get() != 1 {
		t.Errorf("expected 1, got %v", result.Get())
	}

	result = ReturnResult(-1)
	if result.IsSuccess() {
		t.Errorf("expected failure, got %v", result)
	}

	if result.GetError() == nil {
		t.Errorf("expected error, got %v", result)
	}

	if result.GetOrDefault(10) != 10 {
		t.Errorf("expected nil, got %v", result)
	}
}
