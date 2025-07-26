package lib

import "testing"

func TestVarParser(t *testing.T) {
	AddVariable(map[string]interface{}{
		"name": "John",
		"age":  30,
	})

	AddVariable(map[string]interface{}{
		"city": "New York",
	})

	text := "Hello ${name}, you are ${age} years old and live in ${city}"
	parsed := Parse(text)

	if parsed != "Hello John, you are 30 years old and live in New York" {
		t.Errorf("Expected 'Hello John, you are 30 years old and live in New York', got '%s'", parsed)
	}
}
