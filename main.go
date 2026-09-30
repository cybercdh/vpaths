/*

vpaths
- takes a list of URIs and prints all paths within the structure, e.g.
- https://www.example.com/foo/bar/baz/index.html
- https://www.example.com/foo/bar/baz/
- https://www.example.com/foo/bar/
- https://www.example.com/foo/
- https://www.example.com/

usage
$ cat urls.txt | vpaths

*/

package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// keep track of urls we've seen
	seen := make(map[string]bool)

	for sc.Scan() {
		for _, p := range expand(sc.Text()) {
			if seen[p] {
				continue
			}
			seen[p] = true
			fmt.Println(p)
		}
	}

	// check there were no errors reading stdin (unlikely)
	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[!]\tfailed to read input: %s\n", err)
		os.Exit(1)
	}
}

// expand returns the URL itself followed by each parent path up to the host
// root. Every parent is emitted with a trailing slash. The URL itself gets a
// slash form too unless its last segment looks like a file (contains a dot).
// Lines that are not absolute URLs are skipped. The path is kept as given
// (percent-encoding intact) and any query string or fragment is dropped.
func expand(line string) []string {
	line = strings.TrimSpace(line)
	if line == "" || !strings.Contains(line, "://") {
		return nil
	}
	u, err := url.Parse(line)
	if err != nil || u.Host == "" {
		return nil
	}
	base := u.Scheme + "://" + u.Host
	segs := strings.Split(u.EscapedPath(), "/")

	var out []string
	add := func(s string) {
		for _, e := range out {
			if e == s {
				return
			}
		}
		out = append(out, s)
	}
	for i := 0; i < len(segs); i++ {
		p := strings.Join(segs[:len(segs)-i], "/")
		add(base + p)
		if strings.HasSuffix(p, "/") {
			continue
		}
		last := segs[len(segs)-1-i]
		if i == 0 && strings.Contains(last, ".") {
			continue // a file, not a directory
		}
		add(base + p + "/")
	}
	return out
}
