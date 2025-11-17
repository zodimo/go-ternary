package ternary

import (
	"testing"

	"github.com/zodimo/go-lazy"
)

func TestTernary(t *testing.T) {
	tests := []struct {
		name      string
		condition bool
		value1    interface{}
		value2    interface{}
		expected  interface{}
	}{
		{
			name:      "true condition returns value1 (int)",
			condition: true,
			value1:    10,
			value2:    20,
			expected:  10,
		},
		{
			name:      "false condition returns value2 (int)",
			condition: false,
			value1:    10,
			value2:    20,
			expected:  20,
		},
		{
			name:      "true condition returns value1 (string)",
			condition: true,
			value1:    "yes",
			value2:    "no",
			expected:  "yes",
		},
		{
			name:      "false condition returns value2 (string)",
			condition: false,
			value1:    "yes",
			value2:    "no",
			expected:  "no",
		},
		{
			name:      "true condition returns value1 (bool)",
			condition: true,
			value1:    true,
			value2:    false,
			expected:  true,
		},
		{
			name:      "false condition returns value2 (bool)",
			condition: false,
			value1:    true,
			value2:    false,
			expected:  false,
		},
		{
			name:      "true condition with zero values",
			condition: true,
			value1:    0,
			value2:    100,
			expected:  0,
		},
		{
			name:      "false condition with zero values",
			condition: false,
			value1:    100,
			value2:    0,
			expected:  0,
		},
		{
			name:      "true condition with empty strings",
			condition: true,
			value1:    "",
			value2:    "not empty",
			expected:  "",
		},
		{
			name:      "false condition with empty strings",
			condition: false,
			value1:    "not empty",
			value2:    "",
			expected:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result interface{}
			switch v1 := tt.value1.(type) {
			case int:
				result = Ternary(tt.condition, v1, tt.value2.(int))
			case string:
				result = Ternary(tt.condition, v1, tt.value2.(string))
			case bool:
				result = Ternary(tt.condition, v1, tt.value2.(bool))
			}

			if result != tt.expected {
				t.Errorf("Ternary() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestTernaryInt(t *testing.T) {
	result := Ternary(true, 42, 0)
	if result != 42 {
		t.Errorf("Ternary(true, 42, 0) = %d, want 42", result)
	}

	result = Ternary(false, 42, 0)
	if result != 0 {
		t.Errorf("Ternary(false, 42, 0) = %d, want 0", result)
	}
}

func TestTernaryString(t *testing.T) {
	result := Ternary(true, "positive", "negative")
	if result != "positive" {
		t.Errorf("Ternary(true, \"positive\", \"negative\") = %s, want \"positive\"", result)
	}

	result = Ternary(false, "positive", "negative")
	if result != "negative" {
		t.Errorf("Ternary(false, \"positive\", \"negative\") = %s, want \"negative\"", result)
	}
}

func TestTernaryLazy(t *testing.T) {
	// Track if each branch was evaluated
	evaluated1 := false
	evaluated2 := false

	condition := lazy.NewLazy(func() bool { return true })
	value1 := lazy.NewLazy(func() string {
		evaluated1 = true
		return "branch1"
	})
	value2 := lazy.NewLazy(func() string {
		evaluated2 = true
		return "branch2"
	})

	result := TernaryLazy(condition, value1, value2)

	// At this point, nothing should be evaluated yet
	if evaluated1 || evaluated2 {
		t.Error("Branches should not be evaluated before Get() is called")
	}

	// Now get the result
	value := result.Get()

	if value != "branch1" {
		t.Errorf("TernaryLazy() = %s, want \"branch1\"", value)
	}

	// Only value1 should have been evaluated
	if !evaluated1 {
		t.Error("value1 should have been evaluated")
	}
	if evaluated2 {
		t.Error("value2 should not have been evaluated when condition is true")
	}
}

func TestTernaryLazyFalseCondition(t *testing.T) {
	// Track if each branch was evaluated
	evaluated1 := false
	evaluated2 := false

	condition := lazy.NewLazy(func() bool { return false })
	value1 := lazy.NewLazy(func() string {
		evaluated1 = true
		return "branch1"
	})
	value2 := lazy.NewLazy(func() string {
		evaluated2 = true
		return "branch2"
	})

	result := TernaryLazy(condition, value1, value2)

	// At this point, nothing should be evaluated yet
	if evaluated1 || evaluated2 {
		t.Error("Branches should not be evaluated before Get() is called")
	}

	// Now get the result
	value := result.Get()

	if value != "branch2" {
		t.Errorf("TernaryLazy() = %s, want \"branch2\"", value)
	}

	// Only value2 should have been evaluated
	if evaluated1 {
		t.Error("value1 should not have been evaluated when condition is false")
	}
	if !evaluated2 {
		t.Error("value2 should have been evaluated")
	}
}

func TestTernaryLazyInt(t *testing.T) {
	condition := lazy.NewLazy(func() bool { return true })
	value1 := lazy.NewLazy(func() int { return 100 })
	value2 := lazy.NewLazy(func() int { return 200 })

	result := TernaryLazy(condition, value1, value2)
	value := result.Get()

	if value != 100 {
		t.Errorf("TernaryLazy() = %d, want 100", value)
	}
}

func TestTernaryLazyMultipleCalls(t *testing.T) {
	callCount1 := 0
	callCount2 := 0

	condition := lazy.NewLazy(func() bool { return true })
	value1 := lazy.NewLazy(func() string {
		callCount1++
		return "branch1"
	})
	value2 := lazy.NewLazy(func() string {
		callCount2++
		return "branch2"
	})

	result := TernaryLazy(condition, value1, value2)

	// Call Get() multiple times
	val1 := result.Get()
	val2 := result.Get()
	val3 := result.Get()

	// All calls should return the same value
	if val1 != "branch1" || val2 != "branch1" || val3 != "branch1" {
		t.Errorf("All calls should return \"branch1\", got %q, %q, %q", val1, val2, val3)
	}

	// value2 should never be evaluated
	if callCount2 != 0 {
		t.Errorf("value2 should not be evaluated, got %d calls", callCount2)
	}
}

