package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []string
	}{
		"simple":      {input: " Hello World! ", want: []string{"hello", "world!"}},
		"many_spaces": {input: "   Hello  World!   ", want: []string{"hello", "world!"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
