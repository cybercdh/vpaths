package main

import (
	"reflect"
	"testing"
)

func TestExpand(t *testing.T) {
	cases := map[string][]string{
		"https://www.example.com/foo/bar/baz/index.html": {
			"https://www.example.com/foo/bar/baz/index.html",
			"https://www.example.com/foo/bar/baz",
			"https://www.example.com/foo/bar/baz/",
			"https://www.example.com/foo/bar",
			"https://www.example.com/foo/bar/",
			"https://www.example.com/foo",
			"https://www.example.com/foo/",
			"https://www.example.com",
			"https://www.example.com/",
		},
		"https://example.org/hello/world/": {
			"https://example.org/hello/world/",
			"https://example.org/hello/world",
			"https://example.org/hello",
			"https://example.org/hello/",
			"https://example.org",
			"https://example.org/",
		},
		// a dotted directory such as an API version must still get its slash form
		"https://api.example.com/v1.2/users?id=5#top": {
			"https://api.example.com/v1.2/users",
			"https://api.example.com/v1.2/users/",
			"https://api.example.com/v1.2",
			"https://api.example.com/v1.2/",
			"https://api.example.com",
			"https://api.example.com/",
		},
		// percent-encoding is preserved, not decoded and re-joined
		"https://example.com/a%20b/c%2Fd": {
			"https://example.com/a%20b/c%2Fd",
			"https://example.com/a%20b/c%2Fd/",
			"https://example.com/a%20b",
			"https://example.com/a%20b/",
			"https://example.com",
			"https://example.com/",
		},
		"https://example.com": {
			"https://example.com",
			"https://example.com/",
		},
		"":                      nil,
		"   ":                   nil,
		"example.com/no/scheme": nil,
		"http://":               nil,
	}
	for in, want := range cases {
		if got := expand(in); !reflect.DeepEqual(got, want) {
			t.Errorf("expand(%q)\n got %v\nwant %v", in, got, want)
		}
	}
}
