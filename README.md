# go-personnummer [![Build Status](https://travis-ci.org/frozzare/go-personnummer.svg?branch=master)](https://travis-ci.org/frozzare/go-personnummer)

 Validate Swedish personal identity numbers.

 View the [docs](http://godoc.org/github.com/frozzare/go-personnummer).

## Installation

```
$ go get github.com/frozzare/go-personnummer
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

# License

Copyright (c) 2014 Fredrik Forsmo
Licensed under the MIT license.
