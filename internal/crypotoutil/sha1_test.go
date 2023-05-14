// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crypotoutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sha1(test *testing.T) {
	for _, tc := range []struct {
		name     string
		str      string
		slat     string
		expected string
	}{
		{
			name:     "normal",
			str:      "a",
			expected: "86f7e437faa5a7fce15d1ddcb9eaeaea377667b8",
		},
	} {
		got := Sha1(tc.str)
		assert.Equal(test, tc.expected, got)
	}
}
