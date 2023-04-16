# unique

`unique` is simple tool that outputs the unique lines of its input. That input can come from `stdin` or from a file.

The reason `unique` exists is that I need it. If it's useful to you as well - awesome!

## Installation

If you have [Go](https://go.dev/) installed:

```shell
go install github.com/ro-tex/unique@latest
```

If you prefer a binary, you can download a Linux amd64 one from https://github.com/ro-tex/unique/releases.

## Usage

When no arguments are given, `unique` reads from the standard in.  
Running

```shell
echo " one\n one\n two\n one" | unique
```

will output

```
 one
 two
```

In order to read from a file, use the `-f` parameter:

```shell
unique -f file.dat
```

That is equivalent to

```shell
cat file.dat | unique
```

## Development

To run locally, run:

```shell
go run main.go <ARGS>
```

To install the local changes:

```shell
go install
```
