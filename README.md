# unique

Reads a file and outputs all unique lines.

The tool keeps hashes of all unique lines in memory (in a map), so the size of the file it can process is limited by
your RAM. If you give it a large file it might OOM but it needs to be a *really* large file.

The reason this exists is that I need it. If it's useful to you as well - awesome.

## Installation

If you have [go](https://go.dev/) installed:  
`$ go install github.com/ro-tex/unique`

If you prefer a binary, you can download a Linux amd64 one from https://github.com/ro-tex/unique/releases.

## Usage

When no arguments are given, `unique` reads from the standard in.  
Running ```$ echo " one\n one\n two" | unique``` will output

```
 one
 two
```

In order to read from a file, use the `-f` parameter:  
```$ unique -f file.dat```

That is equivalent to `$ cat file.dat | unique`.
