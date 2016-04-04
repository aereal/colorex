# colorex

colorex (colorize + regular expression) colorize input lines that match given patterns.

## Usage

```
# cat ... | colorex COLOR:PATTERN
cat access_log | colorex yellow:404
```

## Build

```
go build ./...
```

## Author

aereal (aereal@aereal.org)
