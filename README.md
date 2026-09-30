# vpaths

Prints all paths from a list of URIs, e.g.

```
$ cat urls.txt
https://www.example.com/foo/bar/baz/index.html
https://example.org/hello/world/yolo.js

$ cat urls.txt | vpaths
https://www.example.com/foo/bar/baz/index.html
https://www.example.com/foo/bar/baz
https://www.example.com/foo/bar/baz/
https://www.example.com/foo/bar
https://www.example.com/foo/bar/
https://www.example.com/foo
https://www.example.com/foo/
https://www.example.com
https://www.example.com/
https://example.org/hello/world/yolo.js
https://example.org/hello/world
https://example.org/hello/world/
https://example.org/hello
https://example.org/hello/
https://example.org
https://example.org/
```

Each parent path is printed in both `/dir` and `/dir/` form. The input URL itself only gets a slash form when its last segment does not look like a file. Query strings and fragments are dropped, and lines that are not absolute URLs are skipped.

## Install

If you have Go installed and configured (i.e. with `$GOPATH/bin` in your `$PATH`):

```
go install github.com/cybercdh/vpaths@latest
```

## Usage

```
$ cat urls.txt | vpaths
```

## Uses

Useful for piping into other toolsets in order to expand your FUZZing. 
