# vpaths

Prints all paths from a list of URIs, e.g.

```
$ cat urls.txt
https://www.example.com/foo/bar/baz/index.html
https://example.org/hello/world/yolo.js

$ cat urls.txt | vhosts
https://www.example.com/foo/bar/baz/index.html
https://www.example.com/foo/bar/baz/
https://www.example.com/foo/bar/
https://www.example.com/foo/
https://www.example.com/
https://example.org/hello/world/yolo.js
https://example.org/hello/world/
https://example.org/hello/
```

## Install

If you have Go installed and configured (i.e. with `$GOPATH/bin` in your `$PATH`):

```
go get -u github.com/cybercdh/vpaths
```

## Usage

```
$ cat urls.txt | vpaths
```

## Uses

Useful for piping into other toolsets in order to expand your FUZZing. 
