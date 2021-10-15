// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crypotoutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Encode(test *testing.T) {
	for _, tc := range []struct {
		name     string
		str      string
		slat     string
		expected string
	}{
		{
			name:     "normal",
			str:      "a",
			slat:     "b",
			expected: "s;!f\xf9u\xacP",
		},
	} {
		got, err := Encode(tc.str, tc.slat)
		assert.Nil(test, err)
		assert.Equal(test, tc.expected, got)
	}
}
