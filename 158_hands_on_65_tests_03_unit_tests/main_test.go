package main

import (
	"log"
	"testing"
)

func TestParadise(t *testing.T) {
	got := paradise("bali")
	want := "my idea of paradise is bali"
	if got != want {
		log.Fatalf("error - want %s and got %s", want, got)
	}
}
