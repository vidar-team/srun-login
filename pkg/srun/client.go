// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package srun

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"github.com/vidar-team/srun-login/internal/crypotoutil"
)

// Client is the client of srun.
type Client struct {
	host     string
	username string
	password string

	ip   string
	acID string
	typ  string
	n    string
}

// NewClient returns a new srun client with the provided host, username and password.
func NewClient(host, username, password string) *Client {
	return &Client{
		host:     host,
		username: username,
		password: password,

		acID: "0",
		typ:  "1",
		n:    "200",
	}
}

type ChallengeResponse struct {
	Challenge string `json:"challenge"`
	ClientIp  string `json:"client_ip"`
	Ecode     int    `json:"ecode"`
	Error     string `json:"error"`
	ErrorMsg  string `json:"error_msg"`
	Expire    string `json:"expire"`
	OnlineIp  string `json:"online_ip"`
	Res       string `json:"res"`
	SrunVer   string `json:"srun_ver"`
	St        int    `json:"st"`
}

// GetChallenge requests srun to get a challenge code.
func (c *Client) GetChallenge() (*ChallengeResponse, error) {
	url := fmt.Sprintf(c.host+"/cgi-bin/get_challenge?callback=_&username=%s&ip=%s", c.username, c.ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "http get")
	}
	defer func() { _ = resp.Body.Close() }()

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}
	if len(responseBytes) < 3 {
		return nil, errors.Errorf("unexpected response body: %q", string(responseBytes))
	}
	responseBytes = responseBytes[2 : len(responseBytes)-1]

	var respJSON ChallengeResponse
	if err := json.Unmarshal(responseBytes, &respJSON); err != nil {
		return nil, errors.Wrap(err, "decode JSON")
	}

	c.ip = respJSON.ClientIp

	return &respJSON, nil
}

type PortalResponse struct {
	ClientIp string `json:"client_ip"`
	Error    string `json:"error"`
	ErrorMsg string `json:"error_msg"`
	OnlineIp string `json:"online_ip"`
	Res      string `json:"res"`
	SrunVer  string `json:"srun_ver"`
	St       int    `json:"st"`
}

// Portal login to srun with the username and password.
// A challenge code from `GetChallenge()` should be provided.
func (c *Client) Portal(challenge string) (*PortalResponse, error) {
	u, err := url.Parse(c.host + "/cgi-bin/srun_portal")
	if err != nil {
		return nil, errors.Wrap(err, "parse URL")
	}

	query := url.Values{}
	query.Set("callback", "_")
	query.Set("action", "login")
	query.Set("username", c.username)

	passwordMd5 := crypotoutil.Md5(c.password, challenge)

	query.Set("password", "{MD5}"+passwordMd5)
	query.Set("os", "Mac OS")
	query.Set("name", "Macintosh")
	query.Set("double_stack", "0")

	userInfo, err := c.EncodeUserInfo(challenge)
	if err != nil {
		return nil, errors.Wrap(err, "encode user info")
	}
	query.Set("info", userInfo)

	checkSum, err := c.MakeChksum(challenge)
	if err != nil {
		return nil, errors.Wrap(err, "make chksum")
	}
	query.Set("chksum", checkSum)

	query.Set("ac_id", c.acID)
	query.Set("ip", c.ip)
	query.Set("n", c.n)
	query.Set("type", c.typ)

	u.RawQuery = query.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, errors.Wrap(err, "http get")
	}
	defer func() { _ = resp.Body.Close() }()
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}
	if len(responseBytes) < 3 {
		return nil, errors.Errorf("unexpected response body: %q", string(responseBytes))
	}

	responseBytes = responseBytes[2 : len(responseBytes)-1]

	var respJSON PortalResponse
	if err := json.Unmarshal(responseBytes, &respJSON); err != nil {
		return nil, errors.Wrap(err, "decode JSON")
	}
	return &respJSON, nil
}
