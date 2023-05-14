// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crypotoutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Md5(test *testing.T) {
	for _, tc := range []struct {
		name     string
		str      string
		slat     string
		expected string
	}{
		{
			name:     "normal",
			str:      "123",
			slat:     "a765ff3138693dbfa6888f05726588de53db72e1679a9341f04f9195d2ee3b8e",
			expected: "b4cc9d0fcff069fcd31afae0b7001434",
		},
	} {
		got := Md5(tc.str, tc.slat)
		assert.Equal(test, tc.expected, got)
	}
}
