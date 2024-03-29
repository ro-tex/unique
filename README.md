# unique

`unique` is a simple tool that outputs the unique lines of its input.

The input can come from `stdin` or from a file. See [usage](#usage).

## Installation

### Use [Homebrew](https://brew.sh/)

```shell
brew install ro-tex/tap/unique
```

or

```
brew tap ro-tex/tap
brew install unique
```

### Build it yourself

You will need [Go](https://go.dev/) for this.

```shell
go install github.com/ro-tex/unique@latest
```

### Grab a binary

If you prefer a binary, you can download a Linux or Mac one from https://github.com/ro-tex/unique/releases.

## Usage

When no arguments are given, `unique` reads from the standard input.
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
