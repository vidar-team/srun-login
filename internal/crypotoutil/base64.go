// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crypotoutil

import (
	"encoding/base64"
)

// SrunAlphaSet is the srun customize base64 alpha set.
const SrunAlphaSet = "LVoJPiCN2R8G90yg+hmFHuacZ1OWMnrsSTXkYpUq/3dlbfKwv6xztjI7DeBE45QA"

// Base64 will base64 encode the given string with the given alpha set.
// If the alphaSet not provided, it uses the default alpha set.
func Base64(str []byte, alphaSet ...string) string {
	encoding := base64.StdEncoding
	if len(alphaSet) != 0 && alphaSet[0] != "" {
		encoding = base64.NewEncoding(alphaSet[0])
	}
	return encoding.EncodeToString(str)
}
