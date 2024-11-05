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

func TestRemoveFromListById(t *testing.T) {
	tests := []struct {
		name    string
		items   []Item
		id      int
		want    []Item
		wantErr bool
	}{
		{
			name: "Remove existing item",
			items: []Item{
				{id: 1, name: "Item1", price: 10.0},
				{id: 2, name: "Item2", price: 20.0},
			},
			id: 1,
			want: []Item{
				{id: 2, name: "Item2", price: 20.0},
			},
			wantErr: false,
		},
		{
			name: "Remove non-existing item",
			items: []Item{
				{id: 1, name: "Item1", price: 10.0},
				{id: 2, name: "Item2", price: 20.0},
			},
			id:      3,
			want:    []Item{{id: 1, name: "Item1", price: 10.0}, {id: 2, name: "Item2", price: 20.0}},
			wantErr: true,
		},
		{
			name:    "Remove from empty list",
			items:   []Item{},
			id:      1,
			want:    []Item{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := removeFromListById(tt.items, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("removeFromListById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !equal(got, tt.want) {
				t.Errorf("removeFromListById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b []Item) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
