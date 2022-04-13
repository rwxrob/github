// Copyright 2022 Robert Muhlestein.
// SPDX-License-Identifier: Apache-2.0

// Package github provides high-level functions that are called from the Go
// Bonzai branch of the same name providing universal access to the core
// functionality. Use of interfaces over structs with public properties
// is to allow the gradual addition of useful data from the underlying
// JSON structs.
package github

import (
	"fmt"
	"os"

	web "github.com/rwxrob/web/pkg"
)

// -------------------------- Package Globals -------------------------

// Host specifies the host to use for all GitHub web requests. It is the
// public "github.com" by default. Anything else will trigger GitHub
// Enterprise assumptions. This is overridden at init() by the GH_HOST
// environment variable consistent with the official GitHub CLI
// tool.
var Host = "github.com"

func init() {
	host := os.Getenv("GH_HOST")
	if host != "" {
		Host = host
	}
}

// APIVersion is the default used when returning the APIURL.
var APIVersion = "v3"

// ------------------------------ Client ------------------------------

// Client specified a GitHub client capable of returned a limited number
// of data points from the API. This interface will grow as the struct
// returned by NewClient sufficiently adds new, common queries.
//
// Public Versus Enterprise
//
// It's not particularly well-known, but the API queries for both public
// GitHub and GitHub Enterprise differ primarily only in the host and
// URL that is used. GitHub Enterprise requires a token with more
// permissions than is permitted by GitHub public API as well.
//
//      github.com (public)
//      github.example.com (enterprise)
//
type Client interface {

	// Host of github.com (the default
	Host() string
	SetHost(a string)
	APIVersion() string
	SetAPIVersion(a string)
	Repo(id string) (map[string]any, error)

	// APIURL returns a full URL string for the given Host based on
	// inferred usage of GitHub Enterprise:
	//
	//     Host == github.com         -> https://api.github.com/
	//     Host == github.example.com -> https://github.example.com/api/v3
	//
	APIURL() string
}

// NewClient returns a new struct pointer fulfilling the Client
// interface.
func NewClient() *client {
	c := new(client)
	c.host = Host
	c.apivers = APIVersion
	return c
}

type client struct {
	host    string
	apivers string
}

func (c *client) Host() string           { return c.host }
func (c *client) SetHost(a string)       { c.host = a }
func (c *client) APIVersion() string     { return c.apivers }
func (c *client) SetAPIVersion(a string) { c.apivers = a }

func (c *client) Repo(id string) (map[string]any, error) {
	data := map[string]any{}
	url := c.APIURL(`repos/` + id)
	req := web.Req{URL: url, Data: data}
	if err := req.Submit(); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *client) APIURL(suf string) string {
	if c.host == "github.com" {
		return "https://api.github.com/" + suf
	}
	return fmt.Sprintf("https://%v/api/%v/%v", c.host, c.apivers, suf)
}

func (c *client) Latest(is string) string {
	return ""
}

// --------------------------- API Functions --------------------------

/*

// LatestURL takes a GitHub git repo source URL and converts into one
// pointing at the API URL that returns JSON data about the latest
// release.
func LatestURL(src string) string {
	var url string
https://api.github.com/repos/twitchdev/twitch-cli/releases/latest | jq -r .name)
curl -sSL "https://github.com/twitchdev/twitch-cli/releases/download/${latest}/twitch-cli_${latest#v}_Linux_x86_64.tar.gz" -o tarbomb.tgz
	return url
}

// LatestVersion returns the version (name) of the latest release from
// the GitHub repo at the give source URL. Only `git@` and `http://`
// source URLs are currently supported (but full repo source URL
// support, including ssh, is planned).
func LatestVersion(src string) string {
	var url string
	return url

}
*/
