package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := HelloWorld()
	want := "Hello, World!"
	if got != want {
		t.Errorf("Got %q and want %q", got, want)
	}
}
