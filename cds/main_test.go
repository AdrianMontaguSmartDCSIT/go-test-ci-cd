package main

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Test with name Aran", "Aran", "Hello Aran"},
		{"Test with name John", "John", "Hello John"},
		{"Test with empty name", "", "Hello "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := sayHello(tt.input)
			if msg != tt.want {
				t.Errorf("got %s; want %s", msg, tt.want)
			}
		})
	}
}
