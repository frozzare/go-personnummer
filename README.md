# go-personnummer [![Build Status](https://travis-ci.org/frozzare/go-personnummer.svg?branch=master)](https://travis-ci.org/frozzare/go-personnummer) [![GoDoc](https://godoc.org/github.com/frozzare/go-personnummer?status.svg)](https://godoc.org/github.com/frozzare/go-personnummer) [![Go Report Card](https://goreportcard.com/badge/github.com/frozzare/go-personnummer)](https://goreportcard.com/report/github.com/frozzare/go-personnummer)

 Validate Swedish social security numbers.

## Installation

```
$ go get -u github.com/frozzare/go-personnummer
```

## Example

```go
package main

import (
	"github.com/frozzare/go-personnummer"
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

MIT © [Fredrik Forsmo](https://github.com/frozzare)