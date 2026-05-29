# passgen

A simple CLI tool for generating secure passwords and copying them to the clipboard.

## Installation

### Go

```bash
go install github.com/Kalebe16/passgen@latest
```

### Binary

Download the latest release from:

https://github.com/Kalebe16/passgen/releases/latest

## Usage

```txt
Usage:
  passgen [flags]

Flags:
  -h, --help         help for passgen
  -l, --length int   password length (default 16)
  -L, --lowercase    include lowercase chars
  -n, --numbers      include numeric chars
  -s, --symbols      include symbol chars
  -u, --uppercase    include uppercase chars
  -v, --version      show version
```

Generate a password with uppercase, lowercase, numeric, and symbol characters:
```bash
passgen --uppercase --lowercase --numbers --symbols --length 32
```

Generate a password with uppercase characters only:
```bash
passgen --uppercase --length 32
```

Generate a password with lowercase characters only:
```bash
passgen --lowercase --length 32
```

Generate a password with numeric characters only:
```bash
passgen --numbers --length 32
```

Generate a password with symbol characters only:
```bash
passgen --symbols --length 32
```

Show the current version:
```bash
passgen --version
```
