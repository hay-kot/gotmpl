# Gotmpl

Tiny CLI template engine for generating files quickly

Gotmpl brings Go's Templates to your terminal along with some helpful functions courtesy of [Sprig](http://masterminds.github.io/sprig/). Usage is extremely simple, after installing the binary, you can use it like so:

```bash
gotmpl render --tmpl ./template.tmpl --data ./data.json > ./output.txt
```
## Features

- [x] Sprig functions
- [x] Support for JSON, YAML, and TOML data files

## Installation

### Homebrew

```bash
brew install hay-kot/tap/gotmpl
```

### Go

```bash
go install https://github.com/hay-kot/gotmpl
```



