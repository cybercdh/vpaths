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
	"io"
	"net/url"
	"os"
	"strings"
)

func main() {

	// take piped input
	var input_urls io.Reader
	input_urls = os.Stdin
	sc := bufio.NewScanner(input_urls)

	// check there were no errors reading stdin (unlikely)
	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[!]	failed to read input: %s\n", err)
	}

	// keep track of urls we've seen
	seen := make(map[string]bool)

	for sc.Scan() {

		// parse each url
		_url := sc.Text()

		if !strings.HasPrefix(_url, "http") {
			continue
		}

		u, err := url.Parse(_url)
		if err != nil {
			continue
		}

		// split the paths from the parsed url
		paths := strings.Split(u.Path, "/")

		// iterate over the paths slice and print
		for i := 0; i < len(paths); i++ {
			path := paths[:len(paths)-i]
			tmp_url := fmt.Sprintf(u.Scheme + "://" + u.Host + strings.Join(path, "/"))

			// if we've seen the tmp_url already, keep moving
			if _, ok := seen[tmp_url]; ok {
				continue
			}

			// add to seen
			seen[tmp_url] = true

			// print the result
			fmt.Printf("%s\n", tmp_url)
		}

	}
}
