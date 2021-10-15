package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	waifu := []string{
		"Elaina",
		"Mashiro",
		"Alice",
	}

	got := Map(waifu, func(s string) string {
		return s + " is my waifu!"
	})

	want := []string{
		"Elaina is my waifu!",
		"Mashiro is my waifu!",
		"Alice is my waifu!",
	}
	assert.Equal(t, got, want)
}
