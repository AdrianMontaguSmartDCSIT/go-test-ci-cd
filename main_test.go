package main

import "testing"

func TestSayHello(t *testing.T) {
	msg := sayHello("Aran")
	want := "Hello Aran"

	if msg != want {
		t.Errorf("got %s; want %s", msg, want)
	}
}
