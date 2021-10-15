// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crypotoutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Base64(test *testing.T) {
	for _, tc := range []struct {
		name     string
		str      []byte
		alphaSet string
		expected string
	}{
		{
			name:     "normal",
			str:      []byte("E99p1ant"),
			expected: "RTk5cDFhbnQ=",
		},
		{
			name:     "srun alpha set",
			str:      []byte("E99p1ant"),
			alphaSet: SrunAlphaSet,
			expected: "hFYeMJiTWq+=",
		},
	} {
		got := Base64(tc.str, tc.alphaSet)
		assert.Equal(test, tc.expected, got)
	}
}
