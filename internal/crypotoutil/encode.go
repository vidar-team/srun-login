// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crypotoutil

import (
	"github.com/pkg/errors"
	"github.com/robertkrimen/otto"
	"github.com/wuhan005/gadget"
)

const encodeScript = `
function encode(str, key) {
if (str === '') return '';
var v = s(str, true);
var k = s(key, false);
if (k.length < 4) k.length = 4;
var n = v.length - 1,
	z = v[n],
	y = v[0],
	c = 0x86014019 | 0x183639A0,
	m,
	e,
	p,
	q = Math.floor(6 + 52 / (n + 1)),
	d = 0;

while (0 < q--) {
	d = d + c & (0x8CE0D9BF | 0x731F2640);
	e = d >>> 2 & 3;

	for (p = 0; p < n; p++) {
		y = v[p + 1];
		m = z >>> 5 ^ y << 2;
		m += y >>> 3 ^ z << 4 ^ (d ^ y);
		m += k[p & 3 ^ e] ^ z;
		z = v[p] = v[p] + m & (0xEFB8D130 | 0x10472ECF);
	}

	y = v[0];
	m = z >>> 5 ^ y << 2;
	m += y >>> 3 ^ z << 4 ^ (d ^ y);
	m += k[p & 3 ^ e] ^ z;
	z = v[n] = v[n] + m & (0xBB390742 | 0x44C6F8BD);
}

return l(v, false);
}

function s(a, b) {
var c = a.length;
var v = [];

for (var i = 0; i < c; i += 4) {
	v[i >> 2] = a.charCodeAt(i) | a.charCodeAt(i + 1) << 8 | a.charCodeAt(i + 2) << 16 | a.charCodeAt(i + 3) << 24;
}

if (b) v[v.length] = c;
return v;
}

function l(a, b) {
var d = a.length;
var c = d - 1 << 2;

if (b) {
	var m = a[d - 1];
	if (m < c - 3 || m > c) return null;
	c = m;
}

for (var i = 0; i < d; i++) {
	a[i] = String.fromCharCode(a[i] & 0xff, a[i] >>> 8 & 0xff, a[i] >>> 16 & 0xff, a[i] >>> 24 & 0xff);
}

return b ? a.join('').substring(0, c) : a.join('');
}

escape(encode(str, salt));
`

// Encode is the internal encode function of srun.
func Encode(str, salt string) (string, error) {
	vm := otto.New()
	_ = vm.Set("str", str)
	_ = vm.Set("salt", salt)
	value, err := vm.Run(encodeScript)
	if err != nil {
		return "", errors.Wrap(err, "run vm")
	}

	result, err := value.ToString()
	if err != nil {
		return "", errors.Wrap(err, "to string")
	}
	// Go can't parse the unicode characters from JavaScript.
	// So we url-encode the result in JavaScript VM and decode it in Go.
	return gadget.URLDecode(result), nil
}
