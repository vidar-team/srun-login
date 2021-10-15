// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package srun

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/wuhan005/gadget"

	"github.com/vidar-team/srun-login/internal/crypotoutil"
)

type userInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IP       string `json:"ip"`
	AcID     string `json:"acid"`
	EncVer   string `json:"enc_ver"`
}

// EncodeUserInfo encodes the current user info with the given challenge code.
func (c *Client) EncodeUserInfo(challenge string) (string, error) {
	// We would like to keep the JSON keys in order, just like the JavaScript did in browser,
	// so struct used here instead of a map.
	jsonBytes, err := json.Marshal(userInfo{
		Username: c.username,
		Password: c.password,
		IP:       c.ip,
		AcID:     c.acID,
		EncVer:   "srun_bx1",
	})
	if err != nil {
		return "", errors.Wrap(err, "encode JSON")
	}

	encode, err := crypotoutil.Encode(string(jsonBytes), challenge)
	if err != nil {
		return "", errors.Wrap(err, "encode")
	}
	return "{SRBX1}" + crypotoutil.Base64([]byte(encode), crypotoutil.SrunAlphaSet), nil
}

// MakeChksum returns the checksum of the current request with the given challenge code.
func (c *Client) MakeChksum(challenge string) (string, error) {
	hmd5, err := crypotoutil.Md5(c.password, challenge)
	if err != nil {
		return "", errors.Wrap(err, "hmd5")
	}

	userInfo, err := c.EncodeUserInfo(challenge)
	if err != nil {
		return "", errors.Wrap(err, "encode user info")
	}

	fileds := []string{
		c.username,
		hmd5,
		c.acID,
		c.ip,
		c.n,
		c.typ,
		userInfo,
	}
	var str string
	for _, f := range fileds {
		str += challenge + f
	}

	return gadget.Sha1(str), nil
}
