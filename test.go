package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Continuous Delivery is awesome!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
