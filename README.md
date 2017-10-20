# personnummer [![Build Status](https://travis-ci.org/personnummer/go.svg?branch=master)](https://travis-ci.org/personnummer/go) [![GoDoc](https://godoc.org/github.com/personnummer/go?status.svg)](https://godoc.org/github.com/personnummer/go) [![Go Report Card](https://goreportcard.com/badge/github.com/personnummer/go)](https://goreportcard.com/report/github.com/personnummer/go)

 Validate Swedish social security numbers.

## Installation

```
go get -u github.com/personnummer/go
```

## Example

```go
package main

import (
	personnummer "github.com/personnummer/go"
)

func main() {
	personnummer.Valid(6403273813)
	//=> true
	
	personnummer.Valid("19900101-0017")
	//=> true
	
	// works with co-ordination numbers
	personnummer.Valid("701063-2391")
	//=> true

	personnummer.Valid("510818-916")
	//=> false
}
```

## License

MIT
