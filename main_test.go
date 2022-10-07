package main

import (
	"testing"
	"time"
)

func TestReverse(t *testing.T) {
	if got, want := Reverse("reverse"), "esrever"; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	time.Sleep(30 * time.Second)

	t.Run("empty string", func(t *testing.T) {
		if got, want := Reverse(""), ""; got != want {
			t.Fatalf("got %v, want %v", got, want)
		}
	})
}
