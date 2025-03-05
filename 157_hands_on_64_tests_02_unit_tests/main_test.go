package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	got := add(3, 5)
	want := 8
	if got != want {
		log.Fatalf("error - want %v and got %v", want, got)

	}
}
