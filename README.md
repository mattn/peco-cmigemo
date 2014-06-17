# peco-cmigemo

C/Migemo matcher for peco

## Installation

```
$ go build peco-cmigemo.go
```

## How to use?

Add below into you `~/.peco/config.json`

```json
"C/Migemo": [
  "/path/to/your/peco-cmigemo", "$QUERY"
]
```

Put `peco-cmigemo` and `dict` into `~/.peco/`.

Then you can choice `C/Migemo` matcher with typing `CTRL-R` few times.
